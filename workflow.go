package go_dingtalk_sdk_wrapper

import (
	"github.com/alibabacloud-go/dingtalk/workflow_1_0"
	"github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type WorkflowClient struct {
	client *workflow_1_0.Client
	config *DingTalkConfig
}

func NewWorkflowClient(client *workflow_1_0.Client, appConfig *DingTalkConfig) *WorkflowClient {
	return &WorkflowClient{
		client: client,
		config: appConfig,
	}
}

func newGetProcessInstanceHeader(token string) *workflow_1_0.GetProcessInstanceHeaders {
	return &workflow_1_0.GetProcessInstanceHeaders{
		XAcsDingtalkAccessToken: tea.String(token),
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
		processID), newGetProcessInstanceHeader(c.config.AccessToken.Token), &service.RuntimeOptions{})
}

func newTerminateProcessInstanceHeader(token string) *workflow_1_0.TerminateProcessInstanceHeaders {
	return &workflow_1_0.TerminateProcessInstanceHeaders{
		XAcsDingtalkAccessToken: tea.String(token),
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
		processID), newTerminateProcessInstanceHeader(c.config.AccessToken.Token), &service.RuntimeOptions{})
}
