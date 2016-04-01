package wechat

import (
	"github.com/xuebing1110/wechat/entry"
)

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
func (cb *Callback) EventSubscribe(appoid string, oid string, back chan interface{}){

}
func (cb *Callback) EventUnsubscribe(appoid string, oid string, back chan interface{}){

}
func (cb *Callback) EventMenu(appoid string, oid string, key string, back chan interface{}){
	
}



