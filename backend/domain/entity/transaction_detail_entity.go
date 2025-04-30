package entity

type TransactionDetailEntity struct {
	ID            uint    `gorm:"primaryKey"`
	Point         float64 `gorm:"not null"`
	BrandName     string  `gorm:"string(100)"`
	BrandId       uint    `gorm:"not null"`
	VoucherId     uint    `gorm:"not null"`
	TransactionId uint    `gorm:"not null"`
}

func (TransactionDetailEntity) TableName() string {
	return "transaction_details"
}
