package go_dingtalk_sdk_wrapper

import (
	"sync"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/dingtalk/contact_1_0"
	"github.com/alibabacloud-go/dingtalk/oauth2_1_0"
	"github.com/alibabacloud-go/dingtalk/workflow_1_0"
	"github.com/alibabacloud-go/tea/tea"
)

var (
	openapiConfig *openapi.Config
	AccessToken   string
	authClient    *oauth2_1_0.Client
)

func init() {
	openapiConfig = newOpenaiConfig()
	authClient, _ = oauth2_1_0.NewClient(openapiConfig)
}

func newOpenaiConfig() *openapi.Config {
	return &openapi.Config{
		Protocol: tea.String("https"),
		RegionId: tea.String("central"),
	}
}

type DingTalkConfig struct {
	AppKey    string
	CorpId    string
	AppSecret string
	AgentId   string
}

type DingTalkClient struct {
	OpenapiConfig *openapi.Config
	// AuthClient       *oauth2_1_0.Client
	AccessTokenCache Cache
	AccessToken      string
	Locker           *sync.Mutex
	DingTalkConfig   DingTalkConfig
	// Needed Client
	WorkflowClient *WorkflowClient
}

func NewDingTalkClient(appConfig DingTalkConfig) *DingTalkClient {

	return &DingTalkClient{
		OpenapiConfig:  newOpenaiConfig(),
		Locker:         new(sync.Mutex),
		DingTalkConfig: appConfig,
	}
}

func InitContactClient() (*contact_1_0.Client, error) {
	cclient, err := contact_1_0.NewClient(openapiConfig)
	if err != nil {
		return nil, err
	}
	return cclient, nil
}

func (d *DingTalkClient) NewWorkflowClient() *DingTalkClient {
	client, _ := workflow_1_0.NewClient(openapiConfig)
	d.WorkflowClient = InitWorkflowClient(client)
	return d
}

func GetAccessToken(config DingTalkConfig) (string, error) {
	res, err := authClient.GetAccessToken(&oauth2_1_0.GetAccessTokenRequest{
		AppKey:    tea.String(config.AppKey),
		AppSecret: tea.String(config.AppSecret),
	})
	if err != nil {
		return "", err
	}
	return *res.Body.AccessToken, nil
}

func (d *DingTalkClient) getAccessToken() (*string, error) {
	res, err := authClient.GetAccessToken(&oauth2_1_0.GetAccessTokenRequest{
		AppKey:    tea.String(d.DingTalkConfig.AppKey),
		AppSecret: tea.String(d.DingTalkConfig.AppSecret),
	})
	if err != nil {
		return nil, err
	}
	return res.Body.AccessToken, nil
}

func (d *DingTalkClient) RefreshAccessToken() error {
	d.Locker.Lock()
	//todo cache
	if AccessToken, err := d.AccessTokenCache.Get(); err == nil && AccessToken != "" {
		d.AccessToken = AccessToken
		//todo log
		d.Locker.Unlock()
		return nil
	}

	token, err := d.getAccessToken()

	if err == nil {
		d.AccessToken = tea.StringValue(token)
		AccessToken = tea.StringValue(token)
		d.AccessTokenCache.Set()
		d.Locker.Unlock()
	}
	return err
}
