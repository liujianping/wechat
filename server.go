package wechat

import (
	"crypto/sha1"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/kataras/iris"
	"github.com/liujianping/wechat/entry"
)

type Application struct {
	uri      string
	token    string
	appid    string
	secret   string
	customer bool
	menu     *entry.Menu
	api      *Client
}

type Handle func(app *Application, request *entry.Request) (response interface{}, err error)

func NewApplication(uri, token, appid, secret string, customer bool) *Application {
	return &Application{
		uri:      uri,
		token:    token,
		appid:    appid,
		secret:   secret,
		customer: customer,
	}
}

func (app *Application) Menu(menu *entry.Menu) *Application {
	app.menu = menu
	return app
}

func (app *Application) Api() *Client {
	if app.api == nil {
		app.api = NewClient(app.appid, app.secret)
	}
	return app.api
}

func (app *Application) signature(signature, timestamp, nonce string) bool {
	strs := sort.StringSlice{app.secret, timestamp, nonce}
	sort.Strings(strs)
	str := ""

	for _, s := range strs {
		str += s
	}

	h := sha1.New()
	h.Write([]byte(str))

	signature_now := fmt.Sprintf("%x", h.Sum(nil))
	if signature == signature_now {
		return true
	}
	return false
}

type Server struct {
	address      string
	applications map[string]*Application
	handles      map[string]Handle
}

func NewServer(address string) *Server {
	return &Server{
		address:      address,
		applications: make(map[string]*Application),
		handles:      make(map[string]Handle),
	}
}

func (srv *Server) Application(app *Application, handle Handle) *Server {
	srv.applications[app.uri] = app
	srv.handles[app.uri] = handle
	return srv
}

func (srv *Server) Start() {
	for uri, app := range srv.applications {
		if app.menu != nil {
			if err := app.Api().DeleteMenu(); err != nil {
				log.Printf("wechat server started faild: %s\n", err.Error())
				return
			}
			if err := app.Api().CreateMenu(app.menu); err != nil {
				log.Printf("wechat server started faild: %s\n", err.Error())
				return
			}
		}
		iris.Get(uri, srv.Get)
		iris.Post(uri, srv.Post)
	}
	iris.Listen(srv.address)
}

func (srv *Server) Get(c *iris.Context) {
	if app, ok := srv.applications[c.PathString()]; ok {
		signature := c.Param("signature")
		timestamp := c.Param("timestamp")
		nonce := c.Param("nonce")
		echostr := c.Param("echostr")

		if app.signature(signature, timestamp, nonce) == true {
			c.Write(echostr)
			return
		}
	}
	c.NotFound()
	return
}

func (srv *Server) Post(c *iris.Context) {
	if app, ok := srv.applications[c.PathString()]; ok {
		var request entry.Request
		if err := c.ReadXML(&request); err == nil {
			if handle, ok := srv.handles[c.PathString()]; ok {
				//! customer service
				if app.customer {
					go handle(app, &request)
					resp := entry.NewTransferToCustomerService(request.FromUserName, request.ToUserName, time.Now().Unix(), "")
					c.XML(200, resp)
					return
				} else {
					if resp, err := handle(app, &request); err == nil {
						if resp != nil {
							c.XML(200, resp)
							return
						}
					}
				}
			}
		}
	}
	c.Write("success")
	return
}
