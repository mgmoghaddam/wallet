package wallet_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	repomocks "wallet/mocks/repomocks/wallet"
	wallet "wallet/service/wallet"
)

func TestWalletService_AddGift(t *testing.T) {
	// Create a new instance of our mock UseCase object.
	mockUseCase := repomocks.NewUseCase(t)

	// Define our expected result.
	expected := &wallet.DTO{
		ID:         1,
		MemberID:   1,
		WalletName: "a",
		Balance:    0,
	}

	// Define our request.
	request := &wallet.AddGiftRequest{
		MemberID: 1,
		WalletID: 1,
		GiftCode: "Test",
	}

	// Expect the AddGift function to be called with our request.
	mockUseCase.On("AddGift", request).Return(expected, nil)

	// Call AddGift on our wallet service.
	result, err := mockUseCase.AddGift(request)

	// Assert that the function did not return an error.
	assert.NoError(t, err)
	// Assert that the function returned our expected result.
	assert.EqualValues(t, expected, result)

	// Assert that the expectations were met.
	mockUseCase.AssertExpectations(t)
}

func TestWalletService_AddGift_Error(t *testing.T) {
	// Create a new instance of our mock UseCase object.
	mockUseCase := repomocks.NewUseCase(t)

	// Define our request.
	request := &wallet.AddGiftRequest{
		MemberID: 1,
		WalletID: 1,
		GiftCode: "Test",
	}

	// Expect the AddGift function to be called with our request.
	mockUseCase.On("AddGift", request).Return(nil, assert.AnError)

	// Call AddGift on our wallet service.
	result, err := mockUseCase.AddGift(request)

	// Assert that the function returned an error.
	assert.Error(t, err)
	// Assert that the function did not return a result.
	assert.Nil(t, result)

	// Assert that the expectations were met.
	mockUseCase.AssertExpectations(t)
}
