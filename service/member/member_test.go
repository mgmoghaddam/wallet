package member_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	repomocks "wallet/mocks/repomocks/member"
	"wallet/service/member"
)

var MockMember = &member.DTO{
	ID:        1,
	Email:     "a@b.com",
	FirstName: "a",
	LastName:  "b",
	Phone:     "+989123456789",
}

var MockRequest = &member.CreateRequest{
	Email:     "a@b.com",
	FirstName: "a",
	LastName:  "b",
	Phone:     "+989123456789",
}

// Test case for successful usage of the `Create` function.
func TestCreate(t *testing.T) {
	mockUseCase := repomocks.NewUseCase(t)

	mockUseCase.On("Create", MockRequest).Return(MockMember, nil)

	result, err := mockUseCase.Create(MockRequest)

	assert.Nil(t, err)
	assert.Equal(t, MockMember, result)
	mockUseCase.AssertCalled(t, "Create", MockRequest)
}

// Test case for unsuccessful usage of the `Create` function.
func TestCreateError(t *testing.T) {
	mockUseCase := repomocks.NewUseCase(t)

	mockUseCase.On("Create", MockRequest).Return(nil, errors.New("Failed to create member"))

	result, err := mockUseCase.Create(MockRequest)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "Failed to create member", err.Error())
	mockUseCase.AssertCalled(t, "Create", MockRequest)
}

// Test case for successful usage of the `GetById` function.
func TestGetById(t *testing.T) {
	mockUseCase := repomocks.NewUseCase(t)

	mockUseCase.On("GetById", int64(1)).Return(MockMember, nil)

	result, err := mockUseCase.GetById(int64(1))

	assert.Nil(t, err)
	assert.Equal(t, MockMember, result)
	mockUseCase.AssertCalled(t, "GetById", int64(1))
}

// Test case for unsuccessful usage of the `GetById` function.
func TestGetByIdError(t *testing.T) {
	mockUseCase := repomocks.NewUseCase(t)

	mockUseCase.On("GetById", int64(1)).Return(nil, errors.New("Failed to get member"))

	result, err := mockUseCase.GetById(int64(1))

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "Failed to get member", err.Error())
	mockUseCase.AssertCalled(t, "GetById", int64(1))
}

// Test case for successful usage of the `GetByPhone` function.
func TestGetByPhone(t *testing.T) {
	mockUseCase := repomocks.NewUseCase(t)

	mockUseCase.On("GetByPhone", "+989123456789").Return(MockMember, nil)

	result, err := mockUseCase.GetByPhone("+989123456789")

	assert.Nil(t, err)
	assert.Equal(t, MockMember, result)
	mockUseCase.AssertCalled(t, "GetByPhone", "+989123456789")
}

// Test case for unsuccessful usage of the `GetByPhone` function.
func TestGetByPhoneError(t *testing.T) {
	mockUseCase := repomocks.NewUseCase(t)

	mockUseCase.On("GetByPhone", "+989123456789").Return(nil, errors.New("Failed to get member"))

	result, err := mockUseCase.GetByPhone("+989123456789")

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "Failed to get member", err.Error())
	mockUseCase.AssertCalled(t, "GetByPhone", "+989123456789")
}

// Test case for successful usage of the `GetMembersByGiftCode` function.
func TestGetMembersByGiftCode(t *testing.T) {
	mockUseCase := repomocks.NewUseCase(t)

	mockUseCase.On("GetMembersByGiftCode", mock.Anything, mock.Anything, mock.Anything).Return([]*member.DTO{MockMember}, nil)

	result, err := mockUseCase.GetMembersByGiftCode("TEST", 1, 1)

	assert.Nil(t, err)
	assert.Equal(t, []*member.DTO{MockMember}, result)
	mockUseCase.AssertCalled(t, "GetMembersByGiftCode", "TEST", 1, 1)
}

// Test case for unsuccessful usage of the `GetMembersByGiftCode` function.
func TestGetMembersByGiftCodeError(t *testing.T) {
	mockUseCase := repomocks.NewUseCase(t)

	mockUseCase.On("GetMembersByGiftCode", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("Failed to get members"))

	result, err := mockUseCase.GetMembersByGiftCode("TEST", 1, 1)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "Failed to get members", err.Error())
	mockUseCase.AssertCalled(t, "GetMembersByGiftCode", "TEST", 1, 1)
}

// Test case for successful usage of the `Update` function.
func TestUpdate(t *testing.T) {
	mockUseCase := repomocks.NewUseCase(t)

	mockUseCase.On("Update", MockMember).Return(MockMember, nil)

	result, err := mockUseCase.Update(MockMember)

	assert.Nil(t, err)
	assert.Equal(t, MockMember, result)
	mockUseCase.AssertCalled(t, "Update", MockMember)
}

// Test case for unsuccessful usage of the `Update` function.
func TestUpdateError(t *testing.T) {
	mockUseCase := repomocks.NewUseCase(t)

	mockUseCase.On("Update", MockMember).Return(nil, errors.New("Failed to update member"))

	result, err := mockUseCase.Update(MockMember)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "Failed to update member", err.Error())
	mockUseCase.AssertCalled(t, "Update", MockMember)
}
