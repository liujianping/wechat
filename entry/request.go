package entry

import (
	"encoding/xml"
	"time"
)

type Request struct{
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   // base struct
	FromUserName string
	MsgType      string
	CreateTime   time.Duration
	Event        string
	EventKey     string
}

type TextRequest struct{
	Request
	Content string
	MsgId int64
}

type ImageRequest struct{
	Request
	PicUrl string
	MediaId string
	MsgId int64
}

type VoiceRequest struct{
	Request
	MediaId string
	Format	string
	Recongnition string
	MsgID int64
}

type VideoRequest struct{
	Request
	MsgType string
	MediaId string
	ThumbMediaId string
	MsgId int64
}

type LocationRequest struct{
	Request
	MsgType string
	Location_X float64
	Location_Y float64
	Scale      float64
	Label      string
	MsgId int64
}

type LinkRequest struct{
	Request
	Title string
	Description string
	Url string
	MsgId int64
}

