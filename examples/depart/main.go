package main

import (
	"fmt"
	"os"

	dt "github.com/xops-infra/go-dingtalk-sdk-wrapper"
)

var client *dt.DingTalkClient

func init() {
	fmt.Println(os.Getenv("dingtalk_id"), os.Getenv("dingtalk_secret"))
	if os.Getenv("dingtalk_id") == "" || os.Getenv("dingtalk_secret") == "" {
		panic("dingtalk_id or dingtalk_secret is empty, please set env dingtalk_id and dingtalk_secret")
	}
	client, _ = dt.NewDingTalkClient(&dt.DingTalkConfig{
		AppKey:    os.Getenv("dingtalk_id"),
		AppSecret: os.Getenv("dingtalk_secret"),
	})
	client.WithDepartClient().WithUserClient()
}

func main() {
	token := client.AccessToken.Token
	// err recover
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	usersAll := make(map[string]*dt.UserInfo, 0)
	allDepartmentIDs, err := client.Depart.GetAllDepartmentIDs(token)
	if err != nil {
		panic(err)
	}

	fmt.Println("all department ids' length: ", len(allDepartmentIDs), "all department ids: ", allDepartmentIDs)

	for _, departmentID := range allDepartmentIDs {
		users, err := client.User.GetUsers(&dt.GetUsersInput{
			DeptID: departmentID,
			Size:   100,
			Cursor: 0,
		}, token)
		if err != nil {
			panic(err)
		}
		for _, v := range users {
			usersAll[v.Email] = v
		}
	}
	fmt.Println("users count:", len(usersAll), "users: \n", usersAll)
}
