package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
	"vue3-server/common/global"
	"vue3-server/model/system"
)

// CreateJwtToken 创建token
func CreateJwtToken(claims system.BaseClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, createClaims(claims))
	return token.SignedString([]byte(global.Config.Jwt.SigningKey))
}

// 根据传来的值生成Claims
func createClaims(claims system.BaseClaims) system.CustomClaims {
	newClaims := system.CustomClaims{
		BaseClaims: claims,
		BufferTime: global.Config.Jwt.BufferSecond,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    global.Config.Jwt.Issuer,                                                                         // 签名的发行者
			NotBefore: jwt.NewNumericDate(time.Now().Add(10 * -1 * time.Minute)),                                        // 签发时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.Config.Jwt.ExpiresSecond) * time.Second)), // 过期时间
		},
	}
	return newClaims
}
