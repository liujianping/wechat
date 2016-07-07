package entry

import "time"

type Token struct {
	ApiError
	Secret   string    `json:"access_token"`
	ExpireIn int64     `json:"expires_in"`
	CreateAt time.Time `json:"-"`
}

type IPList struct {
	ApiError
	items []string `json:"ip_list"`
}
