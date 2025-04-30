package entity

type TransactionEntity struct {
	ID         uint    `gorm:"primaryKey"`
	TotalPoint float64 `gorm:"not null"`
	BrandName  string  `gorm:"string(100)"`
	BrandId    uint    `gorm:"not null"`
	UserId     uint    `gorm:"not null"`
	// User               UserEntity                `gorm:"foreignKey:UserId;references:ID"`
	Brand              BrandEntity               `gorm:"foreignKey:BrandId;references:ID"`
	TransactionDetails []TransactionDetailEntity `gorm:"foreignKey:TransactionId;references:ID"`
}

func (TransactionEntity) TableName() string {
	return "transactions"
}
