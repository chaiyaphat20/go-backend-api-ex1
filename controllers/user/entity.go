package usercontroller

type InputRegister struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
