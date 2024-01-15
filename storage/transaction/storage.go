package transaction

import (
	"database/sql"
	"wallet/db"
)

type Repository interface {
	Insert(t *Transaction) error
	GetByID(id int64) (*Transaction, error)
	GetByWalletID(walletID int64) ([]*Transaction, error)
	GetByWalletIDWithPagination(walletID int64, limit, offset int) ([]*Transaction, error)
	GetByWalletIDAndType(walletID int64, transactionType Type) ([]*Transaction, error)
	GetByWalletIDAndDiscountCode(walletID int64, discountCode string) ([]*Transaction, error)
	GetByWalletIDAndTypeAndDiscountCode(walletID int64, transactionType Type, discountCode string) ([]*Transaction, error)
	GetByDiscountCodeWithPagination(discountCode string, limit, offset int) ([]*Transaction, error)
	DeleteByWalletID(walletID int64) error
	DeleteByWalletIDAndType(walletID int64, transactionType Type) error
	DeleteByWalletIDAndDiscountCode(walletID int64, discountCode string) error
	DeleteByID(id int64) error
	GetBalance(walletID int64) (int64, error)
	WithTX(tx *sql.Tx) (Repository, error)
}

type Storage struct {
	db db.SQLExt
}

func NewStorage(db *sql.DB) Storage {
	return Storage{db: db}
}

// WithTX returns a new storage with the given transaction replacing the db.
func (s Storage) WithTX(tx *sql.Tx) (Repository, error) {
	if tx == nil {
		return nil, db.ErrNoTXProvided
	}
	switch s.db.(type) {
	case *sql.Tx:
		return nil, db.ErrAlreadyInTX
	case *sql.DB:
		return Storage{db: tx}, nil
	}
	return s, nil
}
