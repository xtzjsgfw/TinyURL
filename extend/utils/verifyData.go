package utils

import "regexp"

// @Summary 校验手机号
func VerifyPhoneFormat(phone string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}

// @Summary 校验密码格式
func VerifyPasswordFormat(password string) bool {
	length := len(password)
	if length < 6 || length > 20 {
		return false
	}
	return true
}

// @Summary 校验url格式
func VerifyUrlFormat(url string) bool {
	return true
}