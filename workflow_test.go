package go_dingtalk_sdk_wrapper

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/alibabacloud-go/tea/tea"
)

var (
	// processID = "OzF5M2WCTwuiqZhDtB55Og07561678157055"
	processID = "l992jCwcRuiY93CAh_xzkw07561677635418" // 有附件
	client    *DingTalkClient
)

func init() {
	client = NewDingTalkClient(&DingTalkConfig{
		AppKey:    os.Getenv("dingtalk_id"),
		AppSecret: os.Getenv("dingtalk_secret"),
	}).WithWorkflowClient()
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
	fmt.Println("token:", client.WorkflowClient.tokenDetail)
	resp, err := client.WorkflowClient.ListProcessInstanceIds(&ListWorkflowInput{
		ProcessCode: "PROC-8FF4478A-DC8D-4922-9B07-36DE233B1DD5",
		StartTime:   time.Now().AddDate(0, 0, -3).UnixMilli(),
		EndTime:     time.Now().UnixMilli(),
		Statuses:    []ApprovalStatus{Completed},
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range resp {
		fmt.Println(v)
		// res, err := client.WorkflowClient.GetProcessInstance(v)
		// if err != nil {
		// 	t.Fatal(err)
		// }
		// fmt.Println(res.IsAgree())
	}
}

// test GetProcessInstance
func TestGetProcessInstance(t *testing.T) {
	resp, err := client.WorkflowClient.GetProcessInstance(processID)
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Println(resp.GetAttachmentFileIDs())
	// t.Log(tea.Prettify(resp))
	fmt.Println(resp.GetComment())
}

// test GetAttachmentFileIDs
func TestGetAttachmentFileIDs(t *testing.T) {
	// processID = "GOrQpcoBQOKCaksdkL5T8A07561678973120" // nil
	// processID = "tlfQopnmTGGA0iiD8p1G5Q07561678936579" // filesize string
	processID = "SrqcxV15SUmsGVeUzO5zkA07561676534364" //人员离职
	resp, err := client.WorkflowClient.GetProcessInstance(processID)
	if err != nil {
		t.Fatal(err)
	}
	ids, _ := resp.GetAttachmentFileIDs()
	for _, v := range ids {
		fmt.Println(tea.Prettify(v))
		res, err := client.WorkflowClient.GrantProcessInstanceForDownloadFile(&GrantProcessInstanceForDownloadFileInput{
			FileId:    v.FileID,
			ProcessID: processID,
		})
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(res.Result.DownloadUri)
	}
}
