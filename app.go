package wechat

import (
	"fmt"
	"time"
	"encoding/xml"	
	"io/ioutil"
	"github.com/astaxie/beego/config"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"errors"
	"github.com/liujianping/wechat/entry"
)

type CallbackInterface interface {
	Initialize(app *WeChatApp, api	*ApiClient)
	MsgText(txt *entry.TextRequest, back chan interface{})
	MsgImage(img *entry.ImageRequest, back chan interface{})
	MsgVoice(voice *entry.VoiceRequest, back chan interface{})
	MsgVideo(video *entry.VideoRequest, back chan interface{})
	MsgLink(link *entry.LinkRequest, back chan interface{})
	Location(location *entry.LocationRequest, back chan interface{})
	EventSubscribe(appoid string, oid string, back chan interface{})
	EventUnsubscribe(appoid string, oid string, back chan interface{})
	EventMenu(appoid string, oid string, key string, back chan interface{})
}

type Handle func (app *WeChatApp, data []byte, back chan []byte)

type WeChatApp struct{
	AppHost 	string 
	AppPort 	int
	AppURI 		string 
	AppToken 	string 
	AppId		string 
	AppSecret 	string 
	AppPath		string
	Config 		config.ConfigContainer
	menu	  	*entry.Menu
	cb 			CallbackInterface
	handle 		Handle
	api 	    *ApiClient
	once		sync.Once
}

func NewWeChatApp() *WeChatApp {	
	app := new(WeChatApp)
	app.AppPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	os.Chdir(app.AppPath)

	app.AppHost = ""
	app.AppPort = 80
	app.AppURI = "/"

	app.AppToken = ""
	app.AppId = ""
	app.AppSecret = ""
	return app
}

func (app *WeChatApp) SetConfig(adapter, file string) error {
	var err error
	app.Config, err = config.NewConfig(adapter, file)
	if err != nil {
		return err
	} else {
		app.AppHost = app.Config.String("AppHost")

		if v, err := app.Config.Int("AppPort"); err == nil {
			app.AppPort = v
		}

		if v := app.Config.String("AppURI"); v != "" {
			app.AppURI = v
		}
		if v := app.Config.String("AppToken"); v != "" {
			app.AppToken = v
		}
		if v := app.Config.String("AppId"); v != "" {
			app.AppId = v
		}
		if v := app.Config.String("AppSecret"); v != "" {
			app.AppSecret = v
		}
	}
	return nil
}

func (app *WeChatApp) SetMenu(menu *entry.Menu){
	app.menu = menu
}

func (app *WeChatApp) SetCallback(callback CallbackInterface) {
	app.cb = callback	
}

func (app *WeChatApp) SetHandle(handle Handle) {
	app.handle = handle
}

func (app *WeChatApp) Run() {
	defer func(){
    	if x := recover(); x != nil {
			fmt.Println("wechat : app panic for <", x, ">")
		}
    }()

	if err := app.initialize(); err != nil {
		panic(err)
	} 

	http.HandleFunc(app.AppURI, app.uri)
	http.ListenAndServe(fmt.Sprintf("%s:%d", app.AppHost, app.AppPort), nil)
}


func (app *WeChatApp) initialize() error{
	if app.AppId == "" || app.AppToken == "" || app.AppSecret == "" {
		return errors.New("wechat: app id or token or secret not setting!")
	}

	if app.cb == nil && app.handle == nil {
		return errors.New("wechat: handle & callback both unset")
	}

	app.api = NewApiClient(app.AppToken, app.AppId, app.AppSecret)
	
	if app.cb != nil  {
		app.cb.Initialize(app, app.api)	
	}
	
	return nil
}

func (app *WeChatApp) buildMenu(){
	if app.menu != nil {
		if err := app.api.RemoveMenu(); err != nil {
			panic(err)
		}
		if err := app.api.CreateMenu(app.menu); err != nil {
			panic(err)
		}
	}
}

func (app *WeChatApp) uri(wr http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		signature := req.FormValue("signature")
		timestamp := req.FormValue("timestamp")
		nonce := req.FormValue("nonce")
		echostr := req.FormValue("echostr")

		if app.api.Signature(signature, timestamp, nonce) == true {
			wr.Write([]byte(echostr))
			app.once.Do(app.buildMenu)
		} else {
			wr.Write([]byte(""))
		}
	} else {
		if app.handle != nil {

		}
		if err := app.execute(wr, req); err != nil {
			Warn("wechat:", err)
		} 		
	}
}

func (app *WeChatApp) execute(wr http.ResponseWriter, req *http.Request) error {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	Debug("wechat: data \n", string(data))
	raw := make(chan []byte)
	defer close(raw)

	ch := make(chan interface{})
	defer close(ch)

	timeout := make(chan bool, 1)
	defer close(timeout)

	go func(c chan bool) {
		defer func(){
	    	if x := recover(); x != nil {
				Debug("wechat: ", x)
			}
	    }()

		time.Sleep(3e9) // 等待3秒钟
		c <- true
	}(timeout)

	if app.handle != nil {
		go app.handle(app, data, raw)
	}

	if app.cb != nil {
		request := &entry.Request{}
		err = xml.Unmarshal(data, request)
		if err != nil {
			return err
		}

		event := request.Event
		msgType := request.MsgType
		
		if "event" == msgType {
			//! event
			switch (event){
			case "subscribe":
				go app.cb.EventSubscribe(request.ToUserName, request.FromUserName, ch)
			case "unsubscribe":
				go app.cb.EventUnsubscribe(request.ToUserName, request.FromUserName, ch)
			case "CLICK":
				go app.cb.EventMenu(request.ToUserName, request.FromUserName, request.EventKey, ch)
			case "LOCATION":
				location := &entry.LocationRequest{}
				err = xml.Unmarshal(data, location)
				if err != nil {
					return err
				}
				go app.cb.Location(location, ch)
			default:
				return errors.New("unknown event ")
			}
		} else {
			//! other msg
			switch (msgType){
			case "text":
				text := &entry.TextRequest{}
				err = xml.Unmarshal(data, text)
				if err != nil {
					return err
				}
				go app.cb.MsgText(text, ch)
			case "image":
				image := &entry.ImageRequest{}
				err = xml.Unmarshal(data, image)
				if err != nil {
					return err
				}
				go app.cb.MsgImage(image, ch)			
			case "voice":
				voice := &entry.VoiceRequest{}
				err = xml.Unmarshal(data, voice)
				if err != nil {
					return err
				}
				go app.cb.MsgVoice(voice, ch)			
			case "video":
				video := &entry.VideoRequest{}
				err = xml.Unmarshal(data, video)
				if err != nil {
					return err
				}
				go app.cb.MsgVideo(video, ch)			
			case "location":
				location := &entry.LocationRequest{}
				err = xml.Unmarshal(data, location)
				if err != nil {
					return err
				}
				go app.cb.Location(location, ch)			
			case "link":
				link := &entry.LinkRequest{}
				err = xml.Unmarshal(data, link)
				if err != nil {
					return err
				}
				go app.cb.MsgLink(link, ch)
			}
		}
	}

	select{
	case r := <-raw:
		Debug("wechat: get response \n", string(r))
		wr.Write(r)	
	case b := <-ch:
		response,_ := xml.Marshal(b)
		Debug("wechat: get response \n", string(response))
		wr.Write(response)	
	case <-timeout:
		Warn("wechat: timeout for null response")
		wr.Write([]byte(""))
	}
	
	return nil
}
