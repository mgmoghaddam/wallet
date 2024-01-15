package wallet_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"wallet/service/wallet"
)

func TestDTO_MarshalBinary(t *testing.T) {
	dto := &wallet.DTO{
		ID:         1,
		MemberID:   2,
		WalletName: "TestWallet",
		Balance:    1000,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	binaryData, err := dto.MarshalBinary()
	assert.NoError(t, err)
	assert.NotNil(t, binaryData)
}
