package go_dingtalk_sdk_wrapper

import (
	"fmt"

	dingtalkworkflow_1_0 "github.com/alibabacloud-go/dingtalk/workflow_1_0"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/spf13/cast"
)

type Ticket struct {
	Client *dingtalkworkflow_1_0.Client
}

type GetTicketInput struct {
	ProcessCode string
	StartTime   int64
	EndTime     int64
}

func NewTicket(cli *dingtalkworkflow_1_0.Client) *Ticket {
	return &Ticket{
		Client: cli,
	}
}

func (t *Ticket) GetTickets(input *GetTicketInput, accessToken string) []string {
	listProcessInstanceIdsHeaders := &dingtalkworkflow_1_0.ListProcessInstanceIdsHeaders{}
	listProcessInstanceIdsHeaders.XAcsDingtalkAccessToken = tea.String(accessToken)
	listProcessInstanceIdsRequest := &dingtalkworkflow_1_0.ListProcessInstanceIdsRequest{
		ProcessCode: tea.String(input.ProcessCode),
		StartTime:   tea.Int64(input.StartTime),
		EndTime:     tea.Int64(input.EndTime),
		MaxResults:  tea.Int64(10),
		NextToken:   tea.Int64(0),
	}
	processIDs := make([]string, 0)
	for {
		res, err := t.Client.ListProcessInstanceIdsWithOptions(listProcessInstanceIdsRequest, listProcessInstanceIdsHeaders, &util.RuntimeOptions{})
		if err != nil {
			fmt.Println(err)
			break
		}
		lists := res.Body.Result.List
		for _, v := range lists {
			processIDs = append(processIDs, *v)
		}
		if res.Body.Result.NextToken == nil {
			break
		}
		listProcessInstanceIdsRequest.NextToken = tea.Int64(cast.ToInt64(*res.Body.Result.NextToken))
	}
	return processIDs
}
