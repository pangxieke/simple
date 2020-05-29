# Go实现web服务器
简单实现Go web框架功能。

实现如下功能
- 自定义路由
- 中间件
- 日志

## Controller

继承`BaseController`
```
type HelloController struct {
	BaseController
}
```

方法返回`error`
```
func (this *HelloController) Hello() (err error) {
	fmt.Println("hello controller, method Hello")
	if true {
		err = fmt.Errorf("this is error")
		return
	}
	return this.SuccessJson("success")
}

```
## Router
路由注入
```
r := New()

r.addPath("/hello", Action((*controllers.HelloController).Hello))
```

## Log
简单封装`log`日志

## Run
```
go run cmd/main/main.go
```

## Build
```$xslt
make build
```