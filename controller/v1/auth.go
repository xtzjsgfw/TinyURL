/*
	登陆模块相关接口：验证码接口、注册、登陆、退出
*/
package v1

import (
	"TinyURL/extend/code"
	myRedis "TinyURL/extend/redis"
	"TinyURL/extend/utils"
	"TinyURL/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// AuthController 用户控制器
type AuthController struct {}

// VerificationCode 获取验证码请求参数
type VerificationCode struct {
	Phone string `json:"phone"`
}

// @Summary 发送验证码
func (ac *AuthController) SendCode(c *gin.Context) {
	var verificationCode = VerificationCode{}
	err := c.ShouldBindJSON(&verificationCode)
	if err != nil {
		fmt.Println("【验证码】接收手机号出错")
		utils.ResponseFormat(c, code.RequestParamError, nil)
		return
	}
	// 验证数据等
	verifyStatus := utils.VerifyPhoneFormat(verificationCode.Phone)
	if verifyStatus == false {
		fmt.Println("手机号无效")
		utils.ResponseFormat(c, code.PhoneFormatError, nil)
		return
	}
	// 生成验证码
	codeStr := utils.CreateVerificationCode()

	// 发送短信
	status := utils.YunPian(verificationCode.Phone, codeStr)
	if status != 0 {
		fmt.Println("发送频率太快")
		utils.ResponseFormat(c, code.SendVerificationCodeError, nil)
		return
	}

	// 存入redis，并设置失效时间
	expireTime := 60 * time.Second
	myRedis.Rdb.Set(verificationCode.Phone, codeStr, expireTime)

	// 响应用户
	utils.ResponseFormat(c, code.Success, gin.H{"phone": verificationCode.Phone})
}

type RegisterRequest struct {
	Phone string `json:"phone"`
	Code string `json:"code"`
	Password string `json:"password"`
}

// @Summary 注册功能
func (ac *AuthController) Register(c *gin.Context) {
	var registerRequest RegisterRequest
	err := c.ShouldBindJSON(&registerRequest)
	if err != nil {
		fmt.Println("【注册】接收参数出错")
		utils.ResponseFormat(c, code.RequestParamError, nil)
		return
	}

	// 验证该账户是否存在
	userService := service.UserService{
		Phone: registerRequest.Phone,
		Password: registerRequest.Password,
	}
	userInfo, err := userService.QueryByPhone(userService.Phone)
	if err != nil {
		fmt.Println("出错")
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}
	if userInfo != nil {
		fmt.Println("注册用户已经存在")
		utils.ResponseFormat(c, code.UserIsExistError, nil)
		return
	}

	// 验证验证码
	result, _ := myRedis.Rdb.Get(registerRequest.Phone).Result()
	if result == "" {
		fmt.Println("验证码过期/没有获取验证码")
		utils.ResponseFormat(c, code.GetVerificationCodeError, nil)
		return
	}
	if result != registerRequest.Code {
		utils.ResponseFormat(c, code.VerificationCodeUnmatchError, nil)
		return
	}

	// 验证密码格式是否正确
	passwordStatus := utils.VerifyPasswordFormat(registerRequest.Password)
	if passwordStatus == false {
		utils.ResponseFormat(c, code.PasswordFormatError, nil)
		return
	}

	// 存入数据
	userID, err := userService.StoreUser(userService.Phone, userService.Password)
	if err != nil {
		fmt.Println(err.Error())
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}

	// 返回响应
	utils.ResponseFormat(c, code.Success, gin.H{"userId": userID})
}

type LoginRequest struct {
	Phone string `json:"phone"`
	Password string `json:"password"`
}

// @Summary 登陆功能
func (ac *AuthController) Login(c *gin.Context) {
	var loginRequest LoginRequest
	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		fmt.Println("【登陆】接收参数错误")
		utils.ResponseFormat(c, code.RequestParamError, nil)
		return
	}

	// 验证手机号格式
	verifyPhone := utils.VerifyPhoneFormat(loginRequest.Phone)
	if verifyPhone == false {
		fmt.Println("【登陆】手机号格式错误")
		utils.ResponseFormat(c, code.PhoneFormatError, nil)
		return
	}

	// 验证密码格式
	verifyPassword := utils.VerifyPasswordFormat(loginRequest.Password)
	if verifyPassword == false {
		fmt.Println("【登陆】密码格式错误")
		utils.ResponseFormat(c, code.PasswordFormatError, nil)
		return
	}

	// 手机号是否存在
	userService := service.UserService{
		Phone: loginRequest.Phone,
		Password: loginRequest.Password,
	}
	userInfo, err := userService.QueryByPhone(userService.Phone)
	if err != nil {
		fmt.Println("系统内部错误")
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}
	if userInfo == nil {
		fmt.Println("【登陆】该账号还未注册")
		utils.ResponseFormat(c, code.UserIsNotExistError, nil)
		return
	}
	if userInfo.Password != utils.MakeSha1(loginRequest.Phone + loginRequest.Password)[:30] {
		fmt.Println("【登陆】登陆密码错误")
		utils.ResponseFormat(c, code.UserPasswordUnmatchError, nil)
		return
	}

	authService := service.AuthService{}
	token, err := authService.GenerateToken(userInfo)
	if err != nil {
		utils.ResponseFormat(c, code.ServiceInsideError, nil)
		return
	}

	utils.ResponseFormat(c, code.Success, gin.H{"token": token, "userId": userInfo.ID})
}

func (ac *AuthController) Test(c *gin.Context) {
	//jwtInstance := jwt.NewJWT()
	//claims, err := jwtInstance.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NCwicGhvbmUiOiIxMzE4MjAwOTUwMyIsImxldmVsIjowLCJleHAiOjE1OTc3Njc4OTAsImlzcyI6IlRpbnlVUkwifQ.uTxzNHgeuh-3Kd-I8f3Z28SbndFp30nEJgA04kdil4Q")
	//if err != nil {
	//	fmt.Printf("我错误了")
	//	return
	//}
	//fmt.Println(claims)
	c.JSON(200, gin.H{
		"msg": "请求成功",
	})
}