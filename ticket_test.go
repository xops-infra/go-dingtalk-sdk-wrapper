package go_dingtalk_sdk_wrapper

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var (
	ticketCli *WorkflowClient
)

func init() {
	config := DingTalkConfig{
		AppKey:    os.Getenv("dingtalk_id"),
		AppSecret: os.Getenv("dingtalk_secret"),
	}
	wclient, _ := InitWorkflowClientV2()
	ticketCli = NewWorkflowClient(wclient, config)
}

func TestTickets(t *testing.T) {
	os.Setenv("task_dingtalk_process_code", "PROC-8FF4478A-DC8D-4922-9B07-")
	input := &GetTicketInput{
		ProcessCode: os.Getenv("task_dingtalk_process_code"),
		StartTime:   time.Now().AddDate(0, 0, -1).UnixMilli(),
		EndTime:     time.Now().UnixMilli(),
	}
	ids := ticketCli.GetTickets(input)
	t.Log(len(ids))
	for _, v := range ids {
		fmt.Println(v)
		if ticketCli.IsTicketApproved(v) {
			t.Log("approved")
		} else {
			t.Log("not approved")
		}
	}
}

// testGetTickets tests the GetTickets method.
func TestGetTickets(t *testing.T) {
	processID := "l992jCwcRuiY93CAh_xzkw0756---"
	comment := Comment{
		ProcessID:     processID,
		Comment:       "test",
		AlertPersion:  "",
		CommentUserID: "290702421220925755---",
	}
	if ticketCli.IsTicketApproved(processID) {
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
	processInstance, err := ticketCli.GetTicket(processID)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*processInstance)
}
