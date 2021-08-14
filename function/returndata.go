package function

//返回相关的函数

type Dataerr struct {
	Code  int
	Msg string
	Data interface{}
}
//成功返回的函数
func Json_success(code int,msg string,data interface{})*Dataerr{
	return &Dataerr{
		code,
		msg,
		data,
	}
}
//失败返回的函数
func Json_fail(code int,msg string,data interface{})*Dataerr{
	return &Dataerr{
		code,
		msg,
		data,
	}
}