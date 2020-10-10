package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type: varchar(20)"`
	Phone    string `gorm:"type: char(11); unique;not null" json:"phone"`
	Password string `gorm:"type: varchar(100)" json:"password"`
	IsActive int    `gorm:"type:int(1)" json:"is_active"` // 为0表示不可用，为1表示可用
	Level    int    `gorm:"type: int; default: 0" json:"level"`
}

func (user *User) FindOne(condition map[string]interface{}) (*User, error) {
	var userInfo User
	result := DB.Where(condition).First(&userInfo)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	if userInfo.ID > 0 {
		return &userInfo, nil
	}
	return nil, nil
}

func (user *User) Insert() (UserID uint, err error) {
	result := DB.Create(&user)
	UserID = user.ID
	if result.Error != nil {
		err = result.Error
	}
	return
}
