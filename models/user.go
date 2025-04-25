package models

import "time"

// table name = users
type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;not null"`
	FullName  string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);not null;unique"`
	Password  string `gorm:"type:varchar(255);not null"`
	IsAdmin   bool   `gorm:"type:bool;default:false"` //ถ้า db เป็น is_active จะใช้ column:is_active
	CreatedAt time.Time
	UpdatedAt time.Time
}
