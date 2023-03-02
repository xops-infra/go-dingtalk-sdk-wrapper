package go_dingtalk_sdk_wrapper

import (
	"fmt"
	"os"
	"testing"
)

var (
	contantClient *ContactClient
)

func init() {
	cli, err := InitContactClient()
	if err != nil {
		panic(err)
	}
	config := DingTalkConfig{
		AppKey:    os.Getenv("dingtalk_id"),
		AppSecret: os.Getenv("dingtalk_secret"),
	}
	contantClient = NewContactClient(cli, config)
}

// test DepartmentList
func TestDepartmentList(t *testing.T) {

	departs, err := contantClient.DepartmentList()
	if err != nil {
		t.Error(err)
	}
	for _, deaprt := range departs {
		fmt.Println(deaprt)
	}
}
