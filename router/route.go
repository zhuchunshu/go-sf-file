package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zhuchunshu/sf-file/modules/controllers"
	"github.com/zhuchunshu/sf-file/modules/middleware"
)

func InitRouter(app *fiber.App) {
	// 上传文件
	app.Post("/file", controllers.UploadFile).Use(middleware.HasFile)
	// 上传图片
	app.Post("/image", controllers.UploadImage).Use(middleware.HasFile)
}
