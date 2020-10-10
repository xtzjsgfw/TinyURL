package code

import "net/http"

type Code struct {
	Status  int    `json:"status"`	// HTTP状态
	Code    int    `json:"code"`	// 业务码
	Message string `json:"message"`	// 业务响应信息
}

var (
	Success = &Code{http.StatusOK, 200000, "请求成功"}

	RequestParamError = &Code{http.StatusBadRequest, 400001, "请求参数有误"}
	ServiceInsideError = &Code{http.StatusInternalServerError, 500001, "服务器内部错误"}

	// 1000xxx 用户模块
	// 验证码接口相关
	PhoneFormatError = &Code{http.StatusBadRequest, 100002, "验证手机号无效"}
	SendVerificationCodeError = &Code{http.StatusBadRequest, 100003, "发送验证码失败"}
	// 注册接口相关
	PasswordFormatError = &Code{http.StatusBadRequest, 100003, "密码格式错误"}
	GetVerificationCodeError = &Code{http.StatusBadRequest, 100004, "获取验证码失败"}
	VerificationCodeUnmatchError = &Code{http.StatusBadRequest, 100005, "验证码错误"}
	UserIsExistError = &Code{http.StatusBadRequest, 100006, "注册用户已存在"}
	// 登陆接口相关
	UserIsNotExistError = &Code{http.StatusBadRequest, 100007, "该账号还未注册"}
	UserPasswordUnmatchError = &Code{http.StatusBadRequest, 100008, "登陆密码错误"}

	// Token相关
	TokenNotFound = &Code{http.StatusUnauthorized, 401001, "未携带Token，没有权限访问"}
	TokenInvalid = &Code{http.StatusUnauthorized, 401002, "Token无效"}


	// 600xxx Link模块
	// Link相关
	LinkNotExist = &Code{http.StatusBadRequest, 600000, "该短链接不存在"}
)
