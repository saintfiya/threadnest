package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	// 可根据需要自行添加字段
	UserID               int64  `json:"user_id"`
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

const TokenExpireDuration = time.Hour * 24 * 7

var customSecret []byte

// Init 注入 JWT 签名密钥。生产环境应通过环境变量提供随机密钥。
func Init(secret string) error {
	if len(secret) < 32 {
		return fmt.Errorf("jwt secret must contain at least 32 bytes")
	}
	customSecret = []byte(secret)
	return nil
}

func signingSecret() ([]byte, error) {
	if len(customSecret) == 0 {
		return nil, errors.New("jwt package is not initialized")
	}
	return customSecret, nil
}

// GenToken 生成JWT
func GenToken(userID int64, username string) (string, error) {
	secret, err := signingSecret()
	if err != nil {
		return "", err
	}
	// 创建一个我们自己的声明
	claims := CustomClaims{
		userID,
		username, // 自定义字段
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "my-project", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(secret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	secret, err := signingSecret()
	if err != nil {
		return nil, err
	}
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	var mc = new(CustomClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return secret, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, err
	}
	if token.Valid { //校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
