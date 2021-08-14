package function

import "time"

//关于时间的公共函数

/***
时间戳的定义，主要用于结构体的创建时间戳
 */
type Datatime struct {
	Created_at 			time.Time		`orm:"auto_now_add;type(datetime)"`
	Updated_at			time.Time		`orm:"auto_now;type(datetime)"`
}