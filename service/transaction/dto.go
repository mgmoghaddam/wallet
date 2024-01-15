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

type DTO struct {
	ID              int64     `json:"id"`
	WalletID        int64     `json:"walletID"`
	Amount          int64     `json:"amount"`
	TransactionType Type      `json:"transactionType"`
	Description     string    `json:"description"`
	DiscountCode    string    `json:"discountCode"`
	CreatedAt       time.Time `json:"createdAt"`
}

type CreateRequest struct {
	WalletID        int64  `json:"walletID"`
	Amount          int64  `json:"amount"`
	TransactionType Type   `json:"transactionType"`
	Description     string `json:"description"`
	DiscountCode    string `json:"discountCode"`
}
