package models

type Blog struct {
	ID     uint   `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Topic  string `json:"topic"`
	UserID uint   `json:"-"`                                                         //KF
	User   User   `json:"-" gorm:"foreignKey:UserID;references:ID;onDelete:CASCADE"` // Define foreign key relationship
}
