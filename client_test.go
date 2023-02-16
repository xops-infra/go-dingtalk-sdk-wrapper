package go_dingtalk_sdk_wrapper

import (
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestNewDingTalkClient(t *testing.T) {
	redis := new(redis.Client)
	config := DingTalkConfig{}
	NewDingTalkClient(redis, config).NewWorkflowClient()
}
