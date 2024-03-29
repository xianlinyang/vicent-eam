package main

import (
	_ "Eam_Server/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"token", "key", "Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	//2024-03-29 s tag=v1.0.1

	// 通过 /images/资源路径  可以访问static/images目录的内容
	// 例: /images/user/1.jpg 实际访问的是 static/images/user/1.jpg
	beego.SetStaticPath("/images", "static/img")

	// 通过 /css/资源路径  可以访问static/css目录的内容
	beego.SetStaticPath("/css", "static/css")

	// 通过 /js/资源路径  可以访问static/js目录的内容
	beego.SetStaticPath("/js", "static/js")

	beego.Run()
}
