package generate

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// 用于签名的字符串
var mySigningKey = []byte("myaof.cn")

// CustomClaims 自定义声明类型 并内嵌 jwt.RegisteredClaims
// jwt 包自带的 jwt.RegisteredClaims 只包含了官方字段
// 假设我们这里需要额外记录一个 username 字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

const TokenExpireDuration = time.Hour * 24

// GenRegisteredClaims 使用默认声明生成 jwt
func GenRegisteredClaims() (string, error) {
	// 1.创建 Claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		Issuer:    "may",
	}

	// 2.生成 token 对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 3.生成签名字符串
	return token.SignedString(mySigningKey)
}

func GenTokenByCustom(username string) (string, error) {
	// 1. 创建 Claims
	claims := CustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "may",
		},
	}
	// 2. 生成 token 对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 3. 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(mySigningKey)
}

// ParseRegisteredClaims 解析 jwt
func ParseRegisteredClaims(tokenString string) (*CustomClaims, error) {
	// 解析 token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")

	//parse, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
	//	return mySigningKey, nil
	//})
	//if err != nil {
	//	return false
	//}
	//return parse.Valid
}
