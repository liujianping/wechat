# wechat Go 开发框架

[![GoDoc](http://godoc.org/github.com/liujianping/wechat?status.png)](http://godoc.org/github.com/liujianping/wechat)

##  安装
	
框架主要依赖包：

	go get github.com/astaxie/beego/logs
	go get github.com/astaxie/beego/config
	go get github.com/liujianping/wechat

##  使用

### 1) you need import it

	import (
		"github.com/liujianping/wechat"
	)

### 2) create a new wechat app

	app := wechat.NewWeChatApp()

### 3) configure the wechat app

	app.SetConfig("ini", "demo.ini")

根据具体微信服务号进行配置, demo.in

	AppHost=""
	AppPort=8081
	AppURI="/weixin"
	AppToken="wx_gh_385c8159a341"
	AppId="wx1ac20fc50222a85e"
	AppSecret="9690c087a0d508bd8fe3faab63fb39a8"


### 3) execute the wechat app

在
	app.Run()

之前, 设置 CallabckInterface 对象

	app.SetCallback()
	
### 4) logging for the wechat app

代码中，可以使用

	wechat.Debug(...)
	wechat.Trace(...)
	wechat.Info(...)

等等接口打印日志，不过打印日志前必须先设置好日志类型

	wechat.SetLogger("console", "")

## about inner packages

### entry package

	import (
		"github.com/liujianping/wechat/entry"
	)

## create menu for the app

菜单的创建

	import "github.com/liujianping/wechat/entry"

	menu := entry.NewMenu()

	btn1 := entry.NewViewButton("新浪","http://sina.com")
	btn2 := entry.NewClickButton("点击","EVENT_MENU_CLICK")
	btn3 := entry.NewButton("更多")
	btn3.Append(entry.NewViewButton("腾讯","http://qq.com"))
	btn3.Append(entry.NewViewButton("百度","http://baidu.com"))
	btn3.Append(entry.NewViewButton("点评","http://dianping.com"))
	menu.Add(btn1)
	menu.Add(btn2)
	menu.Add(btn3)

	client := api.NewApiClient(cs_token, cs_appid, cs_appsecret)
	client.SetCache("redisx",`{"conn":":6379"}`)
	client.CreateMenu(menu)

在WeChatAPP中添加菜单

	menu := entry.NewMenu()

	btn1 := entry.NewViewButton("新浪","http://sina.com")
	btn2 := entry.NewClickButton("点击","EVENT_MENU_CLICK")
	btn3 := entry.NewButton("更多")
	btn3.Append(entry.NewViewButton("腾讯","http://qq.com"))
	btn3.Append(entry.NewViewButton("百度","http://baidu.com"))
	btn3.Append(entry.NewViewButton("点评","http://dianping.com"))
	menu.Add(btn1)
	menu.Add(btn2)
	menu.Add(btn3)

	app.SetMenu(menu)

## How to create a new custom wechat app use the package

- subclass a custom callback from wechat.callback 
- create a custome callback and set to the wechat app
- run the wechat app

具体用例参见 [wechatdemo](https://github.com/liujianping/wechatdemo)
