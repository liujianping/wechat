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
func (e *Echo) MsgText(txt *entry.TextRequest, back chan interface{}){
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

