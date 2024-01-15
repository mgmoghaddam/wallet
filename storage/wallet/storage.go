package wallet

import (
	"database/sql"
	"errors"
	"wallet/db"
)

var (
	ErrNoRowToUpdate = errors.New("no row to update")
)

type Repository interface {
	Create(w *Wallet) error
	UpdateBalance(id int64, balance int64) error
	GetByID(id int64) (*Wallet, error)
	GetByMemberID(memberID int64) ([]*Wallet, error)
	Delete(id int64) error
	DeleteByMemberID(memberID int64) error
	WithTX(tx *sql.Tx) (Repository, error)
}

type Storage struct {
	db db.SQLExt
}

func NewStorage(db *sql.DB) Storage {
	return Storage{db: db}
}

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
