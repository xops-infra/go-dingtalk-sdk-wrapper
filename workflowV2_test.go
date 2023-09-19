package go_dingtalk_sdk_wrapper_test

import (
	"fmt"
	"os"
	"testing"

	dt "github.com/patsnapops/go-dingtalk-sdk-wrapper"
)

var (
	client *dt.DingTalkClient
)

// you should set env: dingtalk_id, dingtalk_secret
func init() {
	fmt.Println(os.Getenv("dingtalk_id"), os.Getenv("dingtalk_secret"))
	client, _ = dt.NewDingTalkClient(&dt.DingTalkConfig{
		AppKey:    os.Getenv("dingtalk_id"),
		AppSecret: os.Getenv("dingtalk_secret"),
	})
	client.WithWorkflowClientV2()
}

// Test CreateProcessInstance
func TestCreateProcessInstance(t *testing.T) {
	id, err := client.Workflow.CreateProcessInstance(&dt.CreateProcessInstanceInput{
		ProcessCode:      "PROC-B85623B4-A372-4684-BB61-1B7E046Cxxxx",
		OriginatorUserID: "2907024212209257xxxx",
		DeptId:           "9022xxxx",
		FormComponentValues: []dt.FormComponentValue{
			{
				Name:  "工单类型",
				Value: "auto",
			},
			{
				Name:  "资源类型",
				Value: "s3Sync",
			},
			{
				Name:  "其他信息",
				Value: "test by patsnapops.",
			},
		},
	}, client.AccessToken.Token)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(id)
}
