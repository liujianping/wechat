package main

import (
	"log"

	"github.com/liujianping/wechat"
	"github.com/liujianping/wechat/entry"
)

func DemoHandle(app *wechat.Application, request *entry.Request) (interface{}, error) {
	log.Printf("app (%v)\n", app)
	log.Printf("msg (%v)\n", request)
	return nil, nil
}

func main() {
	app := wechat.NewApplication("/demo", "demo_secret", "wx02da1455ece52e5a", "9340ce4b0ab01f33e66dcf9650103fb3", false)

	// btn1 := entry.NewButton("baidu").URL("http://baidu.com")
	// btn2 := entry.NewButton("click").Event("event_click")
	// app.Menu(entry.NewMenu(btn1, btn2))

	serv := wechat.NewServer(":8080")
	serv.Application(app, DemoHandle)

	serv.Start()
}
