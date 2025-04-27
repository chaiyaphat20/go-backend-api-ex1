package usercontroller

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"

	"example.com/gin-backend-api/configs"
	"example.com/gin-backend-api/models"
	"example.com/gin-backend-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/matthewhartstonge/argon2"
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
	var input InputLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//สร้าง model
	user := models.User{
		Email:    input.Email,
		Password: input.Password,
	}
	//get email
	userAccount := configs.DB.Where("email =?", input.Email).First(&user)
	if userAccount.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบผู้ใช้งานนี้ในระบบ"})
		return
	}

	//compare password
	ok, _ := argon2.VerifyEncoded([]byte(input.Password), []byte(user.Password))
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ีuser or password invalid"})
		return
	}

	//สร้าง token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 1).Unix(),
	})

	jwtSecret := os.Getenv("JWT_SECRET")
	fmt.Println(jwtSecret)
	accessToken, _ := claims.SignedString([]byte(jwtSecret))

	// ส่ง response กลับ
	c.JSON(http.StatusCreated, gin.H{
		"message":      "Login success",
		"access_token": accessToken,
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

func GetMe(c *gin.Context) {
	user := c.MustGet("user")
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"user":    user,
	})
}

// ต้องคิดก่อนว่า ต้องการ ผลลัพธ์อันเดียว หรือ หลายอัน
// &page=1&page_size=2
// {{url}}/users/search?fullname=chaiyaphat s&page=1&page_size=2
func SearchByFullname(c *gin.Context) {
	fullname := c.Query("fullname")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr) // แปลง string เป็น int

	var users []models.User
	var total int64

	// นับทั้งหมด ห้าม Scopes(utils.Paginate(c)) เด็ดขาด!!
	configs.DB.Model(&models.User{}).
		Where("fullname LIKE ?", "%"+fullname+"%").
		Count(&total)

	// ดึงข้อมูลจริง แบบแบ่งหน้า
	configs.DB.
		Where("fullname LIKE ?", "%"+fullname+"%").
		Scopes(utils.Paginate(c)).
		Find(&users)

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	c.JSON(200, gin.H{
		"data":         users,
		"total_data":   len(users),
		"current_page": page,
		"total_pages":  totalPages,
	})
}

//1:12:42 hr.
