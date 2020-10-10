package models

import "github.com/jinzhu/gorm"

type linkInfo struct {
	gorm.Model
	IP string `gorm:"type: char(15)"`
	Equipment string `gorm:""`
	Browser string
	IPAddress string
	RedirectStatus string
}
