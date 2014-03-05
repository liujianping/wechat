package entry

import (
	"encoding/xml"
	"time"
	"errors"
)


type Response struct{
	ToUserName string 
	FromUserName string 
	MsgType string
	CreateTime time.Duration 
}

type TextResponse struct{
	XMLName xml.Name `xml:"xml"`
	Response
	Content string	
}

type ImageResponse struct{
	XMLName xml.Name `xml:"xml"`
	Response
	MediaId string	`xml:"Image>MediaId"`
}

type VoiceResponse struct{
	XMLName xml.Name `xml:"xml"`
	Response
	MediaId string	`xml:"Voice>MediaId"`	
}

type VideoResponse struct{
	XMLName xml.Name `xml:"xml"`
	Response
	MediaId string	`xml:"Video>MediaId"`	
	Title	string  `xml:"Video>Title"`
	Description string	`xml:"Video>Description"`
}

type MusicResponse struct{
	XMLName xml.Name `xml:"xml"`
	Response
	Title	string  `xml:"Music>Title"`
	Description string	`xml:"Music>Description"`
	MusicUrl string `xml:"Music>MusicUrl"`
	HQMusicUrl string `xml:"Music>HQMusicUrl"`
	ThumbMediaId string `xml:"Music>ThumbMediaId"`	
}

type NewsResponse struct{
	XMLName xml.Name `xml:"xml"`
	Response
	ArticleCount int
	News Articles	`xml:"Articles"`
}

func NewTextResponse(from string, to string, content string) *TextResponse {
	text := new(TextResponse)
	text.FromUserName = from
	text.ToUserName = to
	text.MsgType = "text"
	text.Content = content
	text.CreateTime = time.Duration(time.Now().Unix())
	return text
}

func NewImageResponse(from string, to string, media string) *ImageResponse {
	image := new(ImageResponse)
	image.FromUserName = from
	image.ToUserName = to
	image.MediaId = media
	image.MsgType = "image"
	image.CreateTime = time.Duration(time.Now().Unix())
	return image
}

func NewVoiceResponse(from string, to string, media string) *VoiceResponse {
	voice := new(VoiceResponse)
	voice.FromUserName = from
	voice.ToUserName = to
	voice.MediaId = media
	voice.MsgType = "voice"
	voice.CreateTime = time.Duration(time.Now().Unix())
	return voice
}

func NewVideoResponse(from string, to string, media string, title string, description string) *VideoResponse {
	video := new(VideoResponse)
	video.FromUserName = from
	video.ToUserName = to
	video.MediaId = media
	video.Title = title
	video.Description = description
	video.MsgType = "video"
	video.CreateTime = time.Duration(time.Now().Unix())
	return video
}

func NewMusicResponse(from, to, title, description, musicurl,hqmusicurl,thumb string) *MusicResponse {
	music := new(MusicResponse)
	music.FromUserName = from
	music.ToUserName = to
	music.MsgType = "music"
	music.Title = title
	music.Description = description
	music.MusicUrl = musicurl
	music.HQMusicUrl = hqmusicurl
	music.ThumbMediaId = thumb
	return music
}

func NewNewsResponse(from string, to string) *NewsResponse {
	news := new(NewsResponse)
	news.FromUserName = from
	news.ToUserName = to
	news.MsgType = "news"
	news.ArticleCount = 0
	return news
}

func (news *NewsResponse) Append(article *Article) error{
	if len(news.News.Items) >= 10 {
		return errors.New("entry NewsResponse: news response append exceed 10 articles already.")
	}

	news.News.Items = append(news.News.Items, article)
	news.ArticleCount = len(news.News.Items)
	return nil
}
