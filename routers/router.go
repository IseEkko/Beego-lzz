package routers

import (
	"Beego-JWT-master/controllers"
	"Beego-JWT-master/controllers/MId"
	"github.com/astaxie/beego"
)

func init() {
//登录注册
	ns := beego.NewNamespace("/api",

		beego.NSNamespace("/user",
			beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
			beego.NSRouter("/register", &controllers.UserController{}, "post:CreateUser"),
		),
	)
	beego.Router("/ceshi",&controllers.CeshiController{},"post:Ceshi")
	//中间件的使用
	beego.InsertFilter("/ceshi",beego.BeforeRouter,MId.Mid)
	beego.AddNamespace(ns)
}
