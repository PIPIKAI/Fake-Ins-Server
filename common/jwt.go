package common

import (
	"time"

	"github.com/PIPIKAI/Ins-gin-vue/server/model"

	"github.com/dgrijalva/jwt-go"
)

// Jwt 加密密钥
var jwtKey = []byte("theJwtKeyyy")

type Claims struct {
	UserID uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User, d time.Duration) (string, error) {
	//   设置过期时间
	expireTime := time.Now().Add(d)
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			// 签发人
			Issuer:  "ppk",
			Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
