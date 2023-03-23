package go_dingtalk_sdk_wrapper

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alibabacloud-go/tea/tea"
)

var (
	processID = "OzF5M2WCTwuiqZhDtB55Og07561678157055"
	// processID = "l992jCwcRuiY93CAh_xzkw07561677635418" // 有附件
)

func Init() {
	config := DingTalkConfig{
		AppKey:    os.Getenv("dingtalk_id"),
		AppSecret: os.Getenv("dingtalk_secret"),
	}
	client = NewDingTalkClient(&config).WithWorkflowClient()
	err := client.SetAccessToken()
	if err != nil {
		panic(err)
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
	resp, _ := client.WorkflowClient.GetProcessInstance(processID)
	if resp.IsAgree() {
		t.Log("approved")
	} else {
		t.Log("not approved")
	}
	err := client.WorkflowClient.AddProcessInstancedComment(&comment)
	t.Log(err)
}

// test list
func TestListProcessInstance(t *testing.T) {
	resp := client.WorkflowClient.ListProcessInstanceIds(&ListWorkflowInput{
		ProcessCode: "PROC-8FF4478A-DC8D-4922-9B07-36DE233B1DD5",
		StartTime:   time.Now().AddDate(0, 0, -1).UnixMilli(),
		EndTime:     time.Now().UnixMilli(),
		MaxResults:  10,
	})
	t.Log(tea.Prettify(resp))
}

// test GetProcessInstance
func TestGetProcessInstance(t *testing.T) {
	resp, err := client.WorkflowClient.GetProcessInstance(processID)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp.GetAttachmentFileIDs())
}

// test GrantProcessInstanceForDownloadFile
func TestGetAttachment(t *testing.T) {
	resp, err := client.WorkflowClient.GrantProcessInstanceForDownloadFile(&GrantProcessInstanceForDownloadFileInput{
		ProcessID: processID,
		FileId:    "98285509057",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
	resp2, err := http.Get(*resp.Result.DownloadUri)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp2.StatusCode, resp2.Body)
}
