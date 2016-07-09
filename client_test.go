package wechat

import (
	"log"
	"testing"

	"github.com/liujianping/wechat/entry"
)

func TestGetUserInfo(t *testing.T) {

	client := NewClient("wx02da1455ece52e5a", "9340ce4b0ab01f33e66dcf9650103fb3").Debug(true)

	var user entry.UserInfo
	if err := client.GetUserInfo("o9_Ejs2eLQasNUVFvZtAs2cogCn4", "zh_CN", &user); err != nil {
		t.Errorf("client.GetUserInfo failed: %s", err.Error())
	}

	log.Println("api.GetUserInfo ", user.OpenID, user.NickName, user.HeadImageURL)
}
