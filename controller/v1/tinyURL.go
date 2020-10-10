package v1

import (
	"TinyURL/extend/code"
	"TinyURL/extend/conf"
	"TinyURL/extend/utils"
	"TinyURL/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type URLController struct{}

type IncNumTinyUrlRequest struct {
	OriginUrl string `json:"origin_url"`
}

type CustomTinyUrlRequest struct {
	OriginUrl string `json:"origin_url"`
	KeyWord   string `json:"key_word"`
}

func (uc *URLController) GenerateTinyUrlByIncreaseNum(c *gin.Context) {
	var incNumTinyUrlRequest IncNumTinyUrlRequest
	err := c.ShouldBindJSON(&incNumTinyUrlRequest)
	if err != nil {
		fmt.Println("解析参数出从")
		utils.ResponseFormat(c, code.RequestParamError, nil)
		return
	}
	fmt.Println(incNumTinyUrlRequest)

	// 将长URL存入数据库，获得自增ID
	linkService := service.LinkService{}
	linkID, err := linkService.StoreLink(incNumTinyUrlRequest.OriginUrl)
	if err != nil {
		fmt.Println("存储连接出错")
		fmt.Println(err.Error())
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}
	fmt.Println(linkID)

	// 利用自增ID拼接生成短链接
	tinyCode := utils.Base62Encode(int(linkID))
	tinyUrl := conf.ServerConf.DomainName + tinyCode
	fmt.Println(tinyUrl)

	// 将短链接和key更新到数据库
	linkIDStr := strconv.Itoa(int(linkID))
	updateDate := map[string]interface{} {
		"tiny_url": tinyCode,
		"keyword": linkIDStr,
	}
	if ok := linkService.UpdateLink(linkID, updateDate); !ok {
		fmt.Println("更新失败")
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}

	utils.ResponseFormat(c, code.Success, gin.H{"tiny_url": tinyUrl})
}

func (uc *URLController) GenerateTinyUrlByCustomKey(c *gin.Context) {
	var result interface{}
	_ = c.ShouldBindHeader(result)
	fmt.Println(result)
	a := c.GetHeader("url")
	fmt.Println(a)
}


// 重定向短链接
func (uc *URLController) RequestByTinyUrl(c *gin.Context) {
	// 解析短链接
	var tinyUrl string
	tinyUrl = c.Param("tinyUrl")
	tinyUrl = tinyUrl[1:]

	// 对短链接解码
	urlKeyword := utils.Base62Decode(tinyUrl)

	// 查询短链接对应的长链接
	linkService := &service.LinkService{}
	urlKeywordStr := strconv.Itoa(urlKeyword)
	linkInfo, err := linkService.QueryByKeyword(urlKeywordStr)
	fmt.Println(linkInfo)
	fmt.Println(err)
	if err != nil {
		fmt.Println("查询失败:", err.Error())
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}
	if linkInfo == nil {
		fmt.Println("没有查询到记录")
		utils.ResponseFormat(c, code.LinkNotExist, nil)
		return
	}
	longUrl := linkInfo.LongUrl
	fmt.Println("我是longurl", longUrl)
	c.Redirect(http.StatusFound, longUrl)
	return
}
