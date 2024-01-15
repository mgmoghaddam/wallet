package wallet

import "time"

type DTO struct {
	ID         int64     `json:"id"`
	MemberID   int64     `json:"memberID"`
	WalletName string    `json:"walletName"`
	Balance    int64     `json:"balance"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type CreateRequest struct {
	MemberID   int64  `json:"memberID"`
	WalletName string `json:"walletName"`
	Balance    int64  `json:"balance"`
}

type AddGiftRequest struct {
	MemberID int64  `json:"memberID"`
	WalletID int64  `json:"walletID"`
	GiftCode string `json:"giftCode"`
}
