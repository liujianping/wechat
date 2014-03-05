package wechat

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/config"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"errors"
	"github.com/liujianping/wechat/entry"
)

type WeChatApp struct{
	AppHost 	string 
	AppPort 	int
	AppURI 		string 
	AppToken 	string 
	AppId		string 
	AppSecret 	string 
	AppPath		string
	menu	  	*entry.Menu
	config 		config.ConfigContainer
	callback 	CallbackInterface
	api 	    *ApiClient
	once		sync.Once
	*logs.BeeLogger
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
	app.config, err = config.NewConfig(adapter, file)
	if err != nil {
		return err
	} else {
		app.AppHost = app.config.String("AppHost")

		if v, err := app.config.Int("AppPort"); err == nil {
			app.AppPort = v
		}

		if v := app.config.String("AppURI"); v != "" {
			app.AppURI = v
		}
		if v := app.config.String("AppToken"); v != "" {
			app.AppToken = v
		}
		if v := app.config.String("AppId"); v != "" {
			app.AppId = v
		}
		if v := app.config.String("AppSecret"); v != "" {
			app.AppSecret = v
		}
	}
	return nil
}

func (app *WeChatApp) SetMenu(menu *entry.Menu){
	app.menu = menu
}

func (app *WeChatApp) SetCallback(callback CallbackInterface) {
	app.callback = callback	
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

	if app.callback == nil {
		return errors.New("wechat: callback interface is nil, please set callback first")
	}

	app.api = NewApiClient(app.AppToken, app.AppId, app.AppSecret)
	app.callback.Initialize(app, app.api)

	return nil
}

func (app *WeChatApp) buildMenu(){
	if app.menu != nil {
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

		if app.api.Check(signature, timestamp, nonce) == true {
			wr.Write([]byte(echostr))
			app.once.Do(app.buildMenu)
		} else {
			wr.Write([]byte(""))
		}
	} else {
		if err := app.execute(wr, req); err != nil {
			Warn("wechat:", err)
		} 		
	}
}

func (app *WeChatApp) execute(wr http.ResponseWriter, req *http.Request) error {
	return app.callback.Execute(wr,req)	
}
