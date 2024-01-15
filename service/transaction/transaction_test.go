package transaction_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	repomocks "wallet/mocks/repomocks/transaction"
	transaction "wallet/service/transaction"
)

var MockTransaction = &transaction.DTO{
	WalletID:        1,
	Amount:          100,
	TransactionType: transaction.Recharge,
}
var MockRequest = &transaction.CreateRequest{
	WalletID:        1,
	Amount:          100,
	TransactionType: transaction.Recharge,
}

// Test case for successful usage of the `Create` function.
func TestCreate_Success(t *testing.T) {
	mockUseCase := repomocks.NewUseCase(t)

	// Setup expectations
	mockUseCase.On("Create", mock.Anything).Return(MockTransaction, nil)

	// Call the method
	result, err := mockUseCase.Create(MockRequest)

	// Assert Expectations
	assert.Nil(t, err)
	assert.Equal(t, MockTransaction, result)

	// The `Create` function should have been called with our MockRequest
	mockUseCase.AssertCalled(t, "Create", MockRequest)
}

// Test case for unsuccessful usage of the `Create` function.
func TestCreate_Failure(t *testing.T) {
	mockUseCase := repomocks.NewUseCase(t)

	// Setup expectations
	mockUseCase.On("Create", mock.Anything).Return(nil, errors.New("Failed to create transaction"))

	// Call the method
	result, err := mockUseCase.Create(MockRequest)

	// Assert expectations
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Failed to create transaction")

	// The `Create` function should have been called with our MockRequest
	mockUseCase.AssertCalled(t, "Create", MockRequest)
}

// Test case for successful usage of the `Delete` function.
func TestDelete_Success(t *testing.T) {
	id := int64(1) // defines id as int64
	var mockUseCase = &repomocks.UseCase{}
	mockUseCase.On("Delete", id).Return(nil)

	err := mockUseCase.Delete(id)
	assert.NoError(t, err)
	mockUseCase.AssertExpectations(t)
}

// Test case for unsuccessful usage of the `Delete` function.
func TestDelete_Failure(t *testing.T) {
	id := int64(1) // defines id as int64
	var mockUseCase = &repomocks.UseCase{}
	mockUseCase.On("Delete", id).Return(errors.New("Failed to delete transaction"))

	err := mockUseCase.Delete(id)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "Failed to delete transaction")
	mockUseCase.AssertExpectations(t)
}

// Test case for successful usage of the `DeleteByWalletID` function.
func TestDeleteByWalletID_Success(t *testing.T) {
	id := int64(1) // defines id as int64
	var mockUseCase = &repomocks.UseCase{}
	mockUseCase.On("DeleteByWalletID", id).Return(nil)

	err := mockUseCase.DeleteByWalletID(id)
	assert.NoError(t, err)
	mockUseCase.AssertExpectations(t)
}

// Test case for unsuccessful usage of the `DeleteByWalletID` function.
func TestDeleteByWalletID_Failure(t *testing.T) {
	id := int64(1) // defines id as int64
	var mockUseCase = &repomocks.UseCase{}
	mockUseCase.On("DeleteByWalletID", id).Return(errors.New("Failed to delete transaction"))

	err := mockUseCase.DeleteByWalletID(id)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "Failed to delete transaction")
	mockUseCase.AssertExpectations(t)
}

// Test case for successful usage of the `GetBalance` function.
func TestGetBalance_Success(t *testing.T) {
	id := int64(1) // defines id as int64
	var mockUseCase = &repomocks.UseCase{}
	MockTransaction := int64(1000) // Mock balance value
	mockUseCase.On("GetBalance", id).Return(MockTransaction, nil)
	result, err := mockUseCase.GetBalance(id)
	assert.NoError(t, err)
	assert.Equal(t, MockTransaction, result)
	mockUseCase.AssertExpectations(t)
}

// Test case for unsuccessful usage of the `GetBalance` function.
func TestGetBalance_Failure(t *testing.T) {
	id := int64(1) // defines id as int64
	var mockUseCase = &repomocks.UseCase{}
	mockUseCase.On("GetBalance", id).Return(int64(0), errors.New("Failed to get balance")) // Return 0 instead of nil
	result, err := mockUseCase.GetBalance(id)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "Failed to get balance")
	assert.Equal(t, int64(0), result) // Assert that result is 0
	mockUseCase.AssertExpectations(t)
}

// Test case for successful usage of the `GetByDiscountCodeWithPagination` function.
func TestGetByDiscountCodeWithPagination_Success(t *testing.T) {
	var mockUseCase = &repomocks.UseCase{}
	MockTransaction := []*transaction.DTO{MockTransaction} // Mock transaction value
	mockUseCase.On("GetByDiscountCodeWithPagination", mock.Anything, mock.Anything, mock.Anything).Return(MockTransaction, nil)
	result, err := mockUseCase.GetByDiscountCodeWithPagination("TEST", 1, 1)
	assert.NoError(t, err)
	assert.Equal(t, MockTransaction, result)
	mockUseCase.AssertExpectations(t)
}

// Test case for unsuccessful usage of the `GetByDiscountCodeWithPagination` function.
func TestGetByDiscountCodeWithPagination_Failure(t *testing.T) {
	var mockUseCase = &repomocks.UseCase{}
	mockUseCase.On("GetByDiscountCodeWithPagination", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("Failed to get transaction")) // Return nil instead of MockTransaction
	result, err := mockUseCase.GetByDiscountCodeWithPagination("TEST", 1, 1)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "Failed to get transaction")
	assert.Equal(t, []*transaction.DTO(nil), result) // Assert that result is nil
	mockUseCase.AssertExpectations(t)
}

// Test case for successful usage of the `GetByID` function.
func TestGetByID_Success(t *testing.T) {
	id := int64(1) // defines id as int64
	var mockUseCase = &repomocks.UseCase{}
	mockUseCase.On("GetByID", id).Return(MockTransaction, nil)
	result, err := mockUseCase.GetByID(id)
	assert.NoError(t, err)
	assert.Equal(t, MockTransaction, result)
	mockUseCase.AssertExpectations(t)
}

// Test case for unsuccessful usage of the `GetByID` function.
func TestGetByID_Failure(t *testing.T) {
	id := int64(1) // defines id as int64
	var mockUseCase = &repomocks.UseCase{}
	mockUseCase.On("GetByID", id).Return(nil, errors.New("Failed to get transaction"))
	result, err := mockUseCase.GetByID(id)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "Failed to get transaction")
	assert.Equal(t, (*transaction.DTO)(nil), result) // Assert that result is nil
	mockUseCase.AssertExpectations(t)
}

// Test case for successful usage of the `GetByWalletID` function.
func TestGetByWalletID_Success(t *testing.T) {
	id := int64(1) // defines id as int64
	var mockUseCase = &repomocks.UseCase{}
	MockTransaction := []*transaction.DTO{MockTransaction} // Mock transaction value
	mockUseCase.On("GetByWalletID", id).Return(MockTransaction, nil)
	result, err := mockUseCase.GetByWalletID(id)
	assert.NoError(t, err)
	assert.Equal(t, MockTransaction, result)
	mockUseCase.AssertExpectations(t)
}

// Test case for unsuccessful usage of the `GetByWalletID` function.
func TestGetByWalletID_Failure(t *testing.T) {
	id := int64(1) // defines id as int64
	var mockUseCase = &repomocks.UseCase{}
	mockUseCase.On("GetByWalletID", id).Return(nil, errors.New("Failed to get transaction")) // Return nil instead of MockTransaction
	result, err := mockUseCase.GetByWalletID(id)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "Failed to get transaction")
	assert.Equal(t, []*transaction.DTO(nil), result) // Assert that result is nil
	mockUseCase.AssertExpectations(t)
}

// Test case for successful usage of the `GetByWalletIDAndDiscountCode` function.
func TestGetByWalletIDAndDiscountCode_Success(t *testing.T) {
	id := int64(1) // defines id as int64
	var mockUseCase = &repomocks.UseCase{}
	MockTransactions := []*transaction.DTO{MockTransaction} // Mock transaction value
	mockUseCase.On("GetByWalletIDAndDiscountCode", id, "TEST").Return(MockTransactions, nil)
	result, err := mockUseCase.GetByWalletIDAndDiscountCode(id, "TEST")
	assert.NoError(t, err)
	assert.Equal(t, MockTransactions, result)
	mockUseCase.AssertExpectations(t)
}

// Test case for unsuccessful usage of the `GetByWalletIDAndDiscountCode` function.
func TestGetByWalletIDAndDiscountCode_Failure(t *testing.T) {
	id := int64(1) // defines id as int64
	var mockUseCase = &repomocks.UseCase{}
	mockUseCase.On("GetByWalletIDAndDiscountCode", id, "TEST").Return(nil, errors.New("Failed to get transaction"))
	result, err := mockUseCase.GetByWalletIDAndDiscountCode(id, "TEST")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "Failed to get transaction")
	assert.Equal(t, ([]*transaction.DTO)(nil), result) // Assert that result is nil
	mockUseCase.AssertExpectations(t)
}

// Test case for successful usage of the `GetByWalletIDAndType` function.
func TestGetByWalletIDAndType_Success(t *testing.T) {
	id := int64(1) // defines id as int64
	var mockUseCase = &repomocks.UseCase{}
	MockTransaction := []*transaction.DTO{MockTransaction} // Mock transaction value
	mockUseCase.On("GetByWalletIDAndType", id, transaction.Recharge).Return(MockTransaction, nil)
	result, err := mockUseCase.GetByWalletIDAndType(id, transaction.Recharge)
	assert.NoError(t, err)
	assert.Equal(t, MockTransaction, result)
	mockUseCase.AssertExpectations(t)
}

// Test case for unsuccessful usage of the `GetByWalletIDAndType` function.
func TestGetByWalletIDAndType_Failure(t *testing.T) {
	id := int64(1) // defines id as int64
	var mockUseCase = &repomocks.UseCase{}
	mockUseCase.On("GetByWalletIDAndType", id, transaction.Recharge).Return(nil, errors.New("Failed to get transaction")) // Return nil instead of MockTransaction
	result, err := mockUseCase.GetByWalletIDAndType(id, transaction.Recharge)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "Failed to get transaction")
	assert.Equal(t, []*transaction.DTO(nil), result) // Assert that result is nil
	mockUseCase.AssertExpectations(t)
}

// Test case for successful usage of the `GetByWalletIDAndTypeAndDiscountCode` function.
func TestGetByWalletIDAndTypeAndDiscountCode_Success(t *testing.T) {
	id := int64(1) // defines id as int64
	var mockUseCase = &repomocks.UseCase{}
	MockTransaction := []*transaction.DTO{MockTransaction} // Mock transaction value
	mockUseCase.On("GetByWalletIDAndTypeAndDiscountCode", id, transaction.Recharge, "TEST").Return(MockTransaction, nil)
	result, err := mockUseCase.GetByWalletIDAndTypeAndDiscountCode(id, transaction.Recharge, "TEST")
	assert.NoError(t, err)
	assert.Equal(t, MockTransaction, result)
	mockUseCase.AssertExpectations(t)
}

// Test case for unsuccessful usage of the `GetByWalletIDAndTypeAndDiscountCode` function.
func TestGetByWalletIDAndTypeAndDiscountCode_Failure(t *testing.T) {
	id := int64(1) // defines id as int64
	var mockUseCase = &repomocks.UseCase{}
	mockUseCase.On("GetByWalletIDAndTypeAndDiscountCode", id, transaction.Recharge, "TEST").Return(nil, errors.New("Failed to get transaction")) // Return nil instead of MockTransaction
	result, err := mockUseCase.GetByWalletIDAndTypeAndDiscountCode(id, transaction.Recharge, "TEST")
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "Failed to get transaction")
	assert.Equal(t, []*transaction.DTO(nil), result) // Assert that result is nil
	mockUseCase.AssertExpectations(t)
}

// Test case for successful usage of the `GetByWalletIDWithPagination` function.
func TestGetByWalletIDWithPagination_Success(t *testing.T) {
	var mockUseCase = &repomocks.UseCase{}
	MockTransaction := []*transaction.DTO{MockTransaction} // Mock transaction value
	mockUseCase.On("GetByWalletIDWithPagination", mock.Anything, mock.Anything, mock.Anything).Return(MockTransaction, nil)
	result, err := mockUseCase.GetByWalletIDWithPagination(1, 1, 1)
	assert.NoError(t, err)
	assert.Equal(t, MockTransaction, result)
	mockUseCase.AssertExpectations(t)
}

// Test case for unsuccessful usage of the `GetByWalletIDWithPagination` function.
func TestGetByWalletIDWithPagination_Failure(t *testing.T) {
	var mockUseCase = &repomocks.UseCase{}
	mockUseCase.On("GetByWalletIDWithPagination", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("Failed to get transaction")) // Return nil instead of MockTransaction
	result, err := mockUseCase.GetByWalletIDWithPagination(1, 1, 1)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "Failed to get transaction")
	assert.Equal(t, []*transaction.DTO(nil), result) // Assert that result is nil
	mockUseCase.AssertExpectations(t)
}
