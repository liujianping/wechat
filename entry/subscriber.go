package entry

/*
"subscribe": 1, 
    "openid": "o6_bmjrPTlm6_2sgVt7hMZOPfL2M", 
    "nickname": "Band", 
    "sex": 1, 
    "language": "zh_CN", 
    "city": "广州", 
    "province": "广东", 
    "country": "中国", 
    "headimgurl":    "http://wx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/0", 
   "subscribe_time": 1382694957
*/
type Subscriber struct{
	Subscribe int 			`json:"subscribe"`
	Openid string			`json:"openid"`
	Nickname string			`json:"nickname"`	
	Sex int 				`json:"sex"`
	Language string			`json:"language"`
	City string				`json:"city"`
	Province string			`json:"province"`
	Country string			`json:"country"`
	Headimgurl string		`json:"headimgurl"`
	Subscribe_time int64	`json:"subscribe_time"`
}

