package service

import (
	"TinyURL/extend/utils"
	"TinyURL/models"
)

// UserService 用户服务层逻辑
type UserService struct {
	UserID   uint
	Phone    string
	Password string
}

func (us *UserService) QueryByPhone(phone string) (user *models.User, err error) {
	userModel := &models.User{}
	condition := map[string]interface{}{
		"phone": phone,
	}
	user, err = userModel.FindOne(condition)
	return
}

func (us *UserService) StoreUser(phone, password string) (userID uint, err error) {
	user := &models.User{
		Phone: phone,
		Username: phone,
		Password: password,
		IsActive: 1,
	}
	user.Password = utils.MakeSha1(user.Phone + user.Password)
	userID, err = user.Insert()
	return
}
