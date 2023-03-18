package middleware

import "github.com/gofiber/fiber/v2"

func HasFile(c *fiber.Ctx) error {
	// 判断是否上传了文件

	file, err := c.FormFile("file")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if file == nil {
		return fiber.NewError(fiber.StatusBadRequest, "缺少文件")
	}

	return c.Next()
}
