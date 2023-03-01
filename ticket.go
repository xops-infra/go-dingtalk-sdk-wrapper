package go_dingtalk_sdk_wrapper

import (
	"fmt"

	"github.com/alibabacloud-go/dingtalk/workflow_1_0"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/spf13/cast"
)

// type Ticket struct {
// 	Client *workflow_1_0.Client
// }

type GetTicketInput struct {
	ProcessCode string
	StartTime   int64
	EndTime     int64
}

type Comment struct {
	ProcessID     string
	AccessToken   string
	Comment       string //评论内容
	AlertPersion  string //通知@多人，这里的内容组装需要自己实现  "[周xx](2907024xxxx09257xxxx)[崔xx](303256xxxx8455xxxx)"
	CommentUserID string //默认评论的人，这里没有app用户所以只能选择某个具体的人，比如开发者，如果离职记得更换否则评论会失败
}

func NewTicket(cli *workflow_1_0.Client) *WorkflowClient {
	return &WorkflowClient{
		Client: cli,
	}
}

func (t *WorkflowClient) GetTickets(input *GetTicketInput, accessToken string) []string {
	listProcessInstanceIdsHeaders := &workflow_1_0.ListProcessInstanceIdsHeaders{}
	listProcessInstanceIdsHeaders.XAcsDingtalkAccessToken = tea.String(accessToken)
	listProcessInstanceIdsRequest := &workflow_1_0.ListProcessInstanceIdsRequest{
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

// isTicketApproved returns true if the ticket is approved.
func (t *WorkflowClient) IsTicketApproved(processID string, accessToken string) bool {
	getProcessInstanceHeaders := &workflow_1_0.GetProcessInstanceHeaders{}
	getProcessInstanceHeaders.XAcsDingtalkAccessToken = tea.String(accessToken)
	getProcessInstanceRequest := &workflow_1_0.GetProcessInstanceRequest{
		ProcessInstanceId: tea.String(processID),
	}
	res, err := t.Client.GetProcessInstanceWithOptions(getProcessInstanceRequest, getProcessInstanceHeaders, &util.RuntimeOptions{})
	if err != nil {
		fmt.Println(err)
		return false
	}
	return *res.Body.Result.Status == "COMPLETED"
}

// addTicketComment adds a comment to the ticket.
func (t *WorkflowClient) AddComment(Comment Comment) error {
	addCommentHeaders := &workflow_1_0.AddProcessInstanceCommentHeaders{}
	addCommentHeaders.XAcsDingtalkAccessToken = tea.String(Comment.AccessToken)
	addCommentRequest := &workflow_1_0.AddProcessInstanceCommentRequest{
		CommentUserId:     tea.String(Comment.CommentUserID),
		ProcessInstanceId: tea.String(Comment.ProcessID),
		Text:              tea.String(fmt.Sprintf("%s %s", Comment.AlertPersion, Comment.Comment)),
	}
	_, err := t.Client.AddProcessInstanceCommentWithOptions(addCommentRequest, addCommentHeaders, &util.RuntimeOptions{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// getTicket
func (t *WorkflowClient) GetTicket(processID, accessToken string) (*ProcessInstance, error) {
	getProcessInstanceHeaders := &workflow_1_0.GetProcessInstanceHeaders{}
	getProcessInstanceHeaders.XAcsDingtalkAccessToken = tea.String(accessToken)
	getProcessInstanceRequest := &workflow_1_0.GetProcessInstanceRequest{
		ProcessInstanceId: tea.String(processID),
	}
	res, err := t.Client.GetProcessInstanceWithOptions(getProcessInstanceRequest, getProcessInstanceHeaders, &util.RuntimeOptions{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &ProcessInstance{res.Body.Result}, nil
}
