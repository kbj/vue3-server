package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
	"vue3-server/common/global"
	"vue3-server/model/system"
)

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

// CreateJwtToken 创建token
func CreateJwtToken(claims system.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(global.Config.Jwt.SigningKey))
}

// CreateClaims 根据传来的值生成Claims
func CreateClaims(claims system.BaseClaims) system.CustomClaims {
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

// ParseJwtToken 解析Token
func ParseJwtToken(tokenString string) (*system.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &system.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return global.Config.Jwt.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*system.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

// RefreshJwtToken 刷新新的token
func RefreshJwtToken(oldToken string, claims system.CustomClaims) (string, error) {
	// 使用并发控制
	v, err, _ := global.ConcurrencyControl.Do("JWT:"+oldToken, func() (interface{}, error) {
		return CreateJwtToken(claims)
	})
	return v.(string), err
}
