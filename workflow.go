package go_dingtalk_sdk_wrapper

import "github.com/alibabacloud-go/dingtalk/workflow_1_0"

type workflowClient struct {
	workflowClient *workflow_1_0.Client
}

func InitWorkflowClient(client *workflow_1_0.Client) WorkflowClient {
	return &workflowClient{
		workflowClient: client,
	}
}

type WorkflowClient interface {
}
