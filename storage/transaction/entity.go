package transaction

import "time"

type Type string

const (
	Recharge Type = "recharge"
	Gift     Type = "gift"
	Withdraw Type = "withdraw"
	Payment  Type = "payment"
	Refund   Type = "refund"
	Transfer Type = "transfer"
)

type Transaction struct {
	ID              int64     `db:"id"`
	WalletID        int64     `db:"wallet_id"`
	Amount          int64     `db:"amount"`
	TransactionType Type      `db:"transaction_type"`
	Description     string    `db:"description"`
	DiscountCode    string    `db:"discount_code"`
	CreatedAt       time.Time `db:"created_at"`
}
