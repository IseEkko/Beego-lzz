package MId

import (
	"fmt"
	"github.com/astaxie/beego/context"
)
//中间件使用
var Mid = func(ctx *context.Context){
	fmt.Println("Mid test success")
}