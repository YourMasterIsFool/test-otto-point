package entity

import (
	"time"

	"gorm.io/gorm"
)

type VoucherEntity struct {
	ID        uint    `gorm:"primaryKey"`
	BrandId   uint    `gorm:"not null"`
	Point     float64 `gorm:"not null"`
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// sesuain nama table dengan migration
func (VoucherEntity) TableName() string {
	return "vouchers"
}
