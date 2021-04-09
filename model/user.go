package model

type User struct {
	UserId	uint64  `gorm:"column:userId",json:"userId"` // 自增
	UserName string  `gorm:"column:username",json:"username"` //
	Password string `gorm:"column:password",json:"password"`
	//addTime string `gorm:"column:addTime",json:"addTime"`
}

func (User) TableName() string {
	return "user"
}
