package entry

import "encoding/xml"

const (
	MsgTypeEvent                   = "event"                     // 文本消息
	MsgTypeText                    = "text"                      // 文本消息
	MsgTypeImage                   = "image"                     // 图片消息
	MsgTypeVoice                   = "voice"                     // 语音消息
	MsgTypeVideo                   = "video"                     // 视频消息
	MsgTypeMusic                   = "music"                     // 音乐消息
	MsgTypeNews                    = "news"                      // 图文消息
	MsgTypeTransferCustomerService = "transfer_customer_service" // 将消息转发到多客服
)

const (
	EventSubscribe   = "subscribe"   // 订阅事件
	EventUnSubscribe = "unsubscribe" // 取消订阅事件
	EventScan        = "scan"        // 扫描
	EventLocation    = "location"    // 定位
	EventClick       = "click"       // 自定义菜单点击事件
	EventView        = "view"        // 自定义菜单跳转事件
)

//! msg request referrence: http://mp.weixin.qq.com/wiki/17/fc9a27730e07b9126144d9c96eaf51f9.html
//! evt request referrence: http://mp.weixin.qq.com/wiki/14/f79bdec63116f376113937e173652ba2.html
type Request struct {
	XMLName           xml.Name `xml:"xml"`
	ToUserName        string   `xml:"ToUserName"`
	FromUserName      string   `xml:"FromUserName"`
	CreateTime        int64    `xml:"CreateTime"`
	MsgType           string   `xml:"MsgType"`
	MsgID             string   `xml:"MsgId"`
	MediaID           string   `xml:"MediaId"`
	TextContent       string   `xml:"Content"`
	PictureURL        string   `xml:"PicUrl"`
	VoiceFormat       string   `xml:"Format"`
	VoiceRecognition  string   `xml:"Recognition"`
	VideoThumbMediaID string   `xml:"ThumbMediaId"`
	LocationLabel     string   `xml:"Label"`
	LocationX         float64  `xml:"Location_X"`
	LocationY         float64  `xml:"Location_Y"`
	LocationScale     float64  `xml:"Scale"`
	LinkTitle         string   `xml:"Title"`
	LinkDescription   string   `xml:"Description"`
	LinkURL           string   `xml:"Url"`
	Event             string   `xml:"Event"`
	EventKey          string   `xml:"EventKey"`
	EventTicket       string   `xml:"Ticket"`
	EventLatitude     float64  `xml:"Latitude"`
	EventLongitude    float64  `xml:"Longitude"`
	EventPrecision    float64  `xml:"Precision"`
}
