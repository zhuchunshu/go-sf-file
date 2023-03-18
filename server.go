package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/zhuchunshu/sf-file/modules/config"
	"github.com/zhuchunshu/sf-file/modules/controllers"
	"github.com/zhuchunshu/sf-file/modules/middleware"
	"github.com/zhuchunshu/sf-file/router"
	"os"
	"reflect"
	"strings"
)

var BasePath string

type Router struct {
	app         *fiber.App
	controllers []interface{}
}

func (r *Router) RegisterController(controller interface{}) {
	r.controllers = append(r.controllers, controller)
}

func (r *Router) AutoRegisterRoutes() {
	for _, controller := range r.controllers {
		v := reflect.ValueOf(controller)
		t := v.Type()

		for i := 0; i < t.NumMethod(); i++ {
			method := t.Method(i)
			if strings.HasPrefix(method.Name, "Handle") {
				path := strings.TrimPrefix(method.Name, "Handle")
				path = "/" + strings.ToLower(path)

				handler := v.MethodByName(method.Name).Interface().(func(*fiber.Ctx) error)

				r.app.Get(path, handler)
			}
		}
	}
}

func init() {

	// 设置全局BasePath
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	BasePath = dir

	// 初始化config
	config.Get(BasePath)
}

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader: config.Config.APPName,
		AppName:      "SForum File v1.0.0",
		// 优化json
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		// 错误处理
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Retrieve the custom status code if it's a *fiber.Error
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"code":  fiber.StatusInternalServerError,
					"error": true,
					"msg":   err.Error(),
				})
			}
			// Return from handler
			return nil
		},
	})
	// static
	app.Static("/uploads", config.Config.UploadPath)
	// home
	app.Get("/", controllers.Index)
	// 初始化中间件
	middleware.Init(app)
	// 初始化路由
	router.InitRouter(app)

	// 监听端口
	err := app.Listen(":" + fmt.Sprint(config.Config.Port))
	if err != nil {
		return
	}
}
