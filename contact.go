package go_dingtalk_sdk_wrapper

import (
	"github.com/alibabacloud-go/dingtalk/contact_1_0"
	"github.com/alibabacloud-go/tea/tea"
)

type ContactClient struct {
	Client *contact_1_0.Client
}

func NewContactClient(cli *contact_1_0.Client) *ContactClient {
	return &ContactClient{
		Client: cli,
	}
}

// DepartmentList
func (u *ContactClient) DepartmentList() ([]*int64, error) {
	searchDepartmentRequest := &contact_1_0.SearchDepartmentRequest{
		QueryWord: tea.String(""), // empty string means all
	}
	res, err := u.Client.SearchDepartment(searchDepartmentRequest)
	if err != nil {
		return []*int64{}, err
	}
	return res.Body.List, nil
}
