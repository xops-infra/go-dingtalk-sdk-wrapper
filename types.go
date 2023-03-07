package go_dingtalk_sdk_wrapper

import (
	"fmt"
	"sort"
	"time"

	workflow "github.com/alibabacloud-go/dingtalk/workflow_1_0"
	"github.com/alibabacloud-go/tea/tea"
)

type TokenDetail struct {
	// Dingtalk token
	Token    string `json:"token"`
	CreateAt int64  `json:"create_at"`
	ExpireIn int64  `json:"expire_in"`
}

func (t *TokenDetail) IsExpire() bool {
	timeNow := time.Now().Unix()
	return timeNow > t.CreateAt+t.ExpireIn
}

type ListWorkflowInput struct {
	ProcessCode string
	StartTime   int64
	EndTime     int64
	NextToken   int64
	MaxResults  int64
}

type CommentInput struct {
	ProcessID     string
	Comment       string      //评论内容
	AlertPerson   AlertPerson //通知@多人，  "[周xx](2907024xxxx09257xxxx)[崔xx](303256xxxx8455xxxx)"
	CommentUserID string      //指评论的人
}

type GrantProcessInstanceForDownloadFileInput struct {
	FileId    string
	ProcessID string
}

type AlertPerson map[string]string

func (p AlertPerson) Marshal() string {
	var alterString string
	for k, v := range p {
		alterString += fmt.Sprintf("[%s](%s)", k, v)
	}
	return alterString
}

type Json map[string]interface{}

type ApprovalStatus string

const (
	Running    ApprovalStatus = "RUNNING"
	Completed  ApprovalStatus = "COMPLETED"
	Terminated ApprovalStatus = "TERMINATED"
	New        ApprovalStatus = "NEW"
	Canceled   ApprovalStatus = "CANCELED"
)

type ApprovalResult string

const (
	Agree     ApprovalResult = "agree"
	Refuse    ApprovalResult = "refuse"
	Revoke    ApprovalResult = "revoke"
	Approving ApprovalResult = "approving"
)

// ProcessInstanceResp 重定向工单返回体
type ProcessInstanceResp workflow.GetProcessInstanceResponseBody

type GrantProcessInstanceForDownloadFileResp workflow.GrantProcessInstanceForDownloadFileResponseBody

func (r *ProcessInstanceResp) GetStatus() ApprovalStatus {
	return ApprovalStatus(tea.StringValue(r.Result.Status))
}

func (r *ProcessInstanceResp) GetResult() ApprovalResult {
	// status为COMPLETED且result为同意时
	switch r.GetStatus() {
	case Completed:
		return ApprovalResult(tea.StringValue(r.Result.Result))
	case Terminated:
		return Revoke
	default:
		return Approving
	}
}

func (r *ProcessInstanceResp) IsAgree() bool {
	return r.GetStatus() == Completed && r.GetResult() == Agree
}

func (r *ProcessInstanceResp) getTasks() ApprovalTask {
	return r.Result.Tasks
}

func (r *ProcessInstanceResp) GetApprovedUser() []Json {
	var userIdList []Json
	task := r.getTasks()
	sort.Sort(task)
	for i := 0; i < len(task); i++ {
		if tea.StringValue(task[i].Result) == "AGREE" && tea.StringValue(task[i].Status) == "COMPLETED" {
			userIdList = append(userIdList, Json{"id": tea.StringValue(task[i].UserId), "next": false})
		} else if tea.StringValue(task[i].Result) == "NONE" && tea.StringValue(task[i].Status) == "RUNNING" {
			userIdList = append(userIdList, Json{"id": tea.StringValue(task[i].UserId), "next": true})
		}
	}
	return userIdList
}

// ApprovalTask 审批流程 别名 并自定义根据CreateTime 排序
type ApprovalTask []*workflow.GetProcessInstanceResponseBodyResultTasks

func (t ApprovalTask) Len() int {
	return len(t)
}

func (t ApprovalTask) Less(i, j int) bool {
	return tea.StringValue(t[i].CreateTime) < tea.StringValue(t[j].CreateTime)
}

func (t ApprovalTask) Swap(i, j int) {
	temp := t[i]
	t[i] = t[j]
	t[j] = temp
}

//type Cache interface {
//	Set() error
//	Get() (string, error)
//}
//
//type MemoryCache struct {
//}
//
//func NewMemoryCache() *MemoryCache {
//	return &MemoryCache{}
//}
//
//func (r *MemoryCache) Set() error {
//
//	return nil
//}
//
//func (r *MemoryCache) Get() error {
//	return nil
//}
