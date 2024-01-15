package wallet

import "time"

type Wallet struct {
	ID         int64     `db:"id"`
	MemberID   int64     `db:"member_id"`
	WalletName string    `db:"wallet_name"`
	Balance    int64     `db:"balance"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
