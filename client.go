package go_dingtalk_sdk_wrapper

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/dingtalk/oauth2_1_0"
	"github.com/alibabacloud-go/dingtalk/workflow_1_0"
	"github.com/alibabacloud-go/tea/tea"
	"sync"
)

var openapiConfig *openapi.Config

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
	OpenapiConfig    *openapi.Config
	AuthClient       *oauth2_1_0.Client
	AccessTokenCache Cache
	AccessToken      string
	Locker           *sync.Mutex
	DingTalkConfig   DingTalkConfig
	// Needed Client
	WorkflowClient *workflow_1_0.Client
}

func NewDingTalkClient(appConfig DingTalkConfig) *DingTalkClient {
	authClient, _ := oauth2_1_0.NewClient(openapiConfig)

	return &DingTalkClient{
		OpenapiConfig:  newOpenaiConfig(),
		AuthClient:     authClient,
		Locker:         new(sync.Mutex),
		DingTalkConfig: appConfig,
	}
}

func (d *DingTalkClient) NewWorkflowClient() *DingTalkClient {
	client, _ := workflow_1_0.NewClient(openapiConfig)
	d.WorkflowClient = client
	return d
}

func (d *DingTalkClient) getAccessToken() (string, error) {
	res, err := d.AuthClient.GetAccessToken(&oauth2_1_0.GetAccessTokenRequest{
		AppKey:    tea.String(d.DingTalkConfig.AppKey),
		AppSecret: tea.String(d.DingTalkConfig.AppSecret),
	})
	if err != nil {
		return "", err
	}
	return tea.StringValue(res.Body.AccessToken), nil
}

func (d *DingTalkClient) RefreshAccessToken() error {
	d.Locker.Lock()
	//todo cache
	//if token, err := d.AccessTokenCache.Get(); err == nil {
	//	d.accessToken = token
	//	//todo log
	//	d.Locker.Unlock()
	//	return nil
	//}

	token, err := d.getAccessToken()
	fmt.Println(token, err, "here")
	if err == nil {
		d.AccessToken = token
		//d.AccessTokenCache.Set()
		d.Locker.Unlock()
	}
	return err
}
