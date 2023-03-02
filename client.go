package go_dingtalk_sdk_wrapper

import (
	"fmt"
	"sync"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/dingtalk/oauth2_1_0"
	"github.com/alibabacloud-go/dingtalk/workflow_1_0"
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
	AppKey      string
	CorpId      string
	AppSecret   string
	AgentId     string
	AccessToken *TokenDetail
}

type DingTalkClient struct {
	OpenapiConfig *openapi.Config
	AuthClient    *oauth2_1_0.Client
	//AccessTokenCache
	Locker         *sync.Mutex
	DingTalkConfig *DingTalkConfig
	// Needed Client
	WorkflowClient *WorkflowClient
}

func NewDingTalkClient(appConfig *DingTalkConfig) *DingTalkClient {
	authClient, _ := oauth2_1_0.NewClient(openapiConfig)
	return &DingTalkClient{
		OpenapiConfig:  newOpenaiConfig(),
		AuthClient:     authClient,
		Locker:         new(sync.Mutex),
		DingTalkConfig: appConfig,
	}
}

func (d *DingTalkClient) WithWorkflowClient(appConfig *DingTalkConfig) *DingTalkClient {
	client, _ := workflow_1_0.NewClient(openapiConfig)
	d.WorkflowClient = NewWorkflowClient(client, appConfig)

	return d
}

func (d *DingTalkClient) setAccessToken() error {
	CreateAt := time.Now().Unix()
	res, err := d.AuthClient.GetAccessToken(&oauth2_1_0.GetAccessTokenRequest{
		AppKey:    tea.String(d.DingTalkConfig.AppKey),
		AppSecret: tea.String(d.DingTalkConfig.AppSecret),
	})
	if err != nil {
		return fmt.Errorf("获取dingtalk token异常，因为%s", err.Error())
	}
	d.DingTalkConfig.AccessToken = &TokenDetail{
		Token:    tea.StringValue(res.Body.AccessToken),
		ExpireIn: tea.Int64Value(res.Body.ExpireIn),
		CreateAt: CreateAt,
	}

	return nil
}

func (d *DingTalkClient) SetAccessToken() error {
	if d.DingTalkConfig.AccessToken == nil {
		err := d.setAccessToken()
		if err != nil {
			return err
		}
	}

	if !d.DingTalkConfig.AccessToken.IsExpire() {
		return nil
	}
	err := d.setAccessToken()
	if err != nil {
		return err
	}
	return nil
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
