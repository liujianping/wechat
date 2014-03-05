package wechat

import (
	"errors"
	"encoding/xml"
	"net/http"
	"io/ioutil"
	"time"
	"github.com/liujianping/wechat/entry"
)

type CallbackInterface interface {
	Initialize(app *WeChatApp, api	*ApiClient)
	Execute(wr http.ResponseWriter, req *http.Request) error
	MsgText(txt *entry.TextRequest, back chan interface{})
	MsgImage(img *entry.ImageRequest, back chan interface{})
	MsgVoice(voice *entry.VoiceRequest, back chan interface{})
	MsgVideo(video *entry.VideoRequest, back chan interface{})
	MsgLink(link *entry.LinkRequest, back chan interface{})
	Location(location *entry.LocationRequest, back chan interface{})
	EventSubscribe(oid string, back chan interface{})
	EventUnsubscribe(oid string, back chan interface{})
	EventMenu(oid string, key string, back chan interface{})
}

type Callback struct{
	App 	*WeChatApp
	Api		*ApiClient
}
func (cb *Callback) Initialize(app *WeChatApp, api	*ApiClient){
	cb.App = app
	cb.Api = api
}

func (cb *Callback) MsgText(txt *entry.TextRequest, back chan interface{}){

}
func (cb *Callback) MsgImage(img *entry.ImageRequest, back chan interface{}){
	
}
func (cb *Callback) MsgVoice(voice *entry.VoiceRequest, back chan interface{}){
	
}
func (cb *Callback) MsgVideo(video *entry.VideoRequest, back chan interface{}){
	
}
func (cb *Callback) MsgLink(link *entry.LinkRequest, back chan interface{}){
	
}
func (cb *Callback) Location(location *entry.LocationRequest, back chan interface{}){
	
}
func (cb *Callback) EventSubscribe(oid string, back chan interface{}){
	
}
func (cb *Callback) EventUnsubscribe(oid string, back chan interface{}){
	
}
func (cb *Callback) EventMenu(oid string, key string, back chan interface{}){
	
}
func (cb *Callback) Execute(wr http.ResponseWriter, req *http.Request) error{
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	Debug("wechat: data \n", string(data))

	request := &entry.Request{}
	err = xml.Unmarshal(data, request)
	if err != nil {
		return err
	}

	event := request.Event
	msgType := request.MsgType
	ch := make(chan interface{})
	defer close(ch)

	timeout := make(chan bool, 1)
	defer close(timeout)

	go func() {
		time.Sleep(3e9) // 等待3秒钟
		timeout <- true 
		}()
	
	if "event" == msgType {
		//! event
		switch (event){
		case "subscribe":
			go cb.EventSubscribe(request.FromUserName, ch)
		case "unsubscribe":
			go cb.EventUnsubscribe(request.FromUserName, ch)
		case "CLICK":
			go cb.EventMenu(request.FromUserName, request.EventKey, ch)
		case "LOCATION":
			location := &entry.LocationRequest{}
			err = xml.Unmarshal(data, location)
			if err != nil {
				return err
			}
			go cb.Location(location, ch)
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
			go cb.MsgText(text, ch)
		case "image":
			image := &entry.ImageRequest{}
			err = xml.Unmarshal(data, image)
			if err != nil {
				return err
			}
			go cb.MsgImage(image, ch)			
		case "voice":
			voice := &entry.VoiceRequest{}
			err = xml.Unmarshal(data, voice)
			if err != nil {
				return err
			}
			go cb.MsgVoice(voice, ch)			
		case "video":
			video := &entry.VideoRequest{}
			err = xml.Unmarshal(data, video)
			if err != nil {
				return err
			}
			go cb.MsgVideo(video, ch)			
		case "location":
			location := &entry.LocationRequest{}
			err = xml.Unmarshal(data, location)
			if err != nil {
				return err
			}
			go cb.Location(location, ch)			
		case "link":
			link := &entry.LinkRequest{}
			err = xml.Unmarshal(data, link)
			if err != nil {
				return err
			}
			go cb.MsgLink(link, ch)
		}
	}

	select{
	case <-ch:
		response,_ := xml.Marshal(<-ch)
		wr.Write(response)	
	case <-timeout:
		Warn("timeout")
		wr.Write([]byte(""))
	}
	
	return nil
}



