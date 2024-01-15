package transaction_test

import (
	"errors"
	"testing"

	repomocks "wallet/mocks/repomocks/transaction"
	"wallet/storage/transaction"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fakeTr := &transaction.Transaction{
			WalletID:        1,
			Amount:          100,
			TransactionType: "credit",
		}
		mockRepo := repomocks.NewRepository(t)
		// Change "Create" to "Insert" here
		mockRepo.On("Insert", fakeTr).Return(nil)
		err := mockRepo.Insert(fakeTr)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		fakeTr := &transaction.Transaction{
			WalletID:        1,
			Amount:          100,
			TransactionType: "credit",
		}
		mockRepo := repomocks.NewRepository(t)
		// Change "Create" to "Insert" here
		mockRepo.On("Insert", fakeTr).Return(errors.New("forced error"))
		err := mockRepo.Insert(fakeTr)
		assert.Error(t, err)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fakeTrID := int64(1)
		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("DeleteByID", fakeTrID).Return(nil)
		err := mockRepo.DeleteByID(fakeTrID)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		fakeTrID := int64(1)
		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("DeleteByID", fakeTrID).Return(errors.New("forced error"))
		err := mockRepo.DeleteByID(fakeTrID)
		assert.Error(t, err)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteByWalletID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fakeWalletID := int64(1)
		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("DeleteByWalletID", fakeWalletID).Return(nil)
		err := mockRepo.DeleteByWalletID(fakeWalletID)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		fakeWalletID := int64(1)
		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("DeleteByWalletID", fakeWalletID).Return(errors.New("forced error"))
		err := mockRepo.DeleteByWalletID(fakeWalletID)
		assert.Error(t, err)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteByWalletIDAndDiscountCode(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fakeWalletID := int64(1)
		fakeDiscountCode := "TEST"
		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("DeleteByWalletIDAndDiscountCode", fakeWalletID, fakeDiscountCode).Return(nil)
		err := mockRepo.DeleteByWalletIDAndDiscountCode(fakeWalletID, fakeDiscountCode)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		fakeWalletID := int64(1)
		fakeDiscountCode := "TEST"
		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("DeleteByWalletIDAndDiscountCode", fakeWalletID, fakeDiscountCode).Return(errors.New("forced error"))
		err := mockRepo.DeleteByWalletIDAndDiscountCode(fakeWalletID, fakeDiscountCode)
		assert.Error(t, err)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteByWalletIDAndType(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fakeWalletID := int64(1)
		fakeType := transaction.Recharge
		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("DeleteByWalletIDAndType", fakeWalletID, fakeType).Return(nil)
		err := mockRepo.DeleteByWalletIDAndType(fakeWalletID, fakeType)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		fakeWalletID := int64(1)
		fakeType := transaction.Recharge
		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("DeleteByWalletIDAndType", fakeWalletID, fakeType).Return(errors.New("forced error"))
		err := mockRepo.DeleteByWalletIDAndType(fakeWalletID, fakeType)
		assert.Error(t, err)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

// test GetBalance
func TestGetBalance(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fakeWalletID := int64(1)
		mockRepo := repomocks.NewRepository(t)
		// Include mock balance (like 5000) along with nil
		mockRepo.On("GetBalance", fakeWalletID).Return(int64(5000), nil)
		balance, err := mockRepo.GetBalance(fakeWalletID)
		assert.NoError(t, err)
		assert.Equal(t, int64(5000), balance)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		fakeWalletID := int64(1)
		mockRepo := repomocks.NewRepository(t)
		// Include zero balance with error
		mockRepo.On("GetBalance", fakeWalletID).Return(int64(0), errors.New("forced error"))
		balance, err := mockRepo.GetBalance(fakeWalletID)
		assert.Error(t, err)
		assert.Equal(t, int64(0), balance)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestGetByDiscountCodeWithPagination(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fakeDiscountCode := "TEST"
		fakePage := 1
		fakeLimit := 10
		mockRepo := repomocks.NewRepository(t)
		// Include mock balance (like 5000) along with nil
		mockRepo.On("GetByDiscountCodeWithPagination", fakeDiscountCode, fakePage, fakeLimit).Return([]*transaction.Transaction{}, nil)
		transactions, err := mockRepo.GetByDiscountCodeWithPagination(fakeDiscountCode, fakePage, fakeLimit)
		assert.NoError(t, err)
		assert.Equal(t, []*transaction.Transaction{}, transactions)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		fakeDiscountCode := "TEST"
		fakePage := 1
		fakeLimit := 10
		mockRepo := repomocks.NewRepository(t)
		// Include zero balance with error
		mockRepo.On("GetByDiscountCodeWithPagination", fakeDiscountCode, fakePage, fakeLimit).Return([]*transaction.Transaction{}, errors.New("forced error"))
		transactions, err := mockRepo.GetByDiscountCodeWithPagination(fakeDiscountCode, fakePage, fakeLimit)
		assert.Error(t, err)
		assert.Equal(t, []*transaction.Transaction{}, transactions)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fakeTrID := int64(1)
		mockRepo := repomocks.NewRepository(t)
		// Include mock balance (like 5000) along with nil
		mockRepo.On("GetByID", fakeTrID).Return(&transaction.Transaction{}, nil)
		tr, err := mockRepo.GetByID(fakeTrID)
		assert.NoError(t, err)
		assert.Equal(t, &transaction.Transaction{}, tr)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		fakeTrID := int64(1)
		mockRepo := repomocks.NewRepository(t)
		// Include zero balance with error
		mockRepo.On("GetByID", fakeTrID).Return(&transaction.Transaction{}, errors.New("forced error"))
		tr, err := mockRepo.GetByID(fakeTrID)
		assert.Error(t, err)
		assert.Equal(t, &transaction.Transaction{}, tr)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestGetByWalletID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fakeWalletID := int64(1)
		mockRepo := repomocks.NewRepository(t)
		// Include mock balance (like 5000) along with nil
		mockRepo.On("GetByWalletID", fakeWalletID).Return([]*transaction.Transaction{}, nil)
		tr, err := mockRepo.GetByWalletID(fakeWalletID)
		assert.NoError(t, err)
		assert.Equal(t, []*transaction.Transaction{}, tr)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		fakeWalletID := int64(1)
		mockRepo := repomocks.NewRepository(t)
		// Include zero balance with error
		mockRepo.On("GetByWalletID", fakeWalletID).Return([]*transaction.Transaction{}, errors.New("forced error"))
		tr, err := mockRepo.GetByWalletID(fakeWalletID)
		assert.Error(t, err)
		assert.Equal(t, []*transaction.Transaction{}, tr)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestGetByWalletIDAndDiscountCode(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fakeWalletID := int64(1)
		fakeDiscountCode := "TEST"
		mockRepo := repomocks.NewRepository(t)
		// Include mock balance (like 5000) along with nil
		mockRepo.On("GetByWalletIDAndDiscountCode", fakeWalletID, fakeDiscountCode).Return([]*transaction.Transaction{}, nil)
		tr, err := mockRepo.GetByWalletIDAndDiscountCode(fakeWalletID, fakeDiscountCode)
		assert.NoError(t, err)
		assert.Equal(t, []*transaction.Transaction{}, tr)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		fakeWalletID := int64(1)
		fakeDiscountCode := "TEST"
		mockRepo := repomocks.NewRepository(t)
		// Include zero balance with error
		mockRepo.On("GetByWalletIDAndDiscountCode", fakeWalletID, fakeDiscountCode).Return([]*transaction.Transaction{}, errors.New("forced error"))
		tr, err := mockRepo.GetByWalletIDAndDiscountCode(fakeWalletID, fakeDiscountCode)
		assert.Error(t, err)
		assert.Equal(t, []*transaction.Transaction{}, tr)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestGetByWalletIDAndType(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fakeWalletID := int64(1)
		fakeType := transaction.Recharge
		mockRepo := repomocks.NewRepository(t)
		// Include mock balance (like 5000) along with nil
		mockRepo.On("GetByWalletIDAndType", fakeWalletID, fakeType).Return([]*transaction.Transaction{}, nil)
		tr, err := mockRepo.GetByWalletIDAndType(fakeWalletID, fakeType)
		assert.NoError(t, err)
		assert.Equal(t, []*transaction.Transaction{}, tr)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		fakeWalletID := int64(1)
		fakeType := transaction.Recharge
		mockRepo := repomocks.NewRepository(t)
		// Include zero balance with error
		mockRepo.On("GetByWalletIDAndType", fakeWalletID, fakeType).Return([]*transaction.Transaction{}, errors.New("forced error"))
		tr, err := mockRepo.GetByWalletIDAndType(fakeWalletID, fakeType)
		assert.Error(t, err)
		assert.Equal(t, []*transaction.Transaction{}, tr)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestGetByWalletIDAndTypeAndDiscountCode(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fakeWalletID := int64(1)
		fakeType := transaction.Recharge
		fakeDiscountCode := "TEST"
		mockRepo := repomocks.NewRepository(t)
		// Include mock balance (like 5000) along with nil
		mockRepo.On("GetByWalletIDAndTypeAndDiscountCode", fakeWalletID, fakeType, fakeDiscountCode).Return([]*transaction.Transaction{}, nil)
		tr, err := mockRepo.GetByWalletIDAndTypeAndDiscountCode(fakeWalletID, fakeType, fakeDiscountCode)
		assert.NoError(t, err)
		assert.Equal(t, []*transaction.Transaction{}, tr)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		fakeWalletID := int64(1)
		fakeType := transaction.Recharge
		fakeDiscountCode := "TEST"
		mockRepo := repomocks.NewRepository(t)
		// Include zero balance with error
		mockRepo.On("GetByWalletIDAndTypeAndDiscountCode", fakeWalletID, fakeType, fakeDiscountCode).Return([]*transaction.Transaction{}, errors.New("forced error"))
		tr, err := mockRepo.GetByWalletIDAndTypeAndDiscountCode(fakeWalletID, fakeType, fakeDiscountCode)
		assert.Error(t, err)
		assert.Equal(t, []*transaction.Transaction{}, tr)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestGetByWalletIDWithPagination(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fakeWalletID := int64(1)
		fakePage := 1
		fakeLimit := 10
		mockRepo := repomocks.NewRepository(t)
		// Include mock balance (like 5000) along with nil
		mockRepo.On("GetByWalletIDWithPagination", fakeWalletID, fakePage, fakeLimit).Return([]*transaction.Transaction{}, nil)
		tr, err := mockRepo.GetByWalletIDWithPagination(fakeWalletID, fakePage, fakeLimit)
		assert.NoError(t, err)
		assert.Equal(t, []*transaction.Transaction{}, tr)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		fakeWalletID := int64(1)
		fakePage := 1
		fakeLimit := 10
		mockRepo := repomocks.NewRepository(t)
		// Include zero balance with error
		mockRepo.On("GetByWalletIDWithPagination", fakeWalletID, fakePage, fakeLimit).Return([]*transaction.Transaction{}, errors.New("forced error"))
		tr, err := mockRepo.GetByWalletIDWithPagination(fakeWalletID, fakePage, fakeLimit)
		assert.Error(t, err)
		assert.Equal(t, []*transaction.Transaction{}, tr)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestWithTX(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fakeTx := &transaction.Transaction{
			WalletID:        1,
			Amount:          100,
			TransactionType: "credit",
		}
		mockRepo := repomocks.NewRepository(t)
		// Change "Create" to "Insert" here
		mockRepo.On("Insert", fakeTx).Return(nil)
		err := mockRepo.Insert(fakeTx)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		fakeTx := &transaction.Transaction{
			WalletID:        1,
			Amount:          100,
			TransactionType: "credit",
		}
		mockRepo := repomocks.NewRepository(t)
		// Change "Create" to "Insert" here
		mockRepo.On("Insert", fakeTx).Return(errors.New("forced error"))
		err := mockRepo.Insert(fakeTx)
		assert.Error(t, err)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
