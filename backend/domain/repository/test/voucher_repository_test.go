package repository_test

import (
	"backend/domain/entity"
	voucherdto "backend/dtos/voucher"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockVoucherRepository struct {
	mock.Mock
}

func (m *MockVoucherRepository) Save(dto voucherdto.CreateVoucherDto) (*entity.VoucherEntity, error) {
	args := m.Called(dto)
	return args.Get(0).(*entity.VoucherEntity), args.Error(1)
}

func (m *MockVoucherRepository) Detail(id uint) (*entity.VoucherEntity, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.VoucherEntity), args.Error(1)
}

func (m *MockVoucherRepository) FindByBrandId(brandID uint) ([]*entity.VoucherEntity, error) {
	args := m.Called(brandID)
	return args.Get(0).([]*entity.VoucherEntity), args.Error(1)
}

func (m *MockVoucherRepository) FindByMultipleId(multipleID []uint) ([]*entity.VoucherEntity, error) {
	args := m.Called(multipleID)
	return args.Get(0).([]*entity.VoucherEntity), args.Error(1)
}

func TestVoucherRepository(t *testing.T) {
	mockRepo := new(MockVoucherRepository)

	t.Run("Save", func(t *testing.T) {
		input := voucherdto.CreateVoucherDto{BrandId: 1}
		expected := &entity.VoucherEntity{ID: 1, BrandId: 1}

		mockRepo.On("Save", input).Return(expected, nil)

		result, err := mockRepo.Save(input)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)

		mockRepo.AssertExpectations(t)
	})

	t.Run("Detail", func(t *testing.T) {
		expected := &entity.VoucherEntity{ID: 1, BrandId: 1}

		mockRepo.On("Detail", uint(1)).Return(expected, nil)

		result, err := mockRepo.Detail(1)

		assert.NoError(t, err)
		assert.Equal(t, expected.ID, result.ID)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FindByBrandId", func(t *testing.T) {
		expected := []*entity.VoucherEntity{
			{ID: 1, BrandId: 1, Point: 2000},
			{ID: 2, BrandId: 2, Point: 1000},
		}

		mockRepo.On("FindByBrandId", uint(99)).Return(expected, nil)

		result, err := mockRepo.FindByBrandId(99)

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, 1, result[0].BrandId)

		mockRepo.AssertExpectations(t)
	})

	t.Run("FindByMultipleId", func(t *testing.T) {
		ids := []uint{100, 200}
		expected := []*entity.VoucherEntity{
			{ID: 100, BrandId: 1},
			{ID: 200, BrandId: 2},
		}

		mockRepo.On("FindByMultipleId", mock.MatchedBy(func(arg []uint) bool {
			return len(arg) == 2 && arg[0] == 100 && arg[1] == 200
		})).Return(expected, nil)

		result, err := mockRepo.FindByMultipleId(ids)

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, uint(100), result[0].ID)

		mockRepo.AssertExpectations(t)
	})
}
