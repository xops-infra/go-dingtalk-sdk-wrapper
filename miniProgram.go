package go_dingtalk_sdk_wrapper

/*
Pkg: 小程序
API: https://open.dingtalk.com/document/orgapp/message-notification-overview?spm=a2q3p.21071111.0.0.5e891cfarMi6pk#title-xth-zgb-ppr
*/

import (
	"context"
	"fmt"
	"net/http"

	"github.com/alibabacloud-go/tea/tea"
)

type SendWorkNotificationRequest struct {
	// agent_id
	AgentId *int64 `json:"agent_id" required:"true"` // 应用agentId。 12345
	// userid_list
	UseridList *string `json:"userid_list" example:"user123,user456"` // 接收者的用户userid列表，最大列表长度：100
	// dept_id_list
	DeptIdList *string `json:"dept_id_list" example:"123,456"` // 接收者的部门id列表。最大列表长度为20
	// to_all_user
	ToAllUser *bool `json:"to_all_user" default:"false"` // 是否发送给企业全部用户 当设置为false时必须指定userid_list或dept_id_list其中一个参数的值。
	// msg { "msgtype": "text", "text": { "content": "请提交日报。" } }
	Msg *MessageContent `json:"msg" required:"true"` // 消息内容，消息类型和样例可参考“消息类型与数据格式”。最长不超过2048个字节
}

type MiniProgram interface {
	// 使用机器人发送工作通知
	SendWorkNotification(ctx context.Context, req *SendWorkNotificationRequest, accessToken string) error
	// 发送到群会话，
	SendGroupNotification(ctx context.Context, req *SendGroupNotificationRequest, accessToken string) error //发送消息到企业群接口相关文档，已于2022年09月23日迁移至历史文档（不推荐）目录。不再支持新应用接入，已接入的应用可以正常调用。
}

type miniProgram struct {
	agentId        int64 // 应用agentId。 12345
	requestBuilder requestBuilder
}

// agentId 应用程序的agentId，后台查看
func NewMiniProgram(agentId int64, requestBuilder requestBuilder) MiniProgram {
	return &miniProgram{
		agentId:        agentId,
		requestBuilder: requestBuilder,
	}
}

func (m *miniProgram) SendWorkNotification(ctx context.Context, req *SendWorkNotificationRequest, accessToken string) error {
	var response CommonResponse
	req.AgentId = tea.Int64(m.agentId)
	url := fmt.Sprintf("https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2?access_token=%s", accessToken)
	build, err := m.requestBuilder.build(context.Background(), http.MethodPost, url, req)
	if err != nil {
		return err
	}
	err = m.requestBuilder.sendRequest(build, &response)
	if err != nil {
		return err
	}
	if response.ErrCode != 0 {
		return fmt.Errorf("send work notification error: %s", response.ErrMsg)
	}
	return nil
}

type SendGroupNotificationRequest struct {
	// chatid
	ChatId *string `json:"chatid" required:"true"` // 群会话的id
	// msg
	Msg *MessageContent `json:"msg" required:"true"` // 消息内容，消息类型和样例可参考“消息类型与数据格式”。最长不超过2048个字节
}

func (m *miniProgram) SendGroupNotification(ctx context.Context, req *SendGroupNotificationRequest, accessToken string) error {
	var response CommonResponse
	url := fmt.Sprintf("https://oapi.dingtalk.com/chat/send?access_token=%s", accessToken)
	build, err := m.requestBuilder.build(context.Background(), http.MethodPost, url, req)
	if err != nil {
		return err
	}
	err = m.requestBuilder.sendRequest(build, &response)
	if err != nil {
		return err
	}
	if response.ErrCode != 0 {
		return fmt.Errorf("send group notification error: %s", response.ErrMsg)
	}
	return nil
}
