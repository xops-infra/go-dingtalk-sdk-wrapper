package go_dingtalk_sdk_wrapper

import (
	"github.com/alibabacloud-go/dingtalk/workflow_1_0"
	"github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type WorkflowClient struct {
	Client *workflow_1_0.Client
}

func InitWorkflowClient(client *workflow_1_0.Client) *WorkflowClient {
	return &WorkflowClient{
		Client: client,
	}
}

func newGetProcessInstanceHeader() *workflow_1_0.GetProcessInstanceHeaders {
	return &workflow_1_0.GetProcessInstanceHeaders{
		XAcsDingtalkAccessToken: tea.String(AccessToken),
	}
}

func newGetProcessInstanceRequest(
	processInstanceID string) *workflow_1_0.GetProcessInstanceRequest {
	return &workflow_1_0.GetProcessInstanceRequest{
		ProcessInstanceId: &processInstanceID,
	}
}

func (c *WorkflowClient) GetProcessInstance(
	processID string) (*workflow_1_0.GetProcessInstanceResponse, error) {
	return c.Client.GetProcessInstanceWithOptions(newGetProcessInstanceRequest(
		processID), newGetProcessInstanceHeader(), &service.RuntimeOptions{})
}

func newTerminateProcessInstanceHeader() *workflow_1_0.TerminateProcessInstanceHeaders {
	return &workflow_1_0.TerminateProcessInstanceHeaders{
		XAcsDingtalkAccessToken: tea.String(AccessToken),
	}
}

func newTerminateProcessInstanceRequest(
	processInstanceID string) *workflow_1_0.TerminateProcessInstanceRequest {
	return &workflow_1_0.TerminateProcessInstanceRequest{
		ProcessInstanceId: &processInstanceID,
	}
}

func (c *WorkflowClient) TerminateProcessInstance(
	processID string) (*workflow_1_0.TerminateProcessInstanceResponse, error) {
	return c.Client.TerminateProcessInstanceWithOptions(newTerminateProcessInstanceRequest(
		processID), newTerminateProcessInstanceHeader(), &service.RuntimeOptions{})
}
