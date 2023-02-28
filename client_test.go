package go_dingtalk_sdk_wrapper

import (
	"fmt"
	"testing"
	"time"
)

var (
	client *DingTalkClient
)

func init() {
	client = NewDingTalkClient(DingTalkConfig{
		AppKey:    "",
		AppSecret: "",
	})
	client.NewWorkflowClient()
	token, err := client.getAccessToken()
	if err != nil {
		panic(err)
	}
	client.AccessToken = *token
}

func TestNewDingTalkClient(t *testing.T) {

	input := &GetTicketInput{
		ProcessCode: "PROC-8FF4478A-DC8D-4922-9B07-36DE233B1DD5",
		StartTime:   time.Now().AddDate(0, 0, -1).UnixMilli(),
		EndTime:     time.Now().UnixMilli(),
	}

	ticketCli := NewTicket(client.WorkflowClient.client)
	ids := ticketCli.GetTickets(input, client.AccessToken)
	t.Log(len(ids))
	for _, v := range ids {
		fmt.Println(v)
	}
}
