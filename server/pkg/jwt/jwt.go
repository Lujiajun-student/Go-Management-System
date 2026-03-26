// Package jwt JWT 工具类，生成token和解析token，以及获取当前登录用户的id及用户信息
package jwt

import (
	"Go-Management-System/api/entity"
	"Go-Management-System/common/constant"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// userStdClaims 自定义Claims结构体，包含用户数据以及标准字段
type userStdClaims struct {
	entity.JwtAdmin
	jwt.StandardClaims
}

// TokenExpireDuration 过期时间
const TokenExpireDuration = time.Hour * 24

// Secret token密钥，对称加密，服务端用同一个Secret来签名和验证
var Secret = []byte("admin-go-api")

// 定义错误
var (
	ErrAbsent  = "jwt token absent"  // 令牌不存在
	ErrInvalid = "jwt token invalid" // 令牌无效
)

// GenerateTokenByAdmin 根据用户信息生成token
func GenerateTokenByAdmin(admin entity.SysAdmin) (string, error) {
	// 获取jwt专用的实体，专门构建这个实体的目的是避免敏感信息存到token中
	var jwtAdmin = entity.JwtAdmin{
		Id:       admin.ID,
		Username: admin.Username,
		Nickname: admin.Nickname,
		Icon:     admin.Icon,
		Email:    admin.Email,
		Phone:    admin.Phone,
		Note:     admin.Note,
	}
	// 构建完整的Claims
	c := userStdClaims{
		// 放入业务数据
		jwtAdmin,
		// 放入标准字段
		jwt.StandardClaims{
			// 设置过期时间
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			// 签发人
			Issuer: "admin",
		},
	}
	// 签名并传入Claims来生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用密钥进行哈希运算
	return token.SignedString(Secret)
}

// ValidateToken 解析JWT
func ValidateToken(tokenString string) (*entity.JwtAdmin, error) {
	// 没有token直接返回
	if tokenString == "" {
		return nil, errors.New(ErrAbsent)
	}

	// 通过Secret解析token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return Secret, nil
	})
	if token == nil {
		return nil, errors.New(ErrInvalid)
	}
	// 准备Claims容器
	claims := userStdClaims{}
	// 传入Claims容器来解析token
	_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}

	return &claims.JwtAdmin, nil
}

// GetAdminId 返回用户Id
func GetAdminId(c *gin.Context) (uint, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return 0, errors.New("can't get user id")
	}
	admin, ok := u.(*entity.JwtAdmin)
	if ok {
		return admin.Id, nil
	}
	return 0, errors.New("can't convert to id struct")
}

// GetAdminName 返回用户名
func GetAdminName(c *gin.Context) (string, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return "", errors.New("can't get user name")
	}
	admin, ok := u.(*entity.JwtAdmin)
	if ok {
		return admin.Username, nil
	}
	return "", errors.New("can't convert to api name")
}

// GetAdmin 返回admin信息
func GetAdmin(c *gin.Context) (*entity.JwtAdmin, error) {
	u, exist := c.Get(constant.ContextKeyUserObj)
	if !exist {
		return nil, errors.New("can't get user id")
	}
	admin, ok := u.(*entity.JwtAdmin)
	if ok {
		return admin, nil
	}
	return nil, errors.New("can't convert to api admin struct")
}
