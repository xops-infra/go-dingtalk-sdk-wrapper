package go_dingtalk_sdk_wrapper

import (
	"encoding/json"
	"fmt"
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
	Statuses    []ApprovalStatus
}

type CommentInput struct {
	ProcessID     string
	Comment       string      //评论内容
	AlertPerson   AlertPerson //通知@多人，  "[周xx](2907024xxxx09257xxxx)[崔xx](303256xxxx8455xxxx)"
	CommentUserID string      //指评论的人
}

type CommentResp struct {
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

// to string
func (s ApprovalStatus) String() string {
	return string(s)
}

// 流程实例状态，未传值代表查询所有状态的实例ID列表。
// NEW：新创建
// RUNNING：审批中
// TERMINATED：被终止
// COMPLETED：完成
// CANCELED：取消
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

type ProcessInstanceResp workflow.GetProcessInstanceResponseBody

type ProcessInstanceRespV2 struct {
	Success bool                  `json:"success"`
	Result  ProcessInstanceResult `json:"result"`
}

// 没有完全写完，需要自己加入解析的字段
type ProcessInstanceResult struct {
	Status              *string              `json:"status"`
	Title               *string              `json:"title"`
	FinishTime          *string              `json:"finishTime"`
	CreateTime          *string              `json:"createTime"`
	Result              *string              `json:"result"`
	BusinessId          *string              `json:"businessId"`
	OperationRecords    []OperationRecord    `json:"operationRecords"`
	Tasks               []ApprovalTask       `json:"tasks"`
	FormComponentValues []FormComponentValue `json:"formComponentValues"`
}

type ApprovalTask struct {
	ActivityId        *string `json:"activityId,omitempty" xml:"activityId,omitempty"`
	CreateTime        *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	FinishTime        *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	MobileUrl         *string `json:"mobileUrl,omitempty" xml:"mobileUrl,omitempty"`
	PcUrl             *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	ProcessInstanceId *string `json:"processInstanceId,omitempty" xml:"processInstanceId,omitempty"`
	Result            *string `json:"result,omitempty" xml:"result,omitempty"`
	Status            *string `json:"status,omitempty" xml:"status,omitempty"`
	TaskId            *int64  `json:"taskId,omitempty" xml:"taskId,omitempty"`
	UserId            *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

type OperationRecord struct {
	UserId      *string           `json:"userId"`
	Date        *string           `json:"date"`
	Type        *string           `json:"type"`
	Result      *string           `json:"result"`
	Remark      *string           `json:"remark"`
	Attachments []AttachmentFiled `json:"attachments"`
	CcUserIds   []string          `json:"ccUserIds"`
}

type GrantProcessInstanceForDownloadFileResp workflow.GrantProcessInstanceForDownloadFileResponseBody

type AttachmentFiled struct {
	SpaceID   string    `json:"spaceId"`
	FileName  string    `json:"fileName"`
	Thumbnail Thumbnail `json:"thumbnail"`
	FileSize  any       `json:"fileSize"`
	FileType  string    `json:"fileType"`
	FileID    string    `json:"fileId"`
}
type Thumbnail struct {
	AuthCode    string `json:"authCode"`
	AuthMediaID string `json:"authMediaId"`
	Rotation    int    `json:"rotation"`
	Width       int    `json:"width"`
	MediaID     string `json:"mediaId"`
	Height      int    `json:"height"`
}

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

func (r *ProcessInstanceResp) GetAttachmentFileIDs() ([]AttachmentFiled, error) {
	attachFileds := make([]AttachmentFiled, 0)
	for _, v := range r.Result.FormComponentValues {
		if tea.StringValue(v.ComponentType) == "DDAttachment" && v.Value != nil {
			attachments := []AttachmentFiled{}
			err := json.Unmarshal([]byte(tea.StringValue(v.Value)), &attachments)
			if err != nil {
				return attachFileds, err
			}
			attachFileds = append(attachFileds, attachments...)
		}
	}
	return attachFileds, nil
}

// drop 用户直接在detail获取操作
func (r *ProcessInstanceResp) GetComment() ([]CommentResp, error) {
	var comments []CommentResp
	for _, v := range r.Result.OperationRecords {
		fmt.Println(v)
	}
	return comments, nil
}

type CommonResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	// task_id
	TaskId int `json:"task_id"` // 发送任务的id
	// request_id
	RequestId string `json:"request_id"` // 请求id
}
