package go_dingtalk_sdk_wrapper

/*
官方还没有提供go的sdk
*/

import (
	"github.com/alibabacloud-go/dingtalk/contact_1_0"
	"github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ContactClient struct {
	Client *contact_1_0.Client
	Token  string
}

func NewContactClient(cli *contact_1_0.Client, config DingTalkConfig) *ContactClient {
	token, _ := getAccessToken(config)
	return &ContactClient{
		Client: cli,
		Token:  token,
	}
}

// DepartmentList 403 forbidden
func (c *ContactClient) DepartmentList() ([]*int64, error) {
	searchDepartmentHeaders := &contact_1_0.SearchDepartmentHeaders{
		XAcsDingtalkAccessToken: tea.String(AccessToken),
	}
	searchDepartmentRequest := &contact_1_0.SearchDepartmentRequest{
		QueryWord: tea.String("ops"), // empty string means all
	}
	res, err := c.Client.SearchDepartmentWithOptions(searchDepartmentRequest, searchDepartmentHeaders, &service.RuntimeOptions{})
	if err != nil {
		return []*int64{}, err
	}
	return res.Body.List, nil
}
