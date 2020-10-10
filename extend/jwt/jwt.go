package jwt

import (
	"TinyURL/extend/conf"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	JWTSecret []byte
}

// CustomClaims jwt信息
type CustomClaims struct {
	ID    uint   `json:"id"`
	Phone string `json:"phone"`
	Level int    `json:"level"`
	jwt.StandardClaims
}

var (
	// ErrTokenExpired 验证令牌失效
	ErrTokenExpired = errors.New("Token is expired")
	// ErrTokenNotValidYet 验证令牌未激活
	ErrTokenNotValidYet = errors.New("Token not active yet")
	// ErrTokenMalformed 验证并非属于令牌
	ErrTokenMalformed = errors.New("That's not even a token")
	// ErrTokenInvalid 验证为无效的令牌
	ErrTokenInvalid = errors.New("Couldn't handle this token")
)

func NewJWT() *JWT {
	return &JWT{[]byte(conf.ServerConf.JWTSecret)}
}


// CreateToken 生成 Token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(j.JWTSecret)
	return token, err
}

// ParseToken 解析 Token
func (j *JWT) ParseToken(token string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JWTSecret, nil
	})
	if err != nil {
		if ve,ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			}else if ve.Errors&jwt.ValidationErrorExpired != 0{
				return nil, ErrTokenExpired
			}else if ve.Errors&jwt.ValidationErrorNotValidYet != 0{
				return nil, ErrTokenNotValidYet
			}else {
				return nil, ErrTokenInvalid
			}
		}
	}

	if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
		return claims, nil
	} else {
		return nil, ErrTokenInvalid
	}
}
