package wallet

const walletColumns = "id" + ",member_id,wallet_name,balance,created_at,updated_at"

func (s Storage) Create(w *Wallet) error {
	sqlStmt := `
	INSERT INTO wallet (member_id, wallet_name, balance) VALUES ($1, $2, $3) 
	                     RETURNING id, wallet_name, created_at, updated_at`
	err := s.db.QueryRow(sqlStmt, w.MemberID, w.WalletName, w.Balance).Scan(&w.ID, &w.WalletName,
		&w.CreatedAt, &w.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (s Storage) UpdateBalance(id int64, balance int64) error {
	sqlStmt := `
	UPDATE wallet SET balance = $1, updated_at = now() WHERE id = $2
	                     RETURNING updated_at`
	row, err := s.db.Exec(sqlStmt, balance, id)
	if err != nil {
		return err
	}
	if count, err := row.RowsAffected(); err != nil || count == 0 {
		return ErrNoRowToUpdate
	}
	return nil
}

func (s Storage) GetByID(id int64) (*Wallet, error) {
	sqlStmt := "SELECT " + walletColumns + " FROM wallet WHERE id = $1"
	w := &Wallet{}
	err := s.db.QueryRow(sqlStmt, id).Scan(&w.ID, &w.MemberID, &w.WalletName, &w.Balance, &w.CreatedAt, &w.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (s Storage) GetByMemberID(memberID int64) ([]*Wallet, error) {
	// one member can have multi wallet
	sqlStmt := "SELECT " + walletColumns + " FROM wallet WHERE member_id = $1"
	rows, err := s.db.Query(sqlStmt, memberID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	wallets := make([]*Wallet, 0)
	for rows.Next() {
		w := &Wallet{}
		err := rows.Scan(&w.ID, &w.MemberID, &w.WalletName, &w.Balance, &w.CreatedAt, &w.UpdatedAt)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, w)
	}
	return wallets, nil
}

func (s Storage) Delete(id int64) error {
	sqlStmt := "DELETE FROM wallet WHERE id = $1"
	row, err := s.db.Exec(sqlStmt, id)
	if err != nil {
		return err
	}
	if count, err := row.RowsAffected(); err != nil || count == 0 {
		return ErrNoRowToUpdate
	}
	return nil
}

func (s Storage) DeleteByMemberID(memberID int64) error {
	sqlStmt := "DELETE FROM wallet WHERE member_id = $1"
	row, err := s.db.Exec(sqlStmt, memberID)
	if err != nil {
		return err
	}
	if count, err := row.RowsAffected(); err != nil || count == 0 {
		return ErrNoRowToUpdate
	}
	return nil
}
