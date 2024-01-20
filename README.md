# go-dingtalk-sdk-wrapper
go-dingtalk-sdk-wrapper

wrapper https://github.com/alibabacloud-go/dingtalk to used easily

如果官方已经有最新的 sdk可以实现，这里只需加入功能后给出示例即可，不需要再实现

```bash
go get github.com/xops-infra/go-dingtalk-sdk-wrapper
```

### Usage
1. 钉钉机器人：
    - （http）发送群消息[robot_test.go](examples/robot/main_test.go)
2. 事件订阅 Stream
    - （官方库）Stream方式问答机器人和事件订阅[robot_stream_test.go](examples/dingtalk_stream/main_test.go)
3. 工单相关 Workflow
    - （http）[workflow_test.go](examples/workflow/main_test.go)
