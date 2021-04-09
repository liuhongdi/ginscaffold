package request

type LoginRequest struct {
	UserName string `form:"username" binding:"required,min=3"`
	PassWord string  `form:"password" binding:"required,min=3"`
}