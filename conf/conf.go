package conf

import (
	"fmt"
)

const HOST = "api.weixin.qq.com"

var URIs map[string]string

func MakeURL(key string) string {
	if uri, ok := URIs[key]; ok {
		return fmt.Sprintf("https://%s%s", HOST, uri)
	}
	return fmt.Sprintf("https://%s", HOST)
}

func init() {
	URIs = make(map[string]string)
	URIs["access.token"] = "/cgi-bin/token"
	URIs["callback.ip"] = "/cgi-bin/getcallbackip"
	URIs["menu.create"] = "/cgi-bin/menu/create"
	URIs["menu.get"] = "/cgi-bin/menu/get"
	URIs["menu.delete"] = "/cgi-bin/menu/delete"
	URIs["user.info"] = "/cgi-bin/user/info"
}
