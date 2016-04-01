package entry

import "errors"

type Text struct {
	Body string `json:"content"`
}

type Image struct {
	MediaId string `json:"media_id"`
}

type Voice struct {
	MediaId string `json:"media_id"`
}

type Video struct {
	MediaId string `json:"media_id"`
}

type TextMessage struct {
	ToUser  string `json:"touser"`
	MsgType string `json:"msgtype"`
	Content Text   `json:"text"`
}

func BuildTextMessage(to string, content string) *TextMessage {
	text := TextMessage{ToUser: to, MsgType: "text"}
	text.Content.Body = content
	return &text
}

type ImageMessage struct {
	ToUser  string `json:"touser"`
	MsgType string `json:"msgtype"`
	Content Image  `json:"image"`
}

func BuildImageMessage(to string, media_id string) *ImageMessage {
	image := ImageMessage{ToUser: to, MsgType: "image"}
	image.Content.MediaId = media_id
	return &image
}

type VoiceMessage struct {
	ToUser  string `json:"touser"`
	MsgType string `json:"msgtype"`
	Content Voice  `json:"voice"`
}

func BuildVoiceMessage(to string, media_id string) *VoiceMessage {
	voice := VoiceMessage{ToUser: to, MsgType: "voice"}
	voice.Content.MediaId = media_id
	return &voice
}

type VideoMessage struct {
	ToUser  string `json:"touser"`
	MsgType string `json:"msgtype"`
	Content Video  `json:"video"`
}

func BuildVideoMessage(to string, media_id string) *VideoMessage {
	video := VideoMessage{ToUser: to, MsgType: "video"}
	video.Content.MediaId = media_id
	return &video
}

/*
"music":
    {
      "title":"MUSIC_TITLE",
      "description":"MUSIC_DESCRIPTION",
      "musicurl":"MUSIC_URL",
      "hqmusicurl":"HQ_MUSIC_URL",
      "thumb_media_id":"THUMB_MEDIA_ID"
    }
*/
type Music struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	MusicUrl     string `json:"musicurl"`
	HQMusicUrl   string `json:"hqmusicurl"`
	ThumbMediaId string `json:"thumb_media_id"`
}

type MusicMessage struct {
	ToUser  string `json:"touser"`
	MsgType string `json:"msgtype"`
	Content *Music `json:"music"`
}

func NewMusicMessage(to string) *MusicMessage {
	return &MusicMessage{ToUser: to, MsgType: "music"}
}

func NewMusic(title, description, musicurl, hqmusicurl, thumb_media_id string) *Music {
	return &Music{Title: title, Description: description, MusicUrl: musicurl, HQMusicUrl: hqmusicurl, ThumbMediaId: thumb_media_id}
}

func (msg *MusicMessage) SetMusic(music *Music) {
	msg.Content = music
}

/*
{
    "touser":"OPENID",
    "msgtype":"news",
    "news":{
        "articles": [
         {
             "title":"Happy Day",
             "description":"Is Really A Happy Day",
             "url":"URL",
             "picurl":"PIC_URL"
         },
         {
             "title":"Happy Day",
             "description":"Is Really A Happy Day",
             "url":"URL",
             "picurl":"PIC_URL"
         }
         ]
    }
}
*/
type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PicUrl      string `json:"picurl"`
}

type Articles struct {
	Items []*Article `json:"articles" xml:"item"`
}

type NewsMessage struct {
	ToUser  string   `json:"touser"`
	MsgType string   `json:"msgtype"`
	News    Articles `json:"news"`
}

func NewNewsMessage(to string) *NewsMessage {
	return &NewsMessage{ToUser: to, MsgType: "news"}
}

func NewArticle(title, description, url, picurl string) *Article {
	return &Article{Title: title, Description: description, Url: url, PicUrl: picurl}
}

func (news *NewsMessage) Append(article *Article) error {
	if len(news.News.Items) >= 10 {
		return errors.New("entry message: news message append exceed 10 articles already.")
	}

	news.News.Items = append(news.News.Items, article)
	return nil
}
