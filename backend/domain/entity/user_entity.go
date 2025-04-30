package entity

type UserEntity struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Password string `gorm:"not null"`
	Username string `gorm:"not null"`
}

func (UserEntity) TableName() string {
	return "users"
}
