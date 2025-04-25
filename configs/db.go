package configs

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := os.Getenv("DATABASE_DSN") //dsn = data source name
	fmt.Println("DSN: ", os.Getenv("DATABASE_DSN"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("ไม่สามารถเชื่อมต่อ Database ได้")
	} else {
		fmt.Println("เชื่อมต่อ Database success")
	}

	DB = db
}
