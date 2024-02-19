package go_dingtalk_sdk_wrapper

import (
	"context"
	"net/http"
)

type Department struct {
	DepartID int64  `json:"dept_id"`
	Name     string `json:"name"`
	ParentID int64  `json:"parent_id"`
	// create_dept_group bool
	CreateDeptGroup bool `json:"create_dept_group"`
	// auto_add_user bool
	AutoAddUser bool `json:"auto_add_user"`
}

type GetDepartmentDetailInput struct {
	// dept_id number
	DeptID int64 `json:"dept_id"`
	// language string example:zh_CN,en_US
	Language string `json:"language"`
}

type GetDepartmentDetailResponse struct {
	// request_id
	RequestId string `json:"request_id"`
	// errcode
	ErrCode int `json:"errcode"`
	// errmsg
	ErrMsg string `json:"errmsg"`
	// result
	Result Department `json:"result"`
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
	requestBuilder requestBuilder
}

type Depart interface {
	// 获取部门详情
	GetDepartmentDetail(input *GetDepartmentDetailInput, accessToken string) (*Department, error)
	// 获取部门列表
	GetDepartments(input *GetDepartmentsInput, accessToken string) ([]*Department, error)
	// 获取所有部门
	GetAllDepartments(accessToken string) ([]*Department, error)
	// 获取子部门ID列表
	GetDepartmentIDs(input *GetDepartmentsIDInput, accessToken string) ([]int64, error)
	// 获取所有部门 ID
	GetAllDepartmentIDs(accessToken string) ([]int64, error)
}

func NewDepart(requestBuilder requestBuilder) Depart {
	return &departmentClient{
		requestBuilder: requestBuilder,
	}
}

func (c *departmentClient) GetDepartmentDetail(input *GetDepartmentDetailInput, accessToken string) (*Department, error) {
	var response GetDepartmentDetailResponse
	url := "https://oapi.dingtalk.com/topapi/v2/department/get?access_token=" + accessToken
	build, err := c.requestBuilder.build(context.Background(), http.MethodPost, url, input)
	if err != nil {
		return nil, err
	}
	err = c.requestBuilder.sendRequest(build, &response)
	if err != nil {
		return nil, err
	}
	return &response.Result, nil
}

func (c *departmentClient) GetDepartments(input *GetDepartmentsInput, accessToken string) ([]*Department, error) {
	var departments []*Department
	var response GetDepartmentsResponse
	url := "https://oapi.dingtalk.com/topapi/v2/department/listsub?access_token=" + accessToken
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

func (c *departmentClient) GetAllDepartments(accessToken string) ([]*Department, error) {
	return c.getAllDepartments(1, accessToken)
}

func (c *departmentClient) GetDepartmentIDs(input *GetDepartmentsIDInput, accessToken string) ([]int64, error) {
	var subDepartments []int64
	var response GetDepartmentsIDResponse
	url := "https://oapi.dingtalk.com/topapi/v2/department/listsubid?access_token=" + accessToken
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

func (c *departmentClient) GetAllDepartmentIDs(accessToken string) ([]int64, error) {
	return c.getAllDepartmentIDs(1, accessToken)
}

func (c *departmentClient) getAllDepartmentIDs(departID int64, token string) (departmentIDs []int64, err error) {
	var allDepartmentIDs []int64
	if departID == 1 {
		allDepartmentIDs = append(allDepartmentIDs, 1)
	}

	departIDs, err := c.GetDepartmentIDs(&GetDepartmentsIDInput{
		DeptID: departID,
	}, token)
	if err != nil {
		return nil, err
	}

	allDepartmentIDs = append(allDepartmentIDs, departIDs...)

	if len(departIDs) != 0 {
		for _, departID := range departIDs {
			departIDs, err := c.getAllDepartmentIDs(departID, token)
			if err != nil {
				return nil, err
			}

			allDepartmentIDs = append(allDepartmentIDs, departIDs...)
		}
	}

	return allDepartmentIDs, nil
}

func (c *departmentClient) getAllDepartments(departID int64, accessToken string) (departments []*Department, err error) {
	var allDepartments []*Department

	if departID == 1 {
		rootDepartment, err := c.GetDepartmentDetail(&GetDepartmentDetailInput{DeptID: 1, Language: "zh_CN"}, accessToken)
		if err != nil {
			return nil, err
		}
		allDepartments = append(allDepartments, rootDepartment)
	}

	departs, err := c.GetDepartments(&GetDepartmentsInput{
		DeptID:   departID,
		Language: "zh_CN",
	}, accessToken)
	if err != nil {
		return nil, err
	}

	allDepartments = append(allDepartments, departs...)

	if len(departs) != 0 {
		for _, department := range departs {
			subDeparts, err := c.getAllDepartments(department.DepartID, accessToken)
			if err != nil {
				return nil, err
			}

			allDepartments = append(allDepartments, subDeparts...)
		}
	}

	return allDepartments, nil
}
