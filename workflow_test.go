package go_dingtalk_sdk_wrapper

import (
	"fmt"
	"testing"
)

func TestAddComment(t *testing.T) {

	config := DingTalkConfig{}
	client = NewDingTalkClient(&config).WithWorkflowClient()
	err := client.SetAccessToken()
	if err != nil {
		t.Log(err)
	}
	fmt.Println(client.WorkflowClient.tokenDetail)
	processID := "uIXr5HlNQ-u5YWSQB9LXUg07561677132296"
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

	err = client.WorkflowClient.AddProcessInstancedComment(&comment)
	t.Log(err)
}
