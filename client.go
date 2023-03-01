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
}

func NewDingTalkClient(appConfig DingTalkConfig) (*DingTalkClient, error) {
	token, err := getAccessToken(appConfig)
	if err != nil {
		return nil, err
	}
	return &DingTalkClient{
		OpenapiConfig:  newOpenaiConfig(),
		Locker:         new(sync.Mutex),
		DingTalkConfig: appConfig,
		AccessToken:    token,
	}, nil
}

func InitWorkflowClient() (*workflow_1_0.Client, error) {
	wclient, err := workflow_1_0.NewClient(openapiConfig)
	if err != nil {
		return nil, err
	}
	return wclient, nil
}

func InitContactClient() (*contact_1_0.Client, error) {
	cclient, err := contact_1_0.NewClient(openapiConfig)
	if err != nil {
		return nil, err
	}
	return cclient, nil
}

func getAccessToken(config DingTalkConfig) (string, error) {
	res, err := authClient.GetAccessToken(&oauth2_1_0.GetAccessTokenRequest{
		AppKey:    tea.String(config.AppKey),
		AppSecret: tea.String(config.AppSecret),
	})
	if err != nil {
		return "", err
	}
	return *res.Body.AccessToken, nil
}

func (d *DingTalkClient) RefreshAccessToken() error {
	d.Locker.Lock()
	defer d.Locker.Unlock()
	res, err := authClient.GetAccessToken(&oauth2_1_0.GetAccessTokenRequest{
		AppKey:    tea.String(d.DingTalkConfig.AppKey),
		AppSecret: tea.String(d.DingTalkConfig.AppSecret),
	})
	if err != nil {
		return err
	}
	d.AccessToken = *res.Body.AccessToken
	return nil
}
