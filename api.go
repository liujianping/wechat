package wechat

import (
	"fmt"
	"encoding/json"
	"sort"
	"bytes"
	"crypto/sha1"
	"net/http"
	"io/ioutil"
	"github.com/xuebing1110/wechat/cache"
	"github.com/xuebing1110/wechat/entry"
)

const (
	default_token_key = "wechat.api.default.token.key"
	default_cache_sec = 86400
)

var (
	fmt_token_url string = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	fmt_userinfo_url string = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN"
	fmt_upload_media_url string = "http://file.api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s"
	fmt_download_media_url string = "http://file.api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s"
	fmt_sendmessage_url string = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s"
	fmt_create_menu_url string = "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s"
	fmt_remove_menu_url string = "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=%s"
)

type ApiError struct{
	ErrCode int `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func NewApiError(code int, msg string) *ApiError{
	return &ApiError{ErrCode: code, ErrMsg: msg}
}

func (e *ApiError) Error() string{
	return e.ErrMsg
}

type TokenResponse struct{
	Token string `json:"access_token"`
	Expires_in int64 `json:"expires_in"`
}

type ApiClient struct{
	apptoken string
	appid string
	appsecret string
	cache cache.Cache
}

func NewApiClient(apptoken, appid, appsecret string) *ApiClient {
	return &ApiClient{apptoken: apptoken, appid: appid, appsecret: appsecret}
}

func (c *ApiClient) SetCache(adapter, config string) error {
	mem, err := cache.NewCache(adapter, config)
	if err != nil {
		return err
	}
	c.cache = mem
	return nil
}

func checkJSError(js []byte) error{
	var errmsg ApiError
    if err := json.Unmarshal(js, &errmsg); err != nil {
    	return err
    }

    if errmsg.ErrCode != 0 {
    	return &errmsg
    }

    return nil
}

func (c *ApiClient) Signature(signature, timestamp, nonce string) bool {

	strs := sort.StringSlice{c.apptoken, timestamp, nonce}
	sort.Strings(strs)
	str := ""

	for _,s := range strs {
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

func (c *ApiClient) GetToken() (string, error) {
	if c.cache != nil {
		if v := c.cache.Get(default_token_key); v != nil {
			switch t := v.(type) {
			case string:
				return t, nil
			case []byte:
				return string(t), nil	
			default:
				return "", fmt.Errorf("unexpected type v:", t)
			}
		}
	}

	reponse, err := http.Get(fmt.Sprintf(fmt_token_url, c.appid, c.appsecret))
    if err != nil {
        return "", err
    }

    defer reponse.Body.Close()

    var data []byte
    data, err = ioutil.ReadAll(reponse.Body)
    if err != nil {
    	return "", err
    }

    err = checkJSError(data)
    if err != nil {
    	return "", err
    }

    var tr TokenResponse
    if err = json.Unmarshal(data, &tr); err != nil {
    	return "", err
    }
    if c.cache != nil {
	    c.cache.Put(default_token_key, tr.Token, int64(tr.Expires_in - 10))   	
    }
	
	return tr.Token, nil
}

func (c *ApiClient) Upload() error{
	return nil
}

func (c *ApiClient) Download() error{
	return nil
}

func (c *ApiClient) GetSubscriber(oid string, subscriber *entry.Subscriber) error{
	
	if c.cache != nil {
		if v := c.cache.Get("sub_"+oid); v != nil {
			switch t := v.(type) {
			case []byte:
				if err := json.Unmarshal(t, subscriber); err != nil {
					return err
				}
			}
		}
	}

	token, err := c.GetToken()
	if err != nil {
		return err
	}

	var reponse *http.Response
	reponse, err = http.Get(fmt.Sprintf(fmt_userinfo_url, token, oid))
    if err != nil {
        return err
    }

    defer reponse.Body.Close()

    data, _ := ioutil.ReadAll(reponse.Body)
    err = checkJSError(data)
    if err != nil {
    	return err
    }

    if c.cache != nil {
    	c.cache.Put("sub_" + oid , data, default_cache_sec)
    }
    if err = json.Unmarshal(data, subscriber); err != nil {
		return err
	}

    return nil	
}

func (c *ApiClient) ListSubscribers() error {
	return nil
}

func (c *ApiClient) CreateMenu(menu *entry.Menu) error {
	token, err := c.GetToken()
	if err != nil {
		return err
	}

	data, err := json.Marshal(menu)
	if err != nil {
		return err
	}

	return c.Post(fmt.Sprintf(fmt_create_menu_url, token), data)
}

func (c *ApiClient) GetMenu() error {
	return nil
}

func (c *ApiClient) RemoveMenu() error {
	token, err := c.GetToken()
	if err != nil {
		return err
	}

	reponse, err := http.Get(fmt.Sprintf(fmt_remove_menu_url, token))
    if err != nil {
        return err
    }
    defer reponse.Body.Close()

    data, _ := ioutil.ReadAll(reponse.Body)
    err = checkJSError(data)
    if err != nil {
    	return err
    }

    return nil
}

func (c *ApiClient) Post(url string, json []byte) error {
	reponse, err := http.Post(url, "text/json", bytes.NewBuffer(json))
    if err != nil {
        return err
    }

    defer reponse.Body.Close()

    data, _ := ioutil.ReadAll(reponse.Body)
    err = checkJSError(data)
    if err != nil {
    	return err
    }

    return nil
}

func (c *ApiClient) SendMessage(msg []byte) error {	
	token, err := c.GetToken()
	if err != nil {
		return err
	}
	return c.Post(fmt.Sprintf(fmt_sendmessage_url, token), msg)
}

func (c *ApiClient) SendTextMessage(text *entry.TextMessage) error {
	msg, err := json.Marshal(text)
	if err != nil {
		return err
	}

	return c.SendMessage(msg)
}

func (c *ApiClient) SendImageMessage(image *entry.ImageMessage) error {
	msg, err := json.Marshal(image)
	if err != nil {
		return err
	}

	return c.SendMessage(msg)
}

func (c *ApiClient) SendVoiceMessage(voice *entry.VoiceMessage) error {
	msg, err := json.Marshal(voice)
	if err != nil {
		return err
	}

	return c.SendMessage(msg)
}

func (c *ApiClient) SendVideoMessage(video *entry.VideoMessage) error {
	msg, err := json.Marshal(video)
	if err != nil {
		return err
	}

	return c.SendMessage(msg)
}

func (c *ApiClient) SendMusicMessage(music *entry.MusicMessage) error {
	msg, err := json.Marshal(music)
	if err != nil {
		return err
	}

	return c.SendMessage(msg)
}

func (c *ApiClient) SendNewsMessage(news *entry.NewsMessage) error {
	msg, err := json.Marshal(news)
	if err != nil {
		return err
	}

	return c.SendMessage(msg)
}

func (c *ApiClient) ListGroups() error {
	return nil
}

func (c *ApiClient) CreateGroup() error {
	return nil
}
func (c *ApiClient) UpdateGroup() error {
	return nil
}
func (c *ApiClient) RemoveGroup() error {
	return nil
}
func (c *ApiClient) SearchGroup() error {
	return nil
}
func (c *ApiClient) MovetoGroup() error {
	return nil
}

