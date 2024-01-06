package member

import (
	"database/sql"
	"errors"
	"wallet/db"
)

var (
	ErrNoRowToUpdate = errors.New("no row to update")
)

type Storage struct {
	db db.SQLExt
}

func New(db *sql.DB) Storage {
	return Storage{db: db}
}

// WithTX returns a new storage with the given transaction replacing the db.
func (s Storage) WithTX(tx *sql.Tx) (Storage, error) {
	if tx == nil {
		return Storage{}, db.ErrNoTXProvided
	}
	switch s.db.(type) {
	case *sql.Tx:
		return Storage{}, db.ErrAlreadyInTX
	case *sql.DB:
		return Storage{db: tx}, nil
	}
	return s, nil
}

func (s Storage) ScanMember(scanner db.Scanner) (*Member, error) {
	m := &Member{}
	err := scanner.Scan(&m.ID, &m.FirstName, &m.LastName, &m.Email, &m.Phone, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return m, nil
}
