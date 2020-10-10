package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Link struct {
	gorm.Model
	LongUrl string `gorm:"type:varchar(255); not null" json:"long_url"`
	Keyword string `gorm:"type:varchar(50)" json:"keyword"` // 可用于客户自定义的Key存储
	TinyUrl string `gorm:"type:char(6)" json:"tiny_url"`
}

func (link *Link) Insert() (LinkID uint, err error) {
	result := DB.Create(&link)
	LinkID = link.ID
	if result.Error != nil {
		err = result.Error
	}
	return
}

func (link *Link) UpdateOne(linkID uint, data map[string]interface{}) bool {
	err := DB.Model(&Link{}).Where("id = ?", linkID).Update(data).Error
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func (link *Link) QueryOne(condition map[string]interface{}) (*Link, error) {
	var linkInfo Link
	result := DB.Where(condition).First(&linkInfo)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	if linkInfo.ID > 0 {
		return &linkInfo, nil
	}
	return nil, nil
}