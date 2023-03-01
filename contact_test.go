package go_dingtalk_sdk_wrapper

import (
	"fmt"
	"testing"
)

var (
	contantClient *ContactClient
)

func init() {
	cli, err := InitContactClient()
	if err != nil {
		panic(err)
	}
	contantClient = NewContactClient(cli)

}

// test DepartmentList
func TestDepartmentList(t *testing.T) {

	departs, err := contantClient.DepartmentList()
	if err != nil {
		t.Error(err)
	}
	for _, deaprt := range departs {
		fmt.Println(deaprt)
	}
}
