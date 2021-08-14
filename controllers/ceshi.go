package controllers

import (
	"Beego-JWT-master/function"
	"Beego-JWT-master/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type CeshiController struct {
	beego.Controller
}
type User struct {
	Id       int
	Username string `valid:"Required"`
	Age      int    `valid:"Range(1, 140)"`
}
type CeshiEorr struct{
	Code int
	Eorr string
	Data interface{}
}

func (u *User) Valid(v *validation.Validation) {
fmt.Println("金图")
}


func (u *CeshiController)Ceshi(){
	//var user  User
	//valid := validation.Validation{}
	////ParseForm返回的是错误的信息，获取的信息是在原来的变量里面，这里是user中
	//username := u.ParseForm(&user)
	//if username !=nil{
	//	fmt.Println("错误：",username)
	//}
	//b,err :=valid.Valid(&user)
	//if err != nil {
	//	// handle error
	//}
	//if !b {
	//	// validation does not pass
	//	// blabla...gg
	//	for _, err := range valid.Errors {
   //        ceshi := CeshiEorr{
   //        	422,
   //        	err.Message,
	//		}
	//		u.Data["json"] = &ceshi
	//		u.ServeJSON()
	//	}
	//	return
	//}
	//
	//u.Data["json"] = &user
	//u.ServeJSON()

	user := models.Ceshi{
		Username: "laa",
	}
   data := user.Test(user)
   if data != nil {
	   u.Data["json"] = function.Json_success(200,"测试成功",data)
	   u.ServeJSON()
	}else{
	   u.Data["json"] = function.Json_fail(100,"测试成功",data)
	   u.ServeJSON()
	}

}