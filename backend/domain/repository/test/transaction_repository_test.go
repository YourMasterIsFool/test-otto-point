package repository_test

import (
	"backend/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTransactionRepository adalah mock dari TransactionRepository
type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) Save(e *entity.TransactionEntity) (*entity.TransactionEntity, error) {
	args := m.Called(e)
	return args.Get(0).(*entity.TransactionEntity), args.Error(1)
}

func (m *MockTransactionRepository) Detail(id uint) (*entity.TransactionEntity, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.TransactionEntity), args.Error(1)
}

func TestTransactionRepository(t *testing.T) {
	mockRepo := new(MockTransactionRepository)

	t.Run("Save", func(t *testing.T) {
		input := &entity.TransactionEntity{
			BrandName: "testbrand",
			BrandId:   1,
			UserId:    1,
			ID:        1,
		}
		mockRepo.On("Save", input).Return(input, nil)

		result, err := mockRepo.Save(input)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), result.ID)
		assert.Equal(t, uint(1), result.BrandId)
		assert.Equal(t, uint(1), result.UserId)

		mockRepo.AssertExpectations(t)
	})

	t.Run("Detail", func(t *testing.T) {
		expected := &entity.TransactionEntity{
			BrandName: "testbrand",
			BrandId:   1,
			UserId:    1,
			ID:        1,
		}
		mockRepo.On("Detail", uint(1)).Return(expected, nil)

		result, err := mockRepo.Detail(1)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), result.ID)
		assert.Equal(t, uint(1), result.BrandId)
		assert.Equal(t, uint(1), result.UserId)

		mockRepo.AssertExpectations(t)
	})
}
