package entry

type Button struct {
	Type    string    `json:"type,omitempty"`
	Name    string    `json:"name"`
	Key     string    `json:"key,omitempty"`
	Url     string    `json:"url,omitempty"`
	Buttons []*Button `json:"sub_button,omitempty"`
}

func NewButton(caption string) *Button {
	return &Button{
		Name: caption,
	}
}

func (b *Button) URL(url string) *Button {
	b.Type = "view"
	b.Url = url
	return b
}

func (b *Button) Event(event string) *Button {
	b.Type = "event"
	b.Key = event
	return b
}

func (b *Button) Append(btn *Button) *Button {
	b.Buttons = append(b.Buttons, btn)
	return b
}

type Menu struct {
	Buttons []*Button `json:"button,omitempty"`
}

func NewMenu(btns ...*Button) *Menu {
	var buttons []*Button
	for _, btn := range btns {
		buttons = append(buttons, btn)
	}
	return &Menu{Buttons: buttons}
}

func (m *Menu) Append(btn *Button) {
	m.Buttons = append(m.Buttons, btn)
}
