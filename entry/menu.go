package entry

const (
	// 下面6个类型(包括view类型)的按钮是在公众平台官网发布的菜单按钮类型
	ButtonTypeText  = "text"
	ButtonTypeImage = "img"
	ButtonTypePhoto = "photo"
	ButtonTypeVideo = "video"
	ButtonTypeVoice = "voice"

	// 上面5个类型的按钮不能通过API设置

	ButtonTypeView  = "view"  // 跳转URL
	ButtonTypeClick = "click" // 点击推事件

	// 下面的按钮类型仅支持微信 iPhone5.4.1 以上版本, 和 Android5.4 以上版本的微信用户,
	// 旧版本微信用户点击后将没有回应, 开发者也不能正常接收到事件推送.
	ButtonTypeScanCodePush    = "scancode_push"      // 扫码推事件
	ButtonTypeScanCodeWaitMsg = "scancode_waitmsg"   // 扫码带提示
	ButtonTypePicSysPhoto     = "pic_sysphoto"       // 系统拍照发图
	ButtonTypePicPhotoOrAlbum = "pic_photo_or_album" // 拍照或者相册发图
	ButtonTypePicWeixin       = "pic_weixin"         // 微信相册发图
	ButtonTypeLocationSelect  = "location_select"    // 发送位置

	// 下面的按钮类型专门给第三方平台旗下未微信认证(具体而言, 是资质认证未通过)的订阅号准备的事件类型,
	// 它们是没有事件推送的, 能力相对受限, 其他类型的公众号不必使用.
	ButtonTypeMediaId     = "media_id"     // 下发消息
	ButtonTypeViewLimited = "view_limited" // 跳转图文消息URL
)

type Menu struct {
	Buttons []*Button `json:"button,omitempty"`
	MenuId  int64     `json:"menuid,omitempty"` // 有个性化菜单时查询接口返回值包含这个字段
}

func NewMenu(btns ...*Button) *Menu {
	var buttons []*Button
	for i, btn := range btns {
		if i < 3 {
			buttons = append(buttons, btn)
		}
	}
	return &Menu{Buttons: buttons}
}

type Button struct {
	Type       string    `json:"type,omitempty"`       // 非必须; 菜单的响应动作类型
	Name       string    `json:"name,omitempty"`       // 必须;  菜单标题
	Key        string    `json:"key,omitempty"`        // 非必须; 菜单KEY值, 用于消息接口推送
	Url        string    `json:"url,omitempty"`        // 非必须; 网页链接, 用户点击菜单可打开链接
	MediaId    string    `json:"media_id,omitempty"`   // 非必须; 调用新增永久素材接口返回的合法media_id
	SubButtons []*Button `json:"sub_button,omitempty"` // 非必须; 二级菜单数组
}

func NewButton(caption string) *Button {
	return &Button{
		Name: caption,
	}
}

func (btn *Button) SubButton(btns ...*Button) *Button {
	btn.SubButtons = append(btn.SubButtons, btns...)
	return btn
}

func (btn *Button) Event(key string) *Button {
	btn.Type = ButtonTypeClick
	btn.Key = key
	return btn
}

func (btn *Button) URL(url string) *Button {
	btn.Type = ButtonTypeView
	btn.Url = url
	return btn
}

func (btn *Button) ScanCodePush(key string) *Button {
	btn.Type = ButtonTypeScanCodePush
	btn.Key = key
	return btn
}

func (btn *Button) ScanCodeWaitMsg(key string) *Button {
	btn.Type = ButtonTypeScanCodeWaitMsg
	btn.Key = key
	return btn
}

func (btn *Button) PicSysPhoto(key string) *Button {
	btn.Type = ButtonTypePicSysPhoto
	btn.Key = key
	return btn
}

func (btn *Button) PicPhotoOrAlbum(key string) *Button {
	btn.Type = ButtonTypePicPhotoOrAlbum
	btn.Key = key
	return btn
}

func (btn *Button) PicWeixin(key string) *Button {
	btn.Type = ButtonTypePicWeixin
	btn.Key = key
	return btn
}

func (btn *Button) LocationSelect(key string) *Button {
	btn.Type = ButtonTypeLocationSelect
	btn.Key = key
	return btn
}

func (btn *Button) MediaID(mediaId string) *Button {
	btn.Type = ButtonTypeMediaId
	btn.MediaId = mediaId
	return btn
}

func (btn *Button) ViewLimited(mediaId string) *Button {
	btn.Type = ButtonTypeViewLimited
	btn.MediaId = mediaId
	return btn
}
