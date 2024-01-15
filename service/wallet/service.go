package wallet

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"
	"wallet/client/discount"
	"wallet/db"
	"wallet/internal/config"
	"wallet/service/transaction"
	"wallet/storage/wallet"
)

const walletPrefix = "WALLET:%s"

type UseCase interface {
	Create(r *CreateRequest) (*DTO, error)
	GetByID(id int64) (*DTO, error)
	GetByMemberID(memberID int64) ([]*DTO, error)
	AddGift(r *AddGiftRequest) (*DTO, error)
	Recharge(id, amount int64) (*DTO, error)
	Transfer(fromID, toID, amount int64) (*DTO, error)
	Withdraw(id, amount int64) (*DTO, error)
	Refund(id int64) (*DTO, error)
	Delete(id int64) error
	DeleteByMemberID(memberID int64) error
	GetByDiscountCodeWithPagination(discountCode string, limit, offset int) ([]*DTO, error)
	CreateTransactionAndUpdateWallet(id, amount int64, transactionType transaction.Type, description, discountCode string) (*DTO, error)
	WithTX(tx *sql.Tx) (*Service, error)
}

type Service struct {
	wallet      wallet.Repository
	transaction transaction.UseCase
	rdb         db.RedisClient

	discount discount.Client

	mu   sync.Mutex
	inTx bool
}

func New(
	wallet wallet.Repository,
	transaction transaction.UseCase,
	discount discount.Client,
	rdb db.RedisClient,
) *Service {
	return &Service{
		wallet:      wallet,
		transaction: transaction,
		discount:    discount,
		rdb:         rdb,
	}
}

func (s *Service) WithTX(tx *sql.Tx) (*Service, error) {
	service := *s
	w, err := s.wallet.WithTX(tx)
	if err != nil {
		return nil, err
	}
	service.wallet = w
	service.inTx = true
	return &service, nil
}

func (s *Service) ToDBModel(w *DTO) *wallet.Wallet {
	return &wallet.Wallet{
		ID:       w.ID,
		MemberID: w.MemberID,
		Balance:  w.Balance,
	}

}

func (s *Service) FromDBModel(w *wallet.Wallet) *DTO {
	return &DTO{
		ID:        w.ID,
		MemberID:  w.MemberID,
		Balance:   w.Balance,
		CreatedAt: w.CreatedAt,
		UpdatedAt: w.UpdatedAt,
	}
}

func (s *Service) FromCreateRequest(r *CreateRequest) *wallet.Wallet {
	return &wallet.Wallet{
		WalletName: r.WalletName,
		MemberID:   r.MemberID,
		Balance:    r.Balance,
	}
}

func (s *Service) UpdateOrInsertInRedis(key string, g *DTO, exp time.Duration) error {
	err := s.rdb.Set(context.Background(), config.RDBPrefix()+key, g, exp)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) RetrieveFromRedis(key string) (*DTO, error) {
	w, err := s.rdb.Get(context.Background(), config.RDBPrefix()+key)
	if err != nil {
		return nil, err
	}
	walletRecord := &DTO{}
	err = json.Unmarshal([]byte(w), walletRecord)
	if err != nil {
		return nil, err
	}
	return walletRecord, nil
}

func (s *Service) RemoveGiftFromRedis(w *DTO) {
	key := fmt.Sprintf(walletPrefix, strconv.FormatInt(w.MemberID, 10))
	s.RemoveWithKey(key)
}

func (s *Service) RemoveWithKey(k string) {
	s.rdb.Del(context.Background(), config.RDBPrefix()+k)
}

func (w *DTO) MarshalBinary() ([]byte, error) {
	return json.Marshal(w)
}
