package util

import (
	"SORS/pkg/constant"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaim struct {
	UserID uint
	jwt.RegisteredClaims
}

// SignToken 签发 Token
func SignToken(userID uint) (string, error) {
	// 声明签发Token的键值
	key := []byte(constant.Key)
	// 声明需要放入负载中的数据
	uc := &UserClaim{
		userID,
		jwt.RegisteredClaims{
			// 签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}
	// 获取 Token 对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	// 生成Token并返回
	return token.SignedString(key)
}

// ParseToken 实现一个解析 Token 的方法
func ParseToken(tokenString string) (*UserClaim, error) {
	// 使用 key 去解析 Token
	token, err := jwt.ParseWithClaims(tokenString, &UserClaim{}, func(t *jwt.Token) (interface{}, error) {
		// 将 Key 转化为 byte 切片
		return []byte(constant.Key), nil
	})
	if err != nil {
		return nil, err
	}
	// 从 Token 的负载里拿到 Claims，由于定制化了 Claims，故需要做类型转换
	if uc, ok := token.Claims.(*UserClaim); ok {
		return uc, nil
	}
	return nil, errors.New("ParseToken error")
}
