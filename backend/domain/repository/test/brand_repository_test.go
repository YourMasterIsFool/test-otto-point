package repository_test

import (
	"backend/domain/entity"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockBrandRepository adalah implementasi mock dari BrandRepository
type MockBrandRepository struct {
	mock.Mock
}

func (m *MockBrandRepository) Save(schema entity.BrandEntity) (*entity.BrandEntity, error) {
	args := m.Called(schema)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.BrandEntity), args.Error(1)
}

func (m *MockBrandRepository) Detail(id uint) (*entity.BrandEntity, error) {
	args := m.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.BrandEntity), args.Error(1)
}

func TestBrandRepository(t *testing.T) {
	// Create mock repository
	mockRepo := new(MockBrandRepository)

	// Test case for Save success
	t.Run("Save_Success", func(t *testing.T) {
		newUser := entity.BrandEntity{Name: "test_brand"}
		savedUser := &entity.BrandEntity{ID: 2, Name: "test_brand"}

		mockRepo.On("Save", newUser).Return(savedUser, nil)

		result, err := mockRepo.Save(newUser)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "test_brand", result.Name)
		assert.Equal(t, uint(2), result.ID)

		mockRepo.AssertExpectations(t)
	})

	// Test case for Save failure
	t.Run("Save_Failure", func(t *testing.T) {
		invalidUser := entity.BrandEntity{Name: ""}
		mockRepo.On("Save", invalidUser).Return(nil, errors.New("Name cannot be empty"))

		result, err := mockRepo.Save(invalidUser)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "Name cannot be empty")

		mockRepo.AssertExpectations(t)
	})

	// Test case for Detail success
	t.Run("Detail_Success", func(t *testing.T) {
		mockUser := &entity.BrandEntity{ID: 1, Name: "test@example.com"}
		mockRepo.On("Detail", uint(1)).Return(mockUser, nil)

		result, err := mockRepo.Detail(1)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, uint(1), result.ID)
		assert.Equal(t, "test@example.com", result.Name)

		mockRepo.AssertExpectations(t)
	})

	// Test case for Detail failure
	t.Run("Detail_NotFound", func(t *testing.T) {
		mockRepo.On("Detail", uint(999)).Return(nil, errors.New("user not found"))

		result, err := mockRepo.Detail(999)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "user not found")

		mockRepo.AssertExpectations(t)
	})

}
