package wechat

import (
	"github.com/liujianping/api"
	"github.com/liujianping/wechat/conf"
	"github.com/liujianping/wechat/entry"
)

func (c *Client) CreateMenu(m *entry.Menu) error {
	if err := c.Access(nil); err != nil {
		return err
	}

	agent := api.Post(conf.MakeURL("menu.create")).Debug(c.debug)
	agent.QuerySet("access_token", c.token.Secret)
	agent.JSONData(m, true)

	var e entry.ApiError
	if _, _, err := agent.JSON(&e); err != nil {
		return err
	}

	if e.Code != 0 {
		return e
	}
	return nil
}

func (c *Client) DeleteMenu() error {
	if err := c.Access(nil); err != nil {
		return err
	}

	agent := api.Get(conf.MakeURL("menu.delete")).Debug(c.debug)
	agent.QuerySet("access_token", c.token.Secret)

	var e entry.ApiError
	if _, _, err := agent.JSON(&e); err != nil {
		return err
	}

	if e.Code != 0 {
		return e
	}
	return nil
}
