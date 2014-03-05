# wechat Go 开发框架


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

## How to create a new custom wechat app use the package

- subclass a custom callback from wechat.callback 
- create a custome callback and set to the wechat app
- run the wechat app

具体用例参见 github.com/liujianping/wechatdemo
