package configs

import (
	"fmt"
	"os"

	"example.com/gin-backend-api/models"
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
		fmt.Println(err.Error())
		panic(err)
	}
	fmt.Println("เชื่อมต่อ Database success")

	//Migration
	// db.Migrator().DropTable(&models.User{}) //ลบ table
	db.AutoMigrate(&models.User{}, &models.Blog{})
	DB = db
}
