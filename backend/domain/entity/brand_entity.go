package entity

import (
	"time"

	"gorm.io/gorm"
)

type BrandEntity struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"string(100)"`
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// sesuain nama table dengan migration
func (BrandEntity) TableName() string {
	return "brands"
}
