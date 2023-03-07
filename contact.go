package go_dingtalk_sdk_wrapper

/*
官方还没有提供go的sdk
*/

import (
	contact "github.com/alibabacloud-go/dingtalk/contact_1_0"
	"github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ContactClient struct {
	client      *contact.Client
	tokenDetail *TokenDetail
}

func NewContactClient(client *contact.Client, token *TokenDetail) *ContactClient {
	return &ContactClient{
		client:      client,
		tokenDetail: token,
	}
}

// DepartmentList 403 forbidden
func (c *ContactClient) DepartmentList() ([]*int64, error) {
	searchDepartmentHeaders := &contact.SearchDepartmentHeaders{
		XAcsDingtalkAccessToken: tea.String(c.tokenDetail.Token),
	}
	searchDepartmentRequest := &contact.SearchDepartmentRequest{
		QueryWord: tea.String("ops"), // empty string means all
	}
	res, err := c.client.SearchDepartmentWithOptions(searchDepartmentRequest, searchDepartmentHeaders, &service.RuntimeOptions{})
	if err != nil {
		return []*int64{}, err
	}
	return res.Body.List, nil
}
