package go_dingtalk_sdk_wrapper

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var (
	client *DingTalkClient
)

func init() {
	client = NewDingTalkClient(DingTalkConfig{
		AppKey:    os.Getenv("dingtalk_id"),
		AppSecret: os.Getenv("dingtalk_secret"),
	})
	err := client.NewWorkflowClient()
	if err != nil {
		panic(err)
	}
}

func TestTickets(t *testing.T) {
	os.Setenv("task_dingtalk_process_code", "PROC-8FF4478A-DC8D-4922-9B07-xxxx")
	input := &GetTicketInput{
		ProcessCode: os.Getenv("task_dingtalk_process_code"),
		StartTime:   time.Now().AddDate(0, 0, -1).UnixMilli(),
		EndTime:     time.Now().UnixMilli(),
	}
	ticketCli := NewWorkflowClient(client.WorkflowClient.Client)
	ids := ticketCli.GetTickets(input, client.AccessToken)
	t.Log(len(ids))
	for _, v := range ids {
		fmt.Println(v)
		if ticketCli.IsTicketApproved(v, client.AccessToken) {
			t.Log("approved")
		} else {
			t.Log("not approved")
		}
	}
}

// testGetTickets tests the GetTickets method.
func TestGetTickets(t *testing.T) {
	processID := ""
	comment := Comment{
		ProcessID:     processID,
		AccessToken:   client.AccessToken,
		Comment:       "test",
		AlertPersion:  "",
		CommentUserID: "29070242122092575562",
	}
	ticketCli := NewWorkflowClient(client.WorkflowClient.Client)
	if ticketCli.IsTicketApproved(processID, client.AccessToken) {
		t.Log("approved")
	} else {
		t.Log("not approved")
	}
	// test add comment
	if false == true {
		err := ticketCli.AddComment(comment)
		if err != nil {
			t.Error(err)
		}
	}
	processInstance, err := ticketCli.GetTicket(processID, client.AccessToken)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*processInstance)
}
