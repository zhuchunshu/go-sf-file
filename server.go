package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/zhuchunshu/sf-file/modules/config"
	"os"
	"reflect"
	"strings"
)

var BasePath string

type Router struct {
	app         *fiber.App
	controllers []interface{}
}

func NewRouter(app *fiber.App) *Router {
	return &Router{
		app: app,
	}
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
		ServerHeader: "SForum File Server",
		AppName:      "SForum File v1.0.0",
		// 优化json
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		// 错误处理
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError
			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			// Send custom error page
			err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
			// Return from handler
			return nil
		},
	})

	// 监听端口
	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
