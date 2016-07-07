package entry

import "fmt"

type ApiError struct {
	Code int64  `json:"errcode"`
	Msg  string `json:"errmsg"`
}

func (e ApiError) Error() string {
	return fmt.Sprintf("code: %d message: %s", e.Code, e.Msg)
}
