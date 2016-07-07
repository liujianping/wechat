## wechat 开发框架

[![GoDoc](http://godoc.org/github.com/liujianping/wechat?status.png)](http://godoc.org/github.com/liujianping/wechat)

该框架仅为抛砖之作, 没有实现太多的微信公众号全部的api接口。仅实现接口包括:

-	[获取接口调用凭据](http://mp.weixin.qq.com/wiki/2/88b2bf1265a707c031e51f26ca5e6512.html)
-	[接收消息](http://mp.weixin.qq.com/wiki/17/fc9a27730e07b9126144d9c96eaf51f9.html)
-	[发送消息](http://mp.weixin.qq.com/wiki/18/c66a9f0b5aa952346e46dc39de20f672.html)
-	[自定义菜单](http://mp.weixin.qq.com/wiki/6/95cade7d98b6c1e1040cde5d9a2f9c26.html)
-	[用户管理](http://mp.weixin.qq.com/wiki/17/c807ee0f10ce36226637cebf428a0f6d.html)

###  快速开始

客户端快速开发指南:

````go
	
	import "github.com/liujianping/wechat"
	import "github.com/liujianping/wechat/entry"

	api := wechat.NewClient("appid", "appsecret")

	// 获取令牌
	var token entry.Token
	if err := api.Access(&token); err != nil {

	}

	// 获取用户信息
	var user_info subscriber.UserInfo
	if err := api.GetUserInfo("open_id", "zh_CN", &user_info); err != nil {

	}

	// 更多接口
	...

````

服务端(支持多应用)快速开发指南:

````go

	package main

	import (
		"log"

		"github.com/liujianping/wechat"
		"github.com/liujianping/wechat/entry"
	)

	func DemoHandle(app *wechat.Application, request *entry.Request) (interface{}, error) {
		log.Printf("demo app (%v)\n", app)
		log.Printf("demo msg (%v)\n", request)
		return nil, nil
	}

	func EchoHandle(app *wechat.Application, request *entry.Request) (interface{}, error) {
		log.Printf("echo app (%v)\n", app)
		log.Printf("echo msg (%v)\n", request)
		return nil, nil
	}

	func main() {
		demo := wechat.NewApplication("/demo", "demo_secret", "appid", "secret", false)

		btn11 := entry.NewButton("baidu").URL("http://baidu.com")
		btn12 := entry.NewButton("click").Event("event_click")
		demo.Menu(entry.NewMenu(btn11, btn12))


		echo := wechat.NewApplication("/echo", "echo_secret", "appid", "secret", false)

		btn21 := entry.NewButton("baidu").URL("http://baidu.com")
		btn22 := entry.NewButton("click").Event("event_click")
		echo.Menu(entry.NewMenu(btn21, btn22))

		serv := wechat.NewServer(":8080")
		
		serv.Application(demo, DemoHandle)
		serv.Application(echo, EchoHandle)

		serv.Start()
	}


````

### 例子

参考 demo 实现.


[老版本.README](https://github.com/liujianping/wechat/blob/v0.1/README.md)



