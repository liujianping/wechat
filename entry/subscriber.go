package entry

/*
{
    "subscribe": 1,
    "openid": "o6_bmjrPTlm6_2sgVt7hMZOPfL2M",
    "nickname": "Band",
    "sex": 1,
    "language": "zh_CN",
    "city": "广州",
    "province": "广东",
    "country": "中国",
    "headimgurl":    "http://wx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/0",
   "subscribe_time": 1382694957,
   "unionid": " o6_bmasdasdsad6_2sgVt7hMZOPfL"
   "remark": "",
   "groupid": 0
}
*/
const (
	LangZhCN = "zh_CN" // 简体中文
	LangZhTW = "zh_TW" // 繁体中文
	LangEN   = "en"    // 英文
)

const (
	SexUnknown = 0 // 未知
	SexMale    = 1 // 男性
	SexFemale  = 2 // 女性
)

type UserInfo struct {
	ApiError
	IsSubscriber  int64    `json:"subscribe"`         // 用户是否订阅该公众号标识, 值为0时, 代表此用户没有关注该公众号, 拉取不到其余信息
	OpenID        string   `json:"openid"`            // 用户的标识, 对当前公众号唯一
	NickName      string   `json:"nickname"`          // 用户的昵称
	Sex           int64    `json:"sex"`               // 用户的性别, 值为1时是男性, 值为2时是女性, 值为0时是未知
	Language      string   `json:"language"`          // 用户的语言, zh_CN, zh_TW, en
	City          string   `json:"city"`              // 用户所在城市
	Province      string   `json:"province"`          // 用户所在省份
	Country       string   `json:"country"`           // 用户所在国家
	HeadImageURL  string   `json:"headimgurl"`        // 用户头像, 最后一个数值代表正方形头像大小(有0, 46, 64, 96, 132数值可选, 0代表640*640正方形头像), 用户没有头像时该项为空
	SubscribeTime int64    `json:"subscribe_time"`    // 用户关注时间, 为时间戳. 如果用户曾多次关注, 则取最后关注时间
	UnionID       string   `json:"unionid,omitempty"` // 只有在用户将公众号绑定到微信开放平台帐号后, 才会出现该字段.
	Remark        string   `json:"remark"`            // 公众号运营者对粉丝的备注, 公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	GroupID       int64    `json:"groupid"`           // 用户所在的分组ID
	Tags          []string `json:"tagid_list"`
}
