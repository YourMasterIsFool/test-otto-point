package repository_test

import (
	"backend/domain/entity"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAuthRepository adalah implementasi mock dari AuthRepository
type MockAuthRepository struct {
	mock.Mock
}

func (m *MockAuthRepository) FindByUsername(username string) (*entity.UserEntity, error) {
	args := m.Called(username)

	// Handle nil return value properly
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.UserEntity), args.Error(1)
}

func (m *MockAuthRepository) Save(schema entity.UserEntity) (*entity.UserEntity, error) {
	args := m.Called(schema)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.UserEntity), args.Error(1)
}

func (m *MockAuthRepository) Detail(id uint) (*entity.UserEntity, error) {
	args := m.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.UserEntity), args.Error(1)
}

func TestAuthRepository(t *testing.T) {
	// Create mock repository
	mockRepo := new(MockAuthRepository)

	// Test case for FindByUsername success
	t.Run("FindByUsername_Success", func(t *testing.T) {
		mockUser := &entity.UserEntity{ID: 1, Username: "test@example.com"}
		mockRepo.On("FindByUsername", "test@example.com").Return(mockUser, nil)

		result, err := mockRepo.FindByUsername("test@example.com")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "test@example.com", result.Username)
		assert.Equal(t, uint(1), result.ID)

		mockRepo.AssertExpectations(t)
	})

	// Test case for FindByUsername failure
	t.Run("FindByUsername_NotFound", func(t *testing.T) {
		mockRepo.On("FindByUsername", "nonexistent@example.com").Return(nil, errors.New("user not found"))

		result, err := mockRepo.FindByUsername("nonexistent@example.com")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "user not found")

		mockRepo.AssertExpectations(t)
	})

	// Test case for Save success
	t.Run("Save_Success", func(t *testing.T) {
		newUser := entity.UserEntity{Username: "newuser@example.com"}
		savedUser := &entity.UserEntity{ID: 2, Username: "newuser@example.com"}

		mockRepo.On("Save", newUser).Return(savedUser, nil)

		result, err := mockRepo.Save(newUser)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "newuser@example.com", result.Username)
		assert.Equal(t, uint(2), result.ID)

		mockRepo.AssertExpectations(t)
	})

	// Test case for Save failure
	t.Run("Save_Failure", func(t *testing.T) {
		invalidUser := entity.UserEntity{Username: ""}
		mockRepo.On("Save", invalidUser).Return(nil, errors.New("username cannot be empty"))

		result, err := mockRepo.Save(invalidUser)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "username cannot be empty")

		mockRepo.AssertExpectations(t)
	})

	// Test case for Detail success
	t.Run("Detail_Success", func(t *testing.T) {
		mockUser := &entity.UserEntity{ID: 1, Username: "test@example.com"}
		mockRepo.On("Detail", uint(1)).Return(mockUser, nil)

		result, err := mockRepo.Detail(1)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, uint(1), result.ID)
		assert.Equal(t, "test@example.com", result.Username)

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
