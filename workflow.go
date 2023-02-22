package go_dingtalk_sdk_wrapper

import (
	"github.com/alibabacloud-go/dingtalk/workflow_1_0"
)

type WorkflowClient struct {
	client *workflow_1_0.Client
}

func InitWorkflowClient(client *workflow_1_0.Client) *WorkflowClient {
	return &WorkflowClient{
		client: client,
	}
}

func newGetProcessInstanceHeader() *workflow_1_0.GetProcessInstanceHeaders {
	return &workflow_1_0.GetProcessInstanceHeaders{
		XAcsDingtalkAccessToken: AccessToken,
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
	return c.client.GetProcessInstanceWithOptions(newGetProcessInstanceRequest(
		processID), newGetProcessInstanceHeader(), nil)
}

func newTerminateProcessInstanceHeader() *workflow_1_0.TerminateProcessInstanceHeaders {
	return &workflow_1_0.TerminateProcessInstanceHeaders{
		XAcsDingtalkAccessToken: AccessToken,
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
	return c.client.TerminateProcessInstanceWithOptions(newTerminateProcessInstanceRequest(
		processID), newTerminateProcessInstanceHeader(), nil)
}
