package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/zhuchunshu/sf-file/modules/config"
	"time"
)

func UploadFile(c *fiber.Ctx) error {
	// 从表单中获取文件
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	// 确定上传的文件名和文件路径
	filename := file.Filename
	filename = fmt.Sprintf("%d-%s", time.Now().Unix(), filename)

	// 上传路径
	uploadPath := config.Config.UploadPath
	// 判断上传路径结尾是否有 / 如果有，则去除
	if uploadPath[len(uploadPath)-1:] == "/" {
		uploadPath = uploadPath[:len(uploadPath)-1]
	}
	// 获取当前时间作为文件夹名
	now := time.Now()
	folderName := now.Format("2006/01/02")

	// 文件路径
	filePath := fmt.Sprintf("%s/%s", folderName, filename)

	// 完整路径
	path := fmt.Sprintf("%s/%s", uploadPath, filePath)

	// 将文件保存到服务器上的 "uploads" 文件夹中
	if err := c.SaveFile(file, path); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"code":  200,
		"msg":   "上传成功",
		"error": false,
		"data": fiber.Map{
			"path":     path,
			"filePath": filePath,
			"url":      fmt.Sprintf("%s/%s", config.Config.Url, filePath),
		},
	})
}
