package go_dingtalk_sdk_wrapper

import (
	"context"
	"net/http"
)

type User interface {
	// 获取部门id的用户信息
	GetUsers(*GetUsersInput, string) ([]*UserInfo, error)
}

type UserInfo struct {
	// userid string
	UserID string `json:"userid"`
	// unionid string
	UnionID string `json:"unionid"`
	// name string
	Name string `json:"name"`
	// avatar
	Avatar string `json:"avatar"`
	// state_code 国际电话区号。
	StateCode string `json:"state_code"`
	// mobile
	Mobile string `json:"mobile"`
	// hide_mobile bool
	HideMobile bool `json:"hide_mobile"`
	// telephone string 分机号
	Telephone string `json:"telephone"`
	// job_number
	JobNumber string `json:"job_number"`
	// title
	Title string `json:"title"`
	// email
	Email string `json:"email"`
	// org_email
	OrgEmail string `json:"org_email"`
	// work_place
	WorkPlace string `json:"work_place"`
	// remark
	Remark string `json:"remark"`
	// dept_id_list
	DeptIDList []int64 `json:"dept_id_list"`
	// dept_order number 员工在部门中的排序。
	DeptOrder int64 `json:"dept_order"`
	// extension string 扩展属性。
	Extension string `json:"extension"`
	// hired_date number 入职时间，Unix时间戳，单位毫秒。
	HiredDate int64 `json:"hired_date"`
	// active bool
	Active bool `json:"active"`
	// admin
	Admin bool `json:"admin"`
	// boss
	Boss bool `json:"boss"`
	// leader
	Leader bool `json:"leader"`
	// exclusive_account 是否企业账号
	ExclusiveAccount bool `json:"exclusive_account"`
}

type UserClient struct {
	requestBuilder requestBuilder
}

func NewUser(requestBuilder requestBuilder) User {
	return &UserClient{
		requestBuilder: requestBuilder,
	}
}

func (c *UserClient) GetUsers(input *GetUsersInput, accessToken string) ([]*UserInfo, error) {
	var users []*UserInfo
	var response GetUsersResponse
	url := "https://oapi.dingtalk.com/topapi/v2/user/list?access_token=" + accessToken
	for {
		build, err := c.requestBuilder.build(context.Background(), http.MethodPost, url, input)
		if err != nil {
			return nil, err
		}
		build.Header.Set("Content-Type", "application/json")
		err = c.requestBuilder.sendRequest(build, &response)
		if err != nil {
			return nil, err
		}
		users = append(users, response.Result.List...)
		if !response.Result.HasMore {
			break
		}
		input.Cursor = response.Result.NextCursor
	}

	return users, nil
}

type GetUsersInput struct {
	// dept_id number
	DeptID int64 `json:"dept_id"`
	// cursor number 分页查询的游标，最开始传0，后续传返回参数中的next_cursor值。
	Cursor int64 `json:"cursor" default:"0"`
	// size number 每页条数。
	Size int64 `json:"size" default:"20"`
}

type GetUsersResponse struct {
	// errcode
	ErrCode int64 `json:"errcode"`
	// errmsg
	ErrMsg string `json:"errmsg"`
	// result
	Result struct {
		// has_more bool
		HasMore bool `json:"has_more"`
		// next_cursor number
		NextCursor int64 `json:"next_cursor"`
		// list
		List []*UserInfo `json:"list"`
	} `json:"result"`
}
