package usercontroller

import (
	"errors"
	"net/http"

	"example.com/gin-backend-api/configs"
	"example.com/gin-backend-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "users",
	})
}

func Register(c *gin.Context) {
	// ดึงข้อมูล JSON จาก request body
	var input InputRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		// ถ้า body ไม่ถูกต้อง หรือ validation ไม่ผ่าน
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบว่า user มีอยู่ในระบบแล้วหรือยัง
	var existingUser models.User
	if err := configs.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		// ถ้า err == nil แปลว่ามี user นี้อยู่แล้ว
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// ถ้าเจอ error ที่ไม่ใช่ ErrRecordNotFound แปลว่ามีปัญหาในการ query
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	//สร้าง model
	user := models.User{
		FullName: input.Fullname,
		Email:    input.Email,
		Password: string(input.Password),
	}

	//create จริง
	result := configs.DB.Debug().Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	// ส่ง response กลับ
	c.JSON(http.StatusCreated, gin.H{
		"message": "Register success",
		"data": gin.H{
			"id":       user.ID,
			"fullname": user.FullName,
			"email":    user.Email,
			// อย่า return password กลับไป!
		},
	})
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "login",
	})
}

func GetById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"data": id,
	})
}

func SearchByFullname(c *gin.Context) {
	fullname := c.Query("fullname")
	c.JSON(200, gin.H{
		"data": fullname,
	})
}
