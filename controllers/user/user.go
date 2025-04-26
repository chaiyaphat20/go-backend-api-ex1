package usercontroller

import (
	"net/http"

	"example.com/gin-backend-api/configs"
	"example.com/gin-backend-api/models"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	var users []models.User

	// .Find(&users) เป็น method ที่รับ pointer ของ slice (&users) ไป แล้วมัน mutate หรือเปลี่ยนแปลงค่าใน slice นั้นโดยตรง
	// 	func fill(nums *[]int) {
	//     *nums = append(*nums, 1, 2, 3)
	// }
	// configs.DB.Order("id DESC").Find(&users) //ไม่ใช้ users = configs.DB.Find(&users)  เพราะใช้ pointer

	configs.DB.Raw("select * from users order by id desc").Scan(&users)

	c.JSON(200, gin.H{
		"data": users,
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

	//สร้าง model
	user := models.User{
		Fullname: input.Fullname,
		Email:    input.Email,
		Password: input.Password,
	}
	//check email ซ้ำ
	userExist := configs.DB.Where("email =?", input.Email).First(&user)
	if userExist.RowsAffected == 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "มีผผู้ใช้งาน Email นี้ในระบบแล้ว"})
		return
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
	})
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "login",
	})
}

func GetById(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	result := configs.DB.First(&user, id) //โดน assign value

	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูลนี้"})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}

// ต้องคิดก่อนว่า ต้องการ ผลลัพธ์อันเดียว หรือ หลายอัน
func SearchByFullname(c *gin.Context) {
	fullname := c.Query("fullname") //?fullname=JohnWick

	var users []models.User
	result := configs.DB.Where("fullname LIKE ?", "%"+fullname+"%").Find(&users) //&users คือ ผลลัพท์ ที่โดย mutation
	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูลนี้"})
		return
	}
	c.JSON(200, gin.H{
		"data": users,
	})
}
