package go_dingtalk_sdk_wrapper

import (
	"fmt"
	"testing"
)

func TestDingTalkClient_SetAccessToken(t *testing.T) {
	//fmt.Println(client)
	config := DingTalkConfig{}
	client := NewDingTalkClient(&config)
	for {
		err := client.SetAccessToken()
		if err != nil {
			fmt.Println(err)
		}
		t.Log("get token success")
	}
}
