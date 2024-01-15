package transaction

import (
	"database/sql"
	"sync"
	"wallet/storage/transaction"
)

type UseCase interface {
	Create(r *CreateRequest) (*DTO, error)
	GetByID(id int64) (*DTO, error)
	GetByWalletID(walletID int64) ([]*DTO, error)
	GetByWalletIDWithPagination(walletID int64, limit, offset int) ([]*DTO, error)
	GetByWalletIDAndType(walletID int64, transactionType Type) ([]*DTO, error)
	GetByWalletIDAndDiscountCode(walletID int64, discountCode string) ([]*DTO, error)
	GetByWalletIDAndTypeAndDiscountCode(walletID int64, transactionType Type, discountCode string) ([]*DTO, error)
	GetByDiscountCodeWithPagination(discountCode string, limit, offset int) ([]*DTO, error)
	Delete(id int64) error
	DeleteByWalletID(walletID int64) error
	GetBalance(walletID int64) (int64, error)
	WithTX(tx *sql.Tx) (*Service, error)
}

type Service struct {
	transaction transaction.Repository
	mu          sync.Mutex

	inTx bool
}

func New(
	transaction transaction.Repository,
) *Service {
	return &Service{
		transaction: transaction,
	}
}

func (s *Service) WithTX(tx *sql.Tx) (*Service, error) {
	service := *s
	t, err := s.transaction.WithTX(tx)
	if err != nil {
		return nil, err
	}
	service.transaction = t
	service.inTx = true
	return &service, nil
}

func (s *Service) ToDBModel(t *DTO) *transaction.Transaction {
	return &transaction.Transaction{
		ID:              t.ID,
		WalletID:        t.WalletID,
		Amount:          t.Amount,
		TransactionType: TypeToDBType(t.TransactionType),
		Description:     t.Description,
		DiscountCode:    t.DiscountCode,
	}

}

func (s *Service) FromDBModel(t *transaction.Transaction) *DTO {
	return &DTO{
		ID:              t.ID,
		WalletID:        t.WalletID,
		Amount:          t.Amount,
		TransactionType: DbTypeToType(t.TransactionType),
		Description:     t.Description,
		DiscountCode:    t.DiscountCode,
		CreatedAt:       t.CreatedAt,
	}
}

func (s *Service) FromCreateRequest(r *CreateRequest) *transaction.Transaction {
	return &transaction.Transaction{
		WalletID:        r.WalletID,
		Amount:          r.Amount,
		TransactionType: TypeToDBType(r.TransactionType),
		Description:     r.Description,
		DiscountCode:    r.DiscountCode,
	}
}

func TypeToDBType(t Type) transaction.Type {
	switch t {
	case Recharge:
		return transaction.Recharge
	case Gift:
		return transaction.Gift
	case Withdraw:
		return transaction.Withdraw
	case Payment:
		return transaction.Payment
	case Refund:
		return transaction.Refund
	case Transfer:
		return transaction.Transfer
	default:
		return ""
	}
}

func DbTypeToType(t transaction.Type) Type {
	switch t {
	case transaction.Recharge:
		return Recharge
	case transaction.Gift:
		return Gift
	case transaction.Withdraw:
		return Withdraw
	case transaction.Payment:
		return Payment
	case transaction.Refund:
		return Refund
	case transaction.Transfer:
		return Transfer
	default:
		return ""
	}
}
