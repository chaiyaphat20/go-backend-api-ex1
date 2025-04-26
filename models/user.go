package models

import (
	"time"

	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
)

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

// method BeforeCreate ของ struct Users  หรือ hook
// https://gorm.io/docs/hooks.html
// begin transaction
// BeforeSave
// BeforeCreate
// // save before associations
// // insert into database
// // save after associations
// AfterCreate
// AfterSave
// // commit or rollback transaction

func (user *User) BeforeCreate(db *gorm.DB) error {
	user.Password = hashPassword(user.Password)
	return nil
}

func hashPassword(password string) string {
	argon := argon2.DefaultConfig()
	encoded, _ := argon.HashEncoded([]byte(password))
	return string(encoded)
}
