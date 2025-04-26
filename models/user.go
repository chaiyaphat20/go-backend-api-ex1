package models

import (
	"time"

	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
)

// table name = users
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Fullname  string    `json:"fullname" gorm:"type:varchar(255);not null"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password  string    `json:"-" gorm:"type:varchar(255);not null"`     //ใช้ - ไม่ต้องแสดง password ทั้ง response
	IsAdmin   bool      `json:"is_admin" gorm:"type:bool;default:false"` //ถ้า db เป็น is_active จะใช้ column:is_active
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
