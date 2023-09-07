package go_dingtalk_sdk_wrapper

import (
	"context"
	"net/http"
)

type Department struct {
	DepartID int64  `json:"depart_id"`
	Name     string `json:"name"`
	ParentID int64  `json:"parent_id"`
	// create_dept_group bool
	CreateDeptGroup bool `json:"create_dept_group"`
	// auto_add_user bool
	AutoAddUser bool `json:"auto_add_user"`
}

type GetDepartmentsIDInput struct {
	// dept_id number
	DeptID int64 `json:"dept_id"`
}

type GetDepartmentsInput struct {
	// dept_id number
	DeptID int64 `json:"dept_id"`
	// language string example:zh_CN,en_US
	Language string `json:"language"`
}

type GetDepartmentsIDResponse struct {
	// request_id
	RequestId string `json:"request_id"`
	// errcode
	ErrCode int `json:"errcode"`
	// errmsg
	ErrMsg string `json:"errmsg"`
	// result
	Result struct {
		DepartIDList []int64 `json:"dept_id_list"`
	} `json:"result"`
}

type GetDepartmentsResponse struct {
	// request_id
	RequestId string `json:"request_id"`
	// errcode
	ErrCode int `json:"errcode"`
	// errmsg
	ErrMsg string `json:"errmsg"`
	// result
	Result []*Department `json:"result"`
}

type departmentClient struct {
	accessToken    string
	requestBuilder requestBuilder
}

type Depart interface {
	// 获取部门列表
	GetDepartments(*GetDepartmentsInput) ([]*Department, error)
	// 获取子部门ID列表
	GetDepartmentIDs(*GetDepartmentsIDInput) ([]int64, error)
}

func NewDepart(requestBuilder requestBuilder, accessToken string) Depart {
	return &departmentClient{
		requestBuilder: requestBuilder,
		accessToken:    accessToken,
	}
}

func (c *departmentClient) GetDepartments(input *GetDepartmentsInput) ([]*Department, error) {
	var departments []*Department
	var response GetDepartmentsResponse
	url := "https://oapi.dingtalk.com/topapi/v2/department/listsub?access_token=" + c.accessToken
	build, err := c.requestBuilder.build(context.Background(), http.MethodPost, url, input)
	if err != nil {
		return nil, err
	}
	err = c.requestBuilder.sendRequest(build, &response)
	if err != nil {
		return nil, err
	}
	departments = response.Result
	return departments, nil
}

func (c *departmentClient) GetDepartmentIDs(input *GetDepartmentsIDInput) ([]int64, error) {
	var subDepartments []int64
	var response GetDepartmentsIDResponse
	url := "https://oapi.dingtalk.com/topapi/v2/department/listsubid?access_token=" + c.accessToken
	build, err := c.requestBuilder.build(context.Background(), http.MethodPost, url, input)
	if err != nil {
		return nil, err
	}
	err = c.requestBuilder.sendRequest(build, &response)
	if err != nil {
		return nil, err
	}
	subDepartments = response.Result.DepartIDList
	return subDepartments, nil
}
