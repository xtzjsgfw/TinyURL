package service

import (
	"TinyURL/models"
	"fmt"
)

type LinkService struct {
}


func (ls *LinkService) StoreLink(originUrl string) (linkID uint, err error) {
	linkModel := &models.Link{
		LongUrl: originUrl,
	}
	linkID, err = linkModel.Insert()
	return
}

func (ls *LinkService) UpdateLink(linkID uint, data map[string]interface{}) bool {
	linkModel := &models.Link{}
	UpdateBool := linkModel.UpdateOne(linkID, data)
	return UpdateBool
}

func (ls *LinkService) QueryByKeyword(keyword string) (*models.Link, error) {
	fmt.Println("service中的keyword：", keyword)
	linkModel := &models.Link{}
	condition := map[string]interface{} {
		"id": keyword,
	}
	linkInfo, err := linkModel.QueryOne(condition)
	fmt.Println("service中的", linkInfo)
	fmt.Println("service中的", err)
	return linkInfo, err
}