package service

import (
	"TinyURL/extend/conf"
	"TinyURL/extend/jwt"
	"TinyURL/extend/redis"
	"TinyURL/models"
	"fmt"
	goJWT "github.com/dgrijalva/jwt-go"
	"time"
)

type AuthService struct {

}


func (as *AuthService) GenerateToken(user *models.User) (string, error) {
	jwtInstance := jwt.NewJWT()
	nowTime := time.Now()
	expireTime := time.Duration(conf.ServerConf.JWTExpire) * time.Hour
	claims := jwt.CustomClaims{
		ID: user.ID,
		Phone: user.Phone,
		Level: user.Level,
		StandardClaims: goJWT.StandardClaims{
			ExpiresAt: nowTime.Add(expireTime).Unix(),
			Issuer: "TinyURL",
		},
	}
	token, err := jwtInstance.CreateToken(claims)
	if err != nil {
		return "", err
	}

	 ok, err := redis.Rdb.Set("TOKEN:" + user.Phone, token, expireTime).Result()
	 if err != nil {
	 	fmt.Println(err.Error())
	 	//utils.ResponseFormat(c, code.ServiceInsideError, nil)
	 	return "", err
	 }
	 fmt.Printf("%T", ok)
	 return token, nil
}