package go_dingtalk_sdk_wrapper

import (
	"context"
	"fmt"
	"net/http"
)

/*
官方 SDK 太垃圾了，这里重新封装一下，用 http client 来实现
doc: https://open.dingtalk.com/document/orgapp/create-an-approval-instance
*/

type CreateProcessInstanceInput struct {
	ProcessCode         string               `json:"processCode" binding:"required"`         // 流程模板唯一标识，可在OA管理后台编辑审批表单部分查询
	OriginatorUserID    string               `json:"originatorUserId" binding:"required"`    // 审批实例发起人的userid
	FormComponentValues []FormComponentValue `json:"formComponentValues" binding:"required"` // 审批流表单参数 表单数据内容，控件列表，最大列表长度：150。
	DeptId              string               `json:"deptId"`                                 // 若approvers未传值时（即不直接指定审批人列表），则deptId需必填，若为根部门ID需填-1。
}

type FormComponentValue struct {
	ID            *string `json:"id"`
	ExtValue      *string `json:"extValue"`
	ComponentType *string `json:"componentType"`
	BizAlias      *string `json:"bizAlias"`
	Name          *string `json:"name" binding:"required"`  // 表单控件名称
	Value         *string `json:"value" binding:"required"` // 表单控件值
}

type Workflow interface {
	// 创建审批实例
	CreateProcessInstance(input *CreateProcessInstanceInput, accessToken string) (string, error)
	// 查询审批工单
	GetProcessInstance(processID string, accessToken string) (*ProcessInstanceRespV2, error)
}

type workflowClient struct {
	requestBuilder requestBuilder
}

func NewWorkflowV2(requestBuilder requestBuilder) Workflow {
	return &workflowClient{
		requestBuilder: requestBuilder,
	}
}

type CreateProcessInstanceResponse struct {
	InstanceId string `json:"instanceId"`
}

// 创建审批实例
func (c *workflowClient) CreateProcessInstance(input *CreateProcessInstanceInput, accessToken string) (string, error) {
	var response CreateProcessInstanceResponse
	url := "https://api.dingtalk.com/v1.0/workflow/processInstances"
	build, err := c.requestBuilder.build(context.Background(), http.MethodPost, url, input)
	if err != nil {
		return "", err
	}
	build.Header.Set("x-acs-dingtalk-access-token", accessToken)
	build.Header.Set("Content-Type", "application/json")
	fmt.Println(build.Body)
	err = c.requestBuilder.sendRequest(build, &response)
	if err != nil {
		return "", err
	}
	return response.InstanceId, nil
}

// 查询审批工单
// /v1.0/workflow/processInstances?processInstanceId=a171de6c-8bxxxx
func (c *workflowClient) GetProcessInstance(processID string, accessToken string) (*ProcessInstanceRespV2, error) {
	var response ProcessInstanceRespV2
	url := fmt.Sprintf("https://api.dingtalk.com/v1.0/workflow/processInstances?processInstanceId=%s", processID)
	build, err := c.requestBuilder.build(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	build.Header.Set("x-acs-dingtalk-access-token", accessToken)
	build.Header.Set("Content-Type", "application/json")
	err = c.requestBuilder.sendRequest(build, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
