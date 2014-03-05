package entry

import (
	"errors"
)

type Button struct{
	Type string	`json:"type,omitempty"`
	Name string `json:"name"`
	Key  string `json:"key,omitempty"`
	Url  string `json:"url,omitempty"`
	Sub []*Button `json:"sub_button,omitempty"`
}

func NewButton(name string) *Button{
	return &Button{Name:name}
}

func NewViewButton(name, url string) *Button {
	return &Button{Type:"view", Name:name, Url: url}
}

func NewClickButton(name string, key string) *Button{
	return &Button{Type:"click", Name:name, Key: key}	
}

func (btn *Button) Append(subbtn *Button) error{
	if len(btn.Sub) >= 5 {
		return errors.New("button: exceed max 5 sub buttons")
	}
	for _, b := range btn.Sub {
		if b.Name == subbtn.Name {
			return errors.New("button: sub button exist same name button")
		}
	}
	btn.Type = ""
	btn.Key = ""
	btn.Url = ""
	btn.Sub = append(btn.Sub, subbtn)
	return nil
}

type Menu struct{
	Buttons []*Button `json:"button,omitempty"`
}

func NewMenu() *Menu{
	return &Menu{}
}

func (menu *Menu) Add(btn *Button) error{
	if len(menu.Buttons) >= 3 {
		return errors.New("menu: can't add menu button more than 3.")
	}

	menu.Buttons = append(menu.Buttons, btn)
	return nil
}

