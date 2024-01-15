package wallet

import (
	"context"
	"database/sql"
	"github.com/rs/zerolog/log"
	"time"
	"wallet/db"
	"wallet/internal/serr"
	"wallet/service/transaction"
)

// create wallet
func (s *Service) Create(r *CreateRequest) (*DTO, error) {
	w := s.FromCreateRequest(r)
	err := s.wallet.Create(w)
	if err != nil {
		return nil, err
	}
	return s.FromDBModel(w), nil
}

// get wallet by id
func (s *Service) GetByID(id int64) (*DTO, error) {
	w, err := s.wallet.GetByID(id)
	if err != nil {
		return nil, err
	}
	return s.FromDBModel(w), nil
}

// get wallets by member id
func (s *Service) GetByMemberID(memberID int64) ([]*DTO, error) {
	ws, err := s.wallet.GetByMemberID(memberID)
	if err != nil {
		return nil, err
	}
	var result []*DTO
	for _, w := range ws {
		result = append(result, s.FromDBModel(w))
	}
	return result, nil
}

func (s *Service) CreateTransactionAndUpdateWallet(id, amount int64, transactionType transaction.Type, description, discountCode string) (*DTO, error) {
	w, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}
	err = db.Transaction(context.Background(), func(tx *sql.Tx) error {
		txService, err := s.WithTX(tx)
		if err != nil {
			return err
		}
		tr := &transaction.CreateRequest{
			WalletID:        id,
			Amount:          amount,
			TransactionType: transactionType,
			Description:     description,
			DiscountCode:    discountCode,
		}

		t, err := txService.transaction.Create(tr)
		if err != nil {
			return err
		}
		log.Info().Msgf("transaction: %+v", t)
		err = txService.wallet.UpdateBalance(id, w.Balance)
		if err != nil {
			return err
		}
		w.Balance += t.Amount
		return nil
	})
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (s *Service) AddGift(r *AddGiftRequest) (*DTO, error) {
	layout := "2006-01-02T15:04:05Z07:00"
	g, err := s.discount.GetGiftByCode(r.GiftCode)
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, serr.ValidationErr("gift", "gift not found", serr.ErrGiftNotFound)
	}
	if g.UsedCount >= g.UsageLimit {
		return nil, serr.ValidationErr("gift", "gift usage limit reached", serr.ErrGiftUsageLimitReached)
	}
	expirationDate, err := time.Parse(layout, g.ExpirationDate)
	if err != nil {
		return nil, err
	}
	startDateTime, err := time.Parse(layout, g.StartDateTime)
	if expirationDate.Before(time.Now()) {
		return nil, serr.ValidationErr("gift", "gift expired", serr.ErrGiftExpired)
	}
	if startDateTime.After(time.Now()) {
		return nil, serr.ValidationErr("gift", "gift not started", serr.ErrGiftNotStarted)
	}

	ws, err := s.GetByMemberID(r.MemberID)
	if err != nil {
		return nil, err
	}

	for _, w := range ws {
		t, _ := s.transaction.GetByWalletIDAndDiscountCode(w.ID, r.GiftCode)
		if t != nil {
			return nil, serr.ValidationErr("wallet", "discount code has been used", serr.ErrDiscountCodeUsed)
		}
	}

	gift, err := s.discount.UseGift(r.GiftCode)
	if err != nil {
		return nil, err
	}
	s.RemoveWithKey(":MEMBER:" + r.GiftCode)
	// Create a transaction and update the wallet
	return s.CreateTransactionAndUpdateWallet(r.WalletID, gift.GiftAmount, transaction.Gift, "add gift transaction", gift.Code)
}

func (s *Service) Recharge(id, amount int64) (*DTO, error) {
	return s.CreateTransactionAndUpdateWallet(id, amount, transaction.Recharge, "add recharge transaction", "")
}

func (s *Service) Transfer(fromID, toID, amount int64) (*DTO, error) {
	ws, err := s.GetByID(fromID)
	if err != nil {
		return nil, err
	}
	if ws.Balance < amount {
		return nil, serr.ValidationErr("wallet", "not enough balance", serr.ErrNotEnoughBalance)
	}
	_, err = s.GetByID(toID)
	if err != nil {
		return nil, err
	}
	err = db.Transaction(context.Background(), func(tx *sql.Tx) error {
		txService, err := s.WithTX(tx)
		if err != nil {
			return err
		}
		_, err = txService.CreateTransactionAndUpdateWallet(fromID, -amount, transaction.Transfer, "transfer transaction", "")
		if err != nil {
			return err
		}
		_, err = txService.CreateTransactionAndUpdateWallet(toID, amount, transaction.Transfer, "transfer transaction", "")
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return ws, nil
}

// withdraw wallet balance
func (s *Service) Withdraw(id, amount int64) (*DTO, error) {
	w, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}
	if w.Balance < amount {
		return nil, serr.ValidationErr("wallet", "not enough balance", serr.ErrNotEnoughBalance)
	}
	return s.CreateTransactionAndUpdateWallet(id, -amount, transaction.Withdraw, "withdraw transaction", "")
}

// refund wallet balance by transaction id
func (s *Service) Refund(id int64) (*DTO, error) {
	t, err := s.transaction.GetByID(id)
	if err != nil {
		return nil, err
	}
	if t.TransactionType != transaction.Withdraw {
		return nil, serr.ValidationErr("transaction", "transaction type is not withdraw",
			serr.ErrTransactionTypeNotWithdrawal)
	}
	return s.CreateTransactionAndUpdateWallet(t.WalletID, t.Amount, transaction.Refund, "refund transaction", "")
}

// delete wallet by id and all transactions of that wallet
func (s *Service) Delete(id int64) error {
	err := db.Transaction(context.Background(), func(tx *sql.Tx) error {
		s, err := s.WithTX(tx)
		if err != nil {
			return err
		}
		err = s.transaction.DeleteByWalletID(id)
		if err != nil {
			return err
		}
		err = s.wallet.Delete(id)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// delete wallet by member id and all transactions of that wallet
func (s *Service) DeleteByMemberID(memberID int64) error {
	err := db.Transaction(context.Background(), func(tx *sql.Tx) error {
		s, err := s.WithTX(tx)
		if err != nil {
			return err
		}
		ws, err := s.GetByMemberID(memberID)
		if err != nil {
			return err
		}
		for _, w := range ws {
			err = s.transaction.DeleteByWalletID(w.ID)
			if err != nil {
				return err
			}
			err = s.wallet.Delete(w.ID)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// get all wallets which use specific discount code in transactions of that wallet by offset and limit
func (s *Service) GetByDiscountCodeWithPagination(discountCode string, limit, offset int) ([]*DTO, error) {
	ts, err := s.transaction.GetByDiscountCodeWithPagination(discountCode, limit, offset)
	if err != nil {
		return nil, err
	}
	var result []*DTO
	for _, t := range ts {
		w, err := s.GetByID(t.WalletID)
		if err != nil {
			return nil, err
		}
		result = append(result, w)
	}
	return result, nil
}
