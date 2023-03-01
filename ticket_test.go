package go_dingtalk_sdk_wrapper

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/alibabacloud-go/dingtalk/workflow_1_0"
)

var (
	config  DingTalkConfig
	wclient *workflow_1_0.Client
)

func init() {
	config = DingTalkConfig{
		AppKey:    os.Getenv("dingtalk_id"),
		AppSecret: os.Getenv("dingtalk_secret"),
	}
	wclient, _ = InitWorkflowClient()
}

func TestTickets(t *testing.T) {
	os.Setenv("task_dingtalk_process_code", "PROC-8FF4478A-DC8D-4922-9B07-xxxx")
	input := &GetTicketInput{
		ProcessCode: os.Getenv("task_dingtalk_process_code"),
		StartTime:   time.Now().AddDate(0, 0, -1).UnixMilli(),
		EndTime:     time.Now().UnixMilli(),
	}
	ticketCli := NewWorkflowClient(wclient, config)
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
	processID := ""
	comment := Comment{
		ProcessID:     processID,
		Comment:       "test",
		AlertPersion:  "",
		CommentUserID: "29070242122092575562",
	}
	ticketCli := NewWorkflowClient(wclient, config)
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
