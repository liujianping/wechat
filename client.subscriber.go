package wechat

import (
	"github.com/liujianping/api"
	"github.com/liujianping/wechat/conf"
	"github.com/liujianping/wechat/entry"
)

func (c *Client) GetUserInfo(opendid, lang string, user_info *entry.UserInfo) error {
	if err := c.Access(nil); err != nil {
		return err
	}

	agent := api.Post(conf.MakeURL("user.info")).Debug(c.debug)
	agent.QuerySet("access_token", c.token.Secret)
	agent.QuerySet("openid", opendid)
	agent.QuerySet("lang", lang)

	if _, _, err := agent.JSON(&user_info); err != nil {
		return err
	}

	if user_info.ApiError.Code != 0 {
		return user_info.ApiError
	}
	return nil
}
