package go_dingtalk_sdk_wrapper

import (
	"fmt"
	"sync"
	"time"

	robot "github.com/alibabacloud-go/dingtalk/robot_1_0"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	oauth "github.com/alibabacloud-go/dingtalk/oauth2_1_0"
	workflow "github.com/alibabacloud-go/dingtalk/workflow_1_0"
	"github.com/alibabacloud-go/tea/tea"
)

var (
	openapiConfig *openapi.Config
)

func newOpenaiConfig() *openapi.Config {
	return &openapi.Config{
		Protocol: tea.String("https"),
		RegionId: tea.String("central"),
	}
}

func init() {
	openapiConfig = newOpenaiConfig()
}

type DingTalkConfig struct {
	AppKey    string
	CorpId    string
	AppSecret string
	AgentId   string
}

type DingTalkClient struct {
	OpenapiConfig *openapi.Config
	AuthClient    *oauth.Client
	//AccessTokenCache
	AccessToken    *TokenDetail
	Locker         *sync.Mutex
	DingTalkConfig *DingTalkConfig
	// Needed Client
	WorkflowClient *WorkflowClient
	RobotClient    *RobotClient
	MiniProgram    MiniProgram
	Workflow       Workflow
	Depart         Depart
	User           User
	requestBuilder requestBuilder
}

func NewDingTalkClient(appConfig *DingTalkConfig) (*DingTalkClient, error) {
	authClient, _ := oauth.NewClient(openapiConfig)

	dingTalkClient := &DingTalkClient{
		OpenapiConfig:  newOpenaiConfig(),
		AuthClient:     authClient,
		Locker:         new(sync.Mutex),
		DingTalkConfig: appConfig,
		requestBuilder: newRequestBuilder(),
	}
	err := dingTalkClient.setAccessToken()
	if err != nil {
		return nil, err
	}
	return dingTalkClient, nil
}

func (d *DingTalkClient) WithWorkflowClient() *DingTalkClient {
	client, _ := workflow.NewClient(openapiConfig)
	d.WorkflowClient = NewWorkflowClient(client, d.AccessToken)
	return d
}

func (d *DingTalkClient) WithRobotClient() *DingTalkClient {
	client, _ := robot.NewClient(openapiConfig)
	d.RobotClient = NewRobotClient(client, d.requestBuilder)
	return d
}

func (d *DingTalkClient) WithMiniProgramClient(agentId int64) *DingTalkClient {
	d.MiniProgram = NewMiniProgram(agentId, d.requestBuilder)
	return d
}

func (d *DingTalkClient) WithWorkflowClientV2() *DingTalkClient {
	d.Workflow = NewWorkflowV2(d.requestBuilder)
	return d
}

func (d *DingTalkClient) WithDepartClient() *DingTalkClient {
	d.Depart = NewDepart(d.requestBuilder)
	return d
}

func (d *DingTalkClient) WithUserClient() *DingTalkClient {
	d.User = NewUser(d.requestBuilder)
	return d
}

func (d *DingTalkClient) setAccessToken() error {
	CreateAt := time.Now().Unix()
	res, err := d.AuthClient.GetAccessToken(&oauth.GetAccessTokenRequest{
		AppKey:    tea.String(d.DingTalkConfig.AppKey),
		AppSecret: tea.String(d.DingTalkConfig.AppSecret),
	})
	if err != nil {
		return fmt.Errorf("获取dingtalk token异常，因为%s", err.Error())
	}
	d.AccessToken = &TokenDetail{
		Token:    tea.StringValue(res.Body.AccessToken),
		ExpireIn: tea.Int64Value(res.Body.ExpireIn) - 10,
		CreateAt: CreateAt,
	}
	return nil
}

func (d *DingTalkClient) CronSetAccessToken() error {
	if !d.AccessToken.IsExpire() {
		fmt.Println("过期了")
		return nil
	}
	err := d.setAccessToken()
	return err
}

func (d *DingTalkClient) SetAccessToken() error {
	d.Locker.Lock()
	defer func() {
		d.Locker.Unlock()
	}()
	if d.AccessToken == nil {
		err := d.setAccessToken()
		if err != nil {
			return err
		}
	}

	if !d.AccessToken.IsExpire() {
		return nil
	}
	err := d.setAccessToken()
	return err
}

func (d *DingTalkClient) WorkflowSvc() *WorkflowClient {
	return d.WorkflowClient
}

func (d *DingTalkClient) RobotSvc() *RobotClient {
	return d.RobotClient
}

//func (d *DingTalkClient) RefreshAccessToken() error {
//	d.Locker.Lock()
//	//todo cache
//	if d.DingTalkConfig.AccessToken != nil {
//		if AccessToken, err := d.AccessTokenCache.Get(); err == nil && AccessToken != "" {
//			d.DingTalkConfig.AccessToken = AccessToken
//			//todo log
//			d.Locker.Unlock()
//			return nil
//		}
//	}
//	token, err := d.GetAccessToken()
//
//	if err == nil {
//		d.DingTalkConfig.AccessToken = tea.StringValue(token)
//		d.AccessTokenCache.Set()
//		d.Locker.Unlock()
//	}
//	return err
//}
