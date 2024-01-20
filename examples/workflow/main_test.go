package go_dingtalk_sdk_wrapper

import (
	"fmt"
	"os"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/joho/godotenv"

	dt "github.com/xops-infra/go-dingtalk-sdk-wrapper"
)

var client *dt.DingTalkClient

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	client, _ = dt.NewDingTalkClient(&dt.DingTalkConfig{
		AppKey:    os.Getenv("dingtalk_id"),
		AppSecret: os.Getenv("dingtalk_secret"),
	})
	client.WithWorkflowClientV2()
}

func TestCreateProcessInstance(t *testing.T) {
	id, err := client.Workflow.CreateProcessInstance(&dt.CreateProcessInstanceInput{
		ProcessCode:      os.Getenv("PROCESS_CODE"),
		OriginatorUserID: os.Getenv("ORIGINATOR_USER_ID"),
		DeptId:           os.Getenv("DEPT_ID"),
		FormComponentValues: []dt.FormComponentValue{
			{
				Name:  tea.String("Teams"),
				Value: tea.String("xxxx"),
			},
		},
	}, client.AccessToken.Token)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(id)
}

func TestGetProcessInstance(t *testing.T) {
	resp, err := client.Workflow.GetProcessInstance(os.Getenv("PROCESS_INSTANCE_ID"), client.AccessToken.Token)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(tea.Prettify(resp))
}
