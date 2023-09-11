package main

import (
	"fmt"
	"os"
	"time"

	dt "github.com/patsnapops/go-dingtalk-sdk-wrapper"
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
	// err recover
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	usersAll := make(map[string]*dt.UserInfo, 0)
	departIDChan := make(chan int64, 2000) // 如果发现 chan 夯住了，可以跳大。
	go getDepart(departIDChan)
	for depart := range departIDChan {
		departRes, err := client.Depart.GetDepartmentIDs(&dt.GetDepartmentsIDInput{
			DeptID: depart,
		})
		if err != nil {
			panic(err)
		}
		for _, v := range departRes {
			departIDChan <- v
		}
		users, err := client.User.GetUsers(&dt.GetUsersInput{
			DeptID: depart,
			Size:   20,
			Cursor: 0,
		})
		if err != nil {
			panic(err)
		}
		for _, v := range users {
			usersAll[v.Name] = v
		}
	}
	fmt.Println("users nu:", len(usersAll))
}

func getDepart(c chan int64) {
	input := &dt.GetDepartmentsIDInput{
		DeptID: int64(1),
	}
	departRes, err := client.Depart.GetDepartmentIDs(input)
	if err != nil {
		panic(err)
	}
	for _, v := range departRes {
		c <- v
	}
	for {
		if len(c) == 0 {
			close(c)
		}
		time.Sleep(time.Second)
		// 清理屏幕
		fmt.Println("get depart", len(c))
	}
}
