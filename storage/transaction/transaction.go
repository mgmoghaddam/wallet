package transaction

import (
	"wallet/internal/serr"
)

const transactionColumns = "id,wallet_id,amount,transaction_type,description,discount_code,created_at"

func (s Storage) Insert(t *Transaction) error {
	err := s.db.QueryRow(`
		INSERT INTO transaction
		    (wallet_id, amount, transaction_type, description, discount_code)
		VALUES 
		    ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`, t.WalletID, t.Amount, t.TransactionType, t.Description, t.DiscountCode).Scan(&t.ID, &t.CreatedAt)
	if err != nil {
		return serr.DBError("Insert", "transaction", err)
	}
	return nil
}

func (s Storage) GetByID(id int64) (*Transaction, error) {
	sqlStmt := "SELECT " + transactionColumns + " FROM transaction WHERE id = $1"
	t := &Transaction{}
	err := s.db.QueryRow(sqlStmt, id).Scan(&t.ID, &t.WalletID, &t.Amount, &t.TransactionType, &t.Description, &t.DiscountCode, &t.CreatedAt)
	if err != nil {
		return nil, serr.DBError("GetByID", "transaction", err)
	}
	return t, nil
}

func (s Storage) GetByWalletID(walletID int64) ([]*Transaction, error) {
	sqlStmt := "SELECT " + transactionColumns + " FROM transaction WHERE wallet_id = $1 ORDER BY created_at DESC"
	rows, err := s.db.Query(sqlStmt, walletID)
	if err != nil {
		return nil, serr.DBError("GetByWalletID", "transaction", err)
	}
	defer rows.Close()
	transactions := make([]*Transaction, 0)
	for rows.Next() {
		t := &Transaction{}
		err := rows.Scan(&t.ID, &t.WalletID, &t.Amount, &t.TransactionType, &t.Description, &t.DiscountCode, &t.CreatedAt)
		if err != nil {
			return nil, serr.DBError("GetByWalletID", "transaction", err)
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

func (s Storage) GetByWalletIDWithPagination(walletID int64, limit, offset int) ([]*Transaction, error) {
	sqlStmt := "SELECT " + transactionColumns + " FROM transaction WHERE wallet_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3"
	rows, err := s.db.Query(sqlStmt, walletID, limit, offset)
	if err != nil {
		return nil, serr.DBError("GetByWalletIDWithPagination", "transaction", err)
	}
	defer rows.Close()
	transactions := make([]*Transaction, 0)
	for rows.Next() {
		t := &Transaction{}
		err := rows.Scan(&t.ID, &t.WalletID, &t.Amount, &t.TransactionType, &t.Description, &t.DiscountCode, &t.CreatedAt)
		if err != nil {
			return nil, serr.DBError("GetByWalletIDWithPagination", "transaction", err)
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

func (s Storage) GetByWalletIDAndType(walletID int64, transactionType Type) ([]*Transaction, error) {
	sqlStmt := "SELECT " + transactionColumns + " FROM transaction WHERE wallet_id = $1 AND transaction_type = $2"
	rows, err := s.db.Query(sqlStmt, walletID, transactionType)
	if err != nil {
		return nil, serr.DBError("GetByWalletIDAndType", "transaction", err)
	}
	defer rows.Close()
	transactions := make([]*Transaction, 0)
	for rows.Next() {
		t := &Transaction{}
		err := rows.Scan(&t.ID, &t.WalletID, &t.Amount, &t.TransactionType, &t.Description, &t.DiscountCode, &t.CreatedAt)
		if err != nil {
			return nil, serr.DBError("GetByWalletIDAndType", "transaction", err)
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

func (s Storage) GetByWalletIDAndDiscountCode(walletID int64, discountCode string) ([]*Transaction, error) {
	sqlStmt := "SELECT " + transactionColumns + " FROM transaction WHERE wallet_id = $1 AND discount_code = $2"
	rows, err := s.db.Query(sqlStmt, walletID, discountCode)
	if err != nil {
		return nil, serr.DBError("GetByWalletIDAndDiscountCode", "transaction", err)
	}
	defer rows.Close()
	transactions := make([]*Transaction, 0)
	for rows.Next() {
		t := &Transaction{}
		err := rows.Scan(&t.ID, &t.WalletID, &t.Amount, &t.TransactionType, &t.Description, &t.DiscountCode, &t.CreatedAt)
		if err != nil {
			return nil, serr.DBError("GetByWalletIDAndDiscountCode", "transaction", err)
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

func (s Storage) GetByWalletIDAndTypeAndDiscountCode(walletID int64, transactionType Type, discountCode string) ([]*Transaction, error) {
	sqlStmt := "SELECT " + transactionColumns + " FROM transaction WHERE wallet_id = $1 AND transaction_type = $2 AND discount_code = $3"
	rows, err := s.db.Query(sqlStmt, walletID, transactionType, discountCode)
	if err != nil {
		return nil, serr.DBError("GetByWalletIDAndTypeAndDiscountCode", "transaction", err)
	}
	defer rows.Close()
	transactions := make([]*Transaction, 0)
	for rows.Next() {
		t := &Transaction{}
		err := rows.Scan(&t.ID, &t.WalletID, &t.Amount, &t.TransactionType, &t.Description, &t.DiscountCode, &t.CreatedAt)
		if err != nil {
			return nil, serr.DBError("GetByWalletIDAndTypeAndDiscountCode", "transaction", err)
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

// get by discount code and pagination sorted by created_at
func (s Storage) GetByDiscountCodeWithPagination(discountCode string, limit, offset int) ([]*Transaction, error) {
	sqlStmt := "SELECT " + transactionColumns + " FROM transaction WHERE discount_code = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3"
	rows, err := s.db.Query(sqlStmt, discountCode, limit, offset)
	if err != nil {
		return nil, serr.DBError("GetByDiscountCodeWithPagination", "transaction", err)
	}
	defer rows.Close()
	transactions := make([]*Transaction, 0)
	for rows.Next() {
		t := &Transaction{}
		err := rows.Scan(&t.ID, &t.WalletID, &t.Amount, &t.TransactionType, &t.Description, &t.DiscountCode, &t.CreatedAt)
		if err != nil {
			return nil, serr.DBError("GetByDiscountCodeWithPagination", "transaction", err)
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

// delete all transactions of a wallet
func (s Storage) DeleteByWalletID(walletID int64) error {
	sqlStmt := "DELETE FROM transaction WHERE wallet_id = $1"
	_, err := s.db.Exec(sqlStmt, walletID)
	if err != nil {
		return serr.DBError("DeleteByWalletID", "transaction", err)
	}
	return nil
}

// delete all transactions of a wallet with a specific type
func (s Storage) DeleteByWalletIDAndType(walletID int64, transactionType Type) error {
	sqlStmt := "DELETE FROM transaction WHERE wallet_id = $1 AND transaction_type = $2"
	_, err := s.db.Exec(sqlStmt, walletID, transactionType)
	if err != nil {
		return serr.DBError("DeleteByWalletIDAndType", "transaction", err)
	}
	return nil
}

// delete all transactions of a wallet with a specific discount code
func (s Storage) DeleteByWalletIDAndDiscountCode(walletID int64, discountCode string) error {
	sqlStmt := "DELETE FROM transaction WHERE wallet_id = $1 AND discount_code = $2"
	_, err := s.db.Exec(sqlStmt, walletID, discountCode)
	if err != nil {
		return serr.DBError("DeleteByWalletIDAndDiscountCode", "transaction", err)
	}
	return nil
}

// delete a transaction with transaction id
func (s Storage) DeleteByID(id int64) error {
	sqlStmt := "DELETE FROM transaction WHERE id = $1"
	_, err := s.db.Exec(sqlStmt, id)
	if err != nil {
		return serr.DBError("DeleteByID", "transaction", err)
	}
	return nil
}

// calculate amount of a wallet
func (s Storage) GetBalance(walletID int64) (int64, error) {
	sqlStmt := "SELECT sum(amount) FROM transaction WHERE wallet_id = $1"
	var balance int64
	err := s.db.QueryRow(sqlStmt, walletID).Scan(&balance)
	if err != nil {
		return 0, serr.DBError("GetBalance", "transaction", err)
	}
	return balance, nil
}
