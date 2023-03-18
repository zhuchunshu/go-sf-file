package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zhuchunshu/sf-file/modules/config"
)

func Index(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"code": 200,
		"msg":  "Hello " + config.Config.APPName,
		//"error": false,
	})
}
