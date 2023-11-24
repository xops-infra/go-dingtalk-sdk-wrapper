package go_dingtalk_sdk_wrapper

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/alibabacloud-go/tea/tea"
)

var (
	// processID = "xxxxx"
	processID = "xxxxx" // 有附件
	client    *DingTalkClient
)

// you should set env: dingtalk_id, dingtalk_secret
//func init() {
//	fmt.Println(os.Getenv("dingtalk_id"), os.Getenv("dingtalk_secret"))
//	client, _ = NewDingTalkClient(&DingTalkConfig{
//		AppKey:    os.Getenv("dingtalk_id"),
//		AppSecret: os.Getenv("dingtalk_secret"),
//	})
//	client.WithWorkflowClient().WithMiniProgramClient(12345)
//}

// Test SendWorkNotification
func TestSendWorkNotification(t *testing.T) {
	req := SendWorkNotificationRequest{
		UseridList: tea.String("xxx"),
		ToAllUser:  tea.Bool(false),
		Msg: &MessageContent{
			MsgType: "text",
			Text: TextBody{
				Content: "this is a test for pop.",
			},
		},
	}
	err := client.MiniProgram.SendWorkNotification(context.Background(), &req, client.AccessToken.Token)
	if err != nil {
		t.Error(err)
	}
}

// Test SendGroupNotification
func TestSendGroupNotification(t *testing.T) {
	req := SendGroupNotificationRequest{
		ChatId: tea.String(""),
		Msg: &MessageContent{
			MsgType: "text",
			Text: TextBody{
				Content: "this is a test for pop.",
			},
		},
	}
	err := client.MiniProgram.SendGroupNotification(context.Background(), &req, client.AccessToken.Token)
	if err != nil {
		t.Error(err)
	}
}

func TestAddComment(t *testing.T) {
	fmt.Println(client.WorkflowClient.tokenDetail)
	comment := CommentInput{
		ProcessID:     processID,
		Comment:       "test1231",
		AlertPerson:   map[string]string{"xx": "xx"},
		CommentUserID: "xx",
	}
	fmt.Println(client.WorkflowClient)
	resp, _ := client.WorkflowClient.GetProcessInstance(context.Background(), processID)
	if resp.IsAgree() {
		t.Log("approved")
	} else {
		t.Log("not approved")
	}
	err := client.WorkflowClient.AddProcessInstancedComment(context.Background(), &comment)
	t.Log(err)
}

// test list
func TestListProcessInstance(t *testing.T) {
	resp, err := client.WorkflowClient.ListProcessInstanceIds(context.Background(), &ListWorkflowInput{
		ProcessCode: "PROC-xxxx",
		StartTime:   time.Now().AddDate(0, 0, -20).UnixMilli(),
		EndTime:     time.Now().UnixMilli(),
		Statuses:    []ApprovalStatus{Completed},
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range resp {
		resp, err := client.WorkflowClient.GetProcessInstance(context.Background(), v)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(tea.Prettify(resp))
		break
	}
}

// test GetProcessInstance
func TestGetProcessInstance(t *testing.T) {
	resp, err := client.WorkflowClient.GetProcessInstance(context.Background(), processID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.GetComment())
}

// test GetAttachmentFileIDs
func TestGetAttachmentFileIDs(t *testing.T) {
	processID = "SrqcxV15SUmsGVeUzO5zkA07561676534364" //人员离职
	resp, err := client.WorkflowClient.GetProcessInstance(context.Background(), processID)
	if err != nil {
		t.Fatal(err)
	}
	ids, _ := resp.GetAttachmentFileIDs()
	for _, v := range ids {
		res, err := client.WorkflowClient.GrantProcessInstanceForDownloadFile(context.Background(), &GrantProcessInstanceForDownloadFileInput{
			FileId:    v.FileID,
			ProcessID: processID,
		})
		if err != nil {
			t.Fatal(err)
		}
		t.Log(res.Result.DownloadUri)
	}
}
