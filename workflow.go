package go_dingtalk_sdk_wrapper

import (
	"fmt"
	workflow "github.com/alibabacloud-go/dingtalk/workflow_1_0"
	"github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/spf13/cast"
)

type WorkflowClient struct {
	client      *workflow.Client
	tokenDetail *TokenDetail
}

func NewWorkflowClient(client *workflow.Client, token *TokenDetail) *WorkflowClient {
	return &WorkflowClient{
		client:      client,
		tokenDetail: token,
	}
}

// about get one processInstance
func newGetProcessInstanceHeader(token string) *workflow.GetProcessInstanceHeaders {
	return &workflow.GetProcessInstanceHeaders{
		XAcsDingtalkAccessToken: tea.String(token),
	}
}

func newGetProcessInstanceRequest(
	processInstanceID string) *workflow.GetProcessInstanceRequest {
	return &workflow.GetProcessInstanceRequest{
		ProcessInstanceId: &processInstanceID,
	}
}

func (c *WorkflowClient) GetProcessInstance(
	processID string) (*ProcessInstanceResp, error) {
	resp, err := c.client.GetProcessInstanceWithOptions(newGetProcessInstanceRequest(
		processID), newGetProcessInstanceHeader(c.tokenDetail.Token), &service.RuntimeOptions{})

	if err != nil {
		return nil, err
	}
	return (*ProcessInstanceResp)(resp.Body), nil
}

// about terminated process
func newTerminateProcessInstanceHeader(token string) *workflow.TerminateProcessInstanceHeaders {
	return &workflow.TerminateProcessInstanceHeaders{
		XAcsDingtalkAccessToken: tea.String(token),
	}
}

func newTerminateProcessInstanceRequest(
	processInstanceID string) *workflow.TerminateProcessInstanceRequest {
	return &workflow.TerminateProcessInstanceRequest{
		ProcessInstanceId: &processInstanceID,
	}
}

func (c *WorkflowClient) TerminateProcessInstance(
	processID string) (bool, error) {
	resp, err := c.client.TerminateProcessInstanceWithOptions(newTerminateProcessInstanceRequest(
		processID), newTerminateProcessInstanceHeader(c.tokenDetail.Token), &service.RuntimeOptions{})
	if err != nil {
		return false, err
	}
	return tea.BoolValue(resp.Body.Result), err
}

// about list processInstance
func newListProcessInstanceIdsHeaders(token string) *workflow.ListProcessInstanceIdsHeaders {
	return &workflow.ListProcessInstanceIdsHeaders{
		XAcsDingtalkAccessToken: tea.String(token),
	}
}

func newListProcessInstanceIdsRequest(input *ListWorkflowInput) *workflow.ListProcessInstanceIdsRequest {
	return &workflow.ListProcessInstanceIdsRequest{
		ProcessCode: tea.String(input.ProcessCode),
		StartTime:   tea.Int64(input.StartTime),
		EndTime:     tea.Int64(input.EndTime),
		MaxResults:  tea.Int64(input.MaxResults),
		NextToken:   tea.Int64(input.NextToken),
	}
}

func (c *WorkflowClient) ListProcessInstanceIds(input *ListWorkflowInput) []string {
	var processIDs []string
	for {
		res, err := c.client.ListProcessInstanceIdsWithOptions(newListProcessInstanceIdsRequest(input),
			newListProcessInstanceIdsHeaders(c.tokenDetail.Token),
			&service.RuntimeOptions{})

		if err != nil {
			continue
		}

		processIDs = append(processIDs, tea.StringSliceValue(res.Body.Result.List)...)
		if res.Body.Result.NextToken == nil {
			break
		}
		input.NextToken = cast.ToInt64(tea.StringValue(res.Body.Result.NextToken))
	}
	return processIDs
}

// about add comment for processInstance
func newAddProcessInstanceCommentHeaders(token string) *workflow.AddProcessInstanceCommentHeaders {
	return &workflow.AddProcessInstanceCommentHeaders{
		XAcsDingtalkAccessToken: tea.String(token),
	}
}

func newAddProcessInstanceCommentRequest(input *CommentInput) *workflow.AddProcessInstanceCommentRequest {
	return &workflow.AddProcessInstanceCommentRequest{
		CommentUserId:     tea.String(input.CommentUserID),
		ProcessInstanceId: tea.String(input.ProcessID),
		Text:              tea.String(fmt.Sprintf("%s %s", input.AlertPerson.Marshal(), input.Comment)),
	}
}

func (c *WorkflowClient) AddProcessInstancedComment(input *CommentInput) error {
	_, err := c.client.AddProcessInstanceCommentWithOptions(newAddProcessInstanceCommentRequest(input),
		newAddProcessInstanceCommentHeaders(c.tokenDetail.Token), &service.RuntimeOptions{})
	return err
}
