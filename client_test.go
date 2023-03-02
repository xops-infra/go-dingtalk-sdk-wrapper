package go_dingtalk_sdk_wrapper

import (
	"fmt"
	"testing"
)

var (
	client *DingTalkClient
)

func init() {
	config := DingTalkConfig{
		AppKey:    "",
		AppSecret: "",
	}
	client = NewDingTalkClient(&config).WithWorkflowClient(&config)
}

func TestNewDingTalkClient(t *testing.T) {

	//input := &GetTicketInput{
	//	ProcessCode: "PROC-8FF4478A-DC8D-4922-9B07-***",
	//	StartTime:   time.Now().AddDate(0, 0, -1).UnixMilli(),
	//	EndTime:     time.Now().UnixMilli(),
	//}

	//ticketCli := NewTicket(client.WorkflowClient.Client)
	//ids := ticketCli.GetTickets(input, client.AccessToken)
	//t.Log(len(ids))
	//for _, v := range ids {
	//	fmt.Println(v)
	//}
}

func TestDingTalkClient_SetAccessToken(t *testing.T) {
	err := client.SetAccessToken()
	if err != nil {
		fmt.Println(err)
	}
}
