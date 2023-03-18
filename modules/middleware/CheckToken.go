package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zhuchunshu/sf-file/modules/config"
)

func CheckToken(c *fiber.Ctx) error {
	return authenticate()(c)
}

// 认证中间件
func authenticate() func(*fiber.Ctx) error {
	// 返回一个新的处理函数
	return func(c *fiber.Ctx) error {
		// 检查 Authorization 头部是否存在
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "缺少 Authorization 头部")
		}

		// 检查 Authorization 头部是否以 Bearer 开头
		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			return fiber.NewError(fiber.StatusUnauthorized, "无效的 Authorization 头部格式")
		}

		// 获取访问令牌
		accessToken := authHeader[7:]

		// 检查访问令牌是否有效
		if !isValidToken(accessToken) {
			return fiber.NewError(fiber.StatusUnauthorized, "无效的访问令牌")
		}

		// 调用下一个中间件或处理程序
		return c.Next()
	}
}

// 检查访问令牌是否有效
func isValidToken(accessToken string) bool {
	// 在这里进行访问令牌的验证
	return accessToken == config.Config.Token
}
