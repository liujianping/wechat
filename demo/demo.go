package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/liujianping/wechat"
	"github.com/liujianping/wechat/entry"
)

func DemoHandle(app *wechat.Application, request *entry.Request) (interface{}, error) {
	switch strings.ToLower(request.MsgType) {
	case entry.MsgTypeEvent:
		switch strings.ToLower(request.Event) {
		case entry.EventSubscribe:
			log.Printf("user (%s) subscribed", request.FromUserName)

			var user_info entry.UserInfo
			if err := app.Api().GetUserInfo(request.FromUserName, entry.LangZhCN, &user_info); err != nil {
				return nil, err
			}

			text := entry.NewText(request.FromUserName,
				request.ToUserName,
				time.Now().Unix(),
				fmt.Sprintf("亲爱的(%s), 谢谢您的关注!", user_info.NickName))
			return text, nil

		case entry.EventUnSubscribe:
			log.Printf("user (%s) unsubscribed", request.FromUserName)
		case entry.EventScan:
			log.Printf("user (%s) scan", request.FromUserName)
		case entry.EventLocation:
			log.Printf("user (%s) location", request.FromUserName)
		case entry.EventClick:
			log.Printf("user (%s) menu click (%s)", request.FromUserName, request.EventKey)
		case entry.EventView:
			log.Printf("user (%s) menu view (%s)", request.FromUserName, request.EventKey)
		}
	case entry.MsgTypeText:
		text := entry.NewText(request.FromUserName, request.ToUserName, time.Now().Unix(), request.TextContent)
		return text, nil
	case entry.MsgTypeImage:
	case entry.MsgTypeVoice:
	case entry.MsgTypeVideo:
	case entry.MsgTypeMusic:
	case entry.MsgTypeNews:
	}
	return nil, nil
}

func main() {
	app := wechat.NewApplication("/demo", "demo_secret", "wx02da1455ece52e5a", "9340ce4b0ab01f33e66dcf9650103fb3", false)

	btn1 := entry.NewButton("链接菜单").URL("https://github.com/liujianping/wechat")
	btn2 := entry.NewButton("点击菜单").Event("EVENT_100")
	btn3 := entry.NewButton("更多").SubButton(btn1, btn2)
	app.Menu(entry.NewMenu(btn1, btn2, btn3))

	serv := wechat.NewServer(":8080").Debug(true)
	serv.Application(app, DemoHandle)

	serv.Start()
}
