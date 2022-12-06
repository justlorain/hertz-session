package mysql

import (
	"hertz-session/pkg/consts"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (u *User) TableName() string {
	return consts.UserTableName
}

func CreateUsers(users []*User) error {
	return DB.Create(users).Error
}

func FindUserByNameOrEmail(username, email string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.Where("username = ?", username).Or("email = ?", email).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func CheckUser(username, password string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.Where("username = ? AND password = ?", username, password).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
