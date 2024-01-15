package member_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	repomocks "wallet/mocks/repomocks/member"
	"wallet/storage/member"
)

func TestCreate(t *testing.T) {
	t.Run("create member", func(t *testing.T) {
		fakeMember := &member.Member{
			FirstName: "test",
			LastName:  "test",
			Email:     "a@b.com",
			Phone:     "+989123456789",
		}
		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("Create", fakeMember).Return(nil)
		err := mockRepo.Create(fakeMember)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		fakeMember := &member.Member{
			FirstName: "test",
			LastName:  "test",
			Email:     "a@b.com",
			Phone:     "+989123456789",
		}
		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("Create", fakeMember).Return(errors.New("forced error"))
		err := mockRepo.Create(fakeMember)
		assert.Error(t, err)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestGetAllByPage(t *testing.T) {
	t.Run("get all members", func(t *testing.T) {
		fakeMember := &member.Member{
			FirstName: "test",
			LastName:  "test",
			Email:     "a@b.com",
			Phone:     "+989123456789",
		}
		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("GetAllByPage", 10, 0, false).Return([]*member.Member{fakeMember}, 1, nil)
		members, count, err := mockRepo.GetAllByPage(10, 0, false)
		assert.NoError(t, err)
		assert.Equal(t, 1, count)
		assert.Equal(t, fakeMember, members[0])
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("GetAllByPage", 10, 0, false).Return(nil, 0, errors.New("forced error"))
		members, count, err := mockRepo.GetAllByPage(10, 0, false)
		assert.Error(t, err)
		assert.Equal(t, 0, count)
		assert.Nil(t, members)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("get member by id", func(t *testing.T) {
		fakeMember := &member.Member{
			FirstName: "test",
			LastName:  "test",
			Email:     "a@b.com",
			Phone:     "+989123456789",
		}
		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("GetById", int64(1)).Return(fakeMember, nil)
		member, err := mockRepo.GetById(1)
		assert.NoError(t, err)
		assert.Equal(t, fakeMember, member)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("GetById", int64(1)).Return(nil, errors.New("forced error"))
		member, err := mockRepo.GetById(1)
		assert.Error(t, err)
		assert.Nil(t, member)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestGetByPhone(t *testing.T) {
	t.Run("get member by phone", func(t *testing.T) {
		fakeMember := &member.Member{
			FirstName: "test",
			LastName:  "test",
			Email:     "a@b.com",
			Phone:     "+989123456789",
		}
		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("GetByPhone", "+989123456789").Return(fakeMember, nil)
		member, err := mockRepo.GetByPhone("+989123456789")
		assert.NoError(t, err)
		assert.Equal(t, fakeMember, member)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("GetByPhone", "+989123456789").Return(nil, errors.New("forced error"))
		member, err := mockRepo.GetByPhone("+989123456789")
		assert.Error(t, err)
		assert.Nil(t, member)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update member", func(t *testing.T) {
		fakeMember := &member.Member{
			FirstName: "test",
			LastName:  "test",
			Email:     "a@b.com",
			Phone:     "+989123456789",
		}
		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("Update", fakeMember).Return(nil)
		err := mockRepo.Update(fakeMember)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		fakeMember := &member.Member{
			FirstName: "test",
			LastName:  "test",
			Email:     "a@b.com",
			Phone:     "+989123456789",
		}
		mockRepo := repomocks.NewRepository(t)
		mockRepo.On("Update", fakeMember).Return(errors.New("forced error"))
		err := mockRepo.Update(fakeMember)
		assert.Error(t, err)
		assert.Equal(t, "forced error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
