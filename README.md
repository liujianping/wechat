# wechat Go 开发框架

[![GoDoc](http://godoc.org/github.com/liujianping/wechat?status.png)](http://godoc.org/github.com/liujianping/wechat)

##  [老版本](https://github.com/liujianping/wechat/tree/v0.1)

	[README](https://github.com/liujianping/wechat/blob/v0.1/README.md)

##  快速开始

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

	import "github.com/liujianping/wechat"
	import "github.com/liujianping/wechat/entry"	

	//! 应用1
	app1 := wechat.NewApplication("uri1", "token", "appid", "appsecret", handle1)
	
	btn11 := entry.NewButton("b1").URL("http://"))
	btn12 := entry.NewButton("b2").Event("event_unique_id")
	btn13 := entry.NewButton("b3").Append(entry.NewButton("c1").URL("http://")).Append(entry.NewButton("c2").Event(""))
	app1.SetMenu(entry.NewMenu(btn11, btn12, btn13))

	//! 应用2
	app2 := wechat.NewApplication("uri2", "token", "appid", "appsecret", handle2)

	btn21 := entry.NewButton("b1").URL("http://"))
	btn22 := entry.NewButton("b2").Event("event_unique_id")
	btn23 := entry.NewButton("b3").Append(entry.NewButton("c1").URL("http://")).Append(entry.NewButton("c2").Event(""))
	app2.SetMenu(entry.NewMenu(btn21, btn22, btn23))

	//! 服务
	serv := wechat.NewServer("host:port")

	serv.URIApplication(app1).URIApplication(app2)

	serv.Start()

	serv.Stop()

````

开发自己的微信公众号的服务功能, 只需要实现自己的 IReqeustHandle 接口即可。

````go

	import "github.com/liujianping/wechat"
	import "github.com/liujianping/wechat/entry"

	type Echo struct{
	}

	func (echo Echo) Text(text *entry.Text) {

		wechat.RenderXML(wr, text)
	}

	func (echo Echo) Image(image *entry.Image) {
		
	}

	func (echo Echo) Voice() {
		
	}

	func (echo Echo) Video() {
		
	}

	func (echo Echo) Link() {
		
	}

	func (echo Echo) Location() {
		
	}

	func (echo Echo) Subscribe() {
		
	}

	func (echo Echo) UnSubscribe() {
		
	}

	func (echo Echo) () {
		
	}


````



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

### cache package
	
	该cache主要是修改了 github.com/astaxie/beego/cache 中，redis类型缓存无超时功能的缺陷。
	增加了 redisx 类型, 配置与redis类型相同，但是支持超时功能。

### entry package

该entry包主要提供微信公众平台的的

- 请求消息
- 响应消息
- 客服消息
- 菜单
- 订阅者

几种类型数据的定义

	import (
		"github.com/liujianping/wechat/entry"
	)

## create menu for the app

菜单的创建

	import "github.com/liujianping/wechat"
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

	client := wechat.NewApiClient(cs_token, cs_appid, cs_appsecret)
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

创建一个新的CallbackInteface类型并实现相关接口函数。

````
import "github.com/liujianping/wechat"
import "github.com/liujianping/wechat/entry"

type Echo struct{
	Name string
	wechat.Callback
}
func NewEcho(name string) *Echo{
	return &Echo{name}
}
func (e *Echo) MsgText(txt *entry.Textentry, back chan interface{}){
	wechat.Info("Echo: MsgText ", txt)
}
func (e *Echo) MsgImage(img *entry.ImageRequest, back chan interface{}){
	wechat.Info("Echo: MsgImage ", img)
}
func (e *Echo) MsgVoice(voice *entry.VoiceRequest, back chan interface{}){
	wechat.Info("Echo: MsgVoice ", voice)
}
func (e *Echo) MsgVideo(video *entry.VideoRequest, back chan interface{}){
	wechat.Info("Echo: MsgVideo ", video)
}
func (e *Echo) MsgLink(link *entry.LinkRequest, back chan interface{}){
	wechat.Info("Echo: MsgLink ", link)
}
func (e *Echo) Location(location *entry.LocationRequest, back chan interface{}){
	wechat.Info("Echo: Location ", location)
}

func (e *Echo) EventSubscribe(appoid string, oid string, back chan interface{}){
	wechat.Info("Echo: EventSubscribe ", oid)
	var subscriber entry.Subscriber
	if err := e.Api.GetSubscriber(oid, &subscriber); err != nil {
		wechat.Error("Echo: get subscriber failed ", err)
	}

	response := entry.NewTextResponse(appoid, oid, fmt.Sprintf("%s 欢迎您的关注!", subscriber.Nickname))

	back <- response
}
func (e *Echo) EventUnsubscribe(appoid string, oid string, back chan interface{}){
	wechat.Info("Echo: EventUnsubscribe ", oid)	

}
func (e *Echo) EventMenu(appoid string, oid string, key string, back chan interface{}){
	wechat.Info("Echo: EventMenu ", oid, key)	
}

//! 设置
app.SetCallback(NewEcho("demo"))

````
- create a handle and set to the wechat app

创建一个Handle方法实现

````
func DemoHandle(data []byte, back chan []byte){
	wechat.Info("recieve raw msg:\n", string(data))
	//! TODO: anything you want to do
	//! send empty string back
	back <- []byte("")
}

//! 设置
app.SetHandle(DemoHandle)


````
- run the wechat app

callback 和 handle 二者取一个即可。如果两者均设置了, 两个都会被调用。

具体用例参见:
- [wechatdemo](https://github.com/liujianping/wechatdemo)
- [wechatdemo2](https://github.com/liujianping/wechatdemo2)

