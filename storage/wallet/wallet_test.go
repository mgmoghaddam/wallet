package wallet_test

import (
	"errors"
	"testing"
	repomocks "wallet/mocks/repomocks/wallet"

	"github.com/stretchr/testify/assert"

	wallet "wallet/storage/wallet"
)

func TestCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Prepare
		fakeWallet := &wallet.Wallet{ID: 1, Balance: 1000} // adjust this per your Wallet struct

		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("Create", fakeWallet).Return(nil) // you expect Create function to return nil error, which indicates success

		// Act
		err := mockRepo.Create(fakeWallet)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t) // This checks if the function has been called with correct parameters
	})

	t.Run("error", func(t *testing.T) {
		// Prepare
		fakeWallet := &wallet.Wallet{ID: 2, Balance: 2000} // adjust this per your Wallet struct

		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("Create", fakeWallet).Return(errors.New("forced error")) // you expect Create function to return an error

		// Act
		err := mockRepo.Create(fakeWallet)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t) // This checks if the function has been called with correct parameters
	})
}

func TestDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Prepare
		fakeWalletID := int64(1) // adjust this per your Wallet struct

		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("Delete", fakeWalletID).Return(nil) // you expect Create function to return nil error, which indicates success

		// Act
		err := mockRepo.Delete(fakeWalletID)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t) // This checks if the function has been called with correct parameters
	})

	t.Run("error", func(t *testing.T) {
		// Prepare
		fakeWalletID := int64(2) // adjust this per your Wallet struct

		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("Delete", fakeWalletID).Return(errors.New("forced error")) // you expect Create function to return an error

		// Act
		err := mockRepo.Delete(fakeWalletID)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t) // This checks if the function has been called with correct parameters
	})
}

func TestDeleteByMemberID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Prepare
		fakeMemberID := int64(1) // adjust this per your Wallet struct

		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("DeleteByMemberID", fakeMemberID).Return(nil) // you expect Create function to return nil error, which indicates success

		// Act
		err := mockRepo.DeleteByMemberID(fakeMemberID)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t) // This checks if the function has been called with correct parameters
	})

	t.Run("error", func(t *testing.T) {
		// Prepare
		fakeMemberID := int64(2) // adjust this per your Wallet struct

		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("DeleteByMemberID", fakeMemberID).Return(errors.New("forced error")) // you expect Create function to return an error

		// Act
		err := mockRepo.DeleteByMemberID(fakeMemberID)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t) // This checks if the function has been called with correct parameters
	})
}

func TestGetByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Prepare
		fakeWalletID := int64(1) // adjust this per your Wallet struct
		fakeWallet := &wallet.Wallet{ID: 1, Balance: 1000}

		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("GetByID", fakeWalletID).Return(fakeWallet, nil) // you expect Create function to return nil error, which indicates success

		// Act
		result, err := mockRepo.GetByID(fakeWalletID)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, fakeWallet, result)
		mockRepo.AssertExpectations(t) // This checks if the function has been called with correct parameters
	})

	t.Run("error", func(t *testing.T) {
		// Prepare
		fakeWalletID := int64(2) // adjust this per your Wallet struct

		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("GetByID", fakeWalletID).Return(nil, errors.New("forced error")) // you expect Create function to return an error

		// Act
		result, err := mockRepo.GetByID(fakeWalletID)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "forced error", err.Error())
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t) // This checks if the function has been called with correct parameters
	})
}

func TestGetByMemberID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Prepare
		fakeMemberID := int64(1)                               // adjust this per your Wallet struct
		fakeWallet := []*wallet.Wallet{{ID: 1, Balance: 1000}} // make this a slice

		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("GetByMemberID", fakeMemberID).Return(fakeWallet, nil) // you expect Create function to return nil error, which indicates success

		// Act
		result, err := mockRepo.GetByMemberID(fakeMemberID)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, fakeWallet, result)
		mockRepo.AssertExpectations(t) // This checks if the function has been called with correct parameters
	})

	t.Run("error", func(t *testing.T) {
		// Prepare
		fakeMemberID := int64(2) // adjust this per your Wallet struct

		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("GetByMemberID", fakeMemberID).Return(nil, errors.New("forced error")) // you expect Create function to return an error

		// Act
		result, err := mockRepo.GetByMemberID(fakeMemberID)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "forced error", err.Error())
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t) // This checks if the function has been called with correct parameters
	})
}

func TestUpdateBalance(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Prepare
		fakeWalletID := int64(1) // adjust this per your Wallet struct
		fakeBalance := int64(1000)

		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("UpdateBalance", fakeWalletID, fakeBalance).Return(nil) // you expect Create function to return nil error, which indicates success

		// Act
		err := mockRepo.UpdateBalance(fakeWalletID, fakeBalance)

		// Assert
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t) // This checks if the function has been called with correct parameters
	})

	t.Run("error", func(t *testing.T) {
		// Prepare
		fakeWalletID := int64(2) // adjust this per your Wallet struct
		fakeBalance := int64(2000)

		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("UpdateBalance", fakeWalletID, fakeBalance).Return(errors.New("forced error")) // you expect Create function to return an error

		// Act
		err := mockRepo.UpdateBalance(fakeWalletID, fakeBalance)

		// Assert
		assert.Error(t, err)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t) // This checks if the function has been called with correct parameters
	})
}
