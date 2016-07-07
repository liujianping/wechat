package wechat

import (
	"time"

	"github.com/liujianping/api"
	"github.com/liujianping/wechat/conf"
	"github.com/liujianping/wechat/entry"
)

type Client struct {
	appid  string
	secret string
	token  *entry.Token
	debug  bool
}

func NewClient(appid, secret string) *Client {
	return &Client{
		appid:  appid,
		secret: secret,
		token:  nil,
	}
}

func (c *Client) Debug(flag bool) *Client {
	c.debug = flag
	return c
}

func (c *Client) Access(tk *entry.Token) error {
	if c.token != nil {
		expire := c.token.CreateAt.Add(time.Duration(c.token.ExpireIn) * time.Second)
		if expire.After(time.Now()) {
			if tk != nil {
				tk = c.token
			}
			return nil
		}
	}

	agent := api.Get(conf.MakeURL("access.token")).Debug(c.debug)
	agent.QuerySet("grant_type", "client_credential")
	agent.QuerySet("appid", c.appid)
	agent.QuerySet("secret", c.secret)
	var token entry.Token
	if _, _, err := agent.JSON(&token); err != nil {
		return err
	}
	token.CreateAt = time.Now()
	c.token = &token
	if tk != nil {
		tk = &token
	}
	return nil
}
