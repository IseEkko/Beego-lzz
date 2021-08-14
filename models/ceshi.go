package models

import (
	"Beego-JWT-master/function"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Ceshi struct {
	Id int `orm:"column(id);auto"`
	Username string
	//添加时间戳
    function.Datatime
}

func (this *Ceshi)Test(c Ceshi)interface{}{
	o := orm.NewOrm()
	data ,err := o.Insert(&c)
	if err != nil{
	fmt.Println(err)
	}
   return data
}