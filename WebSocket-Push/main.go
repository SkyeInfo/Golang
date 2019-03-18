package main

import (
	"github.com/astaxie/beego"
	"git.culiu.org/pintuan-go/controllers"
	"golang.org/x/net/websocket"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("ok")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//列表页展示
	beego.Handler("/ws", websocket.Handler(controllers.SockServer))
	//详情信息展示
	beego.Handler("/ds", websocket.Handler(controllers.SockServerDeteail))
	beego.Router("/", &MainController{})
	beego.Run()
}
