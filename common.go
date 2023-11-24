package go_dingtalk_sdk_wrapper

import "fmt"

type LowApiError struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode uint   `json:"errcode"`
}

func (e *LowApiError) Error() string {
	return fmt.Sprintf("%d:%s", e.ErrCode, e.ErrMsg)
}
