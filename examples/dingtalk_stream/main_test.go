package main

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/chatbot"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/client"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/event"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/payload"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/plugin"
)

var cred *client.AppCredentialConfig

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	cred = &client.AppCredentialConfig{
		ClientId:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	}
}

// 简单的应答机器人实现
func OnChatBotMessageReceived(ctx context.Context, data *chatbot.BotCallbackDataModel) ([]byte, error) {
	replyMsg := []byte(fmt.Sprintf("msg received: [%s]", data.Text.Content))

	chatbotReplier := chatbot.NewChatbotReplier()
	if err := chatbotReplier.SimpleReplyText(ctx, data.SessionWebhook, replyMsg); err != nil {
		return nil, err
	}
	if err := chatbotReplier.SimpleReplyMarkdown(ctx, data.SessionWebhook, []byte("Markdown消息"), replyMsg); err != nil {
		return nil, err
	}

	return []byte(""), nil
}

// 简单的插件处理实现
func OnPluginMessageReceived(ctx context.Context, request *plugin.GraphRequest) (*plugin.GraphResponse, error) {
	response := &plugin.GraphResponse{
		Body: `{"text": "hello world", "content": [{"title": "1", "description": "2", "url":"https://www.zhihu.com/question/626551401"},{"title": "2", "description": "2", "url":"https://www.zhihu.com/question/626551401"}]}`,
	}
	return response, nil
}

// 事件处理
func OnEventReceived(ctx context.Context, df *payload.DataFrame) (frameResp *payload.DataFrameResponse, err error) {
	eventHeader := event.NewEventHeaderFromDataFrame(df)
	fmt.Printf("received event, eventId=[%s] eventBornTime=[%d] eventCorpId=[%s] eventType=[%s] eventUnifiedAppId=[%s] data=[%s]\n",
		eventHeader.EventId,
		eventHeader.EventBornTime,
		eventHeader.EventCorpId,
		eventHeader.EventType,
		eventHeader.EventUnifiedAppId,
		df.Data)

	frameResp = payload.NewSuccessDataFrameResponse()
	if err := frameResp.SetJson(event.NewEventProcessResultSuccess()); err != nil {
		return nil, err
	}

	return
}

func main() {
	cli := client.NewStreamClient(client.WithAppCredential(cred))

	//注册事件类型的处理函数
	cli.RegisterAllEventRouter(OnEventReceived)
	//注册callback类型的处理函数
	cli.RegisterChatBotCallbackRouter(OnChatBotMessageReceived)
	//注册插件的处理函数
	cli.RegisterPluginCallbackRouter(OnPluginMessageReceived)

	err := cli.Start(context.Background())
	if err != nil {
		panic(err)
	}

	defer cli.Close()

	select {}
}

func TestMain(t *testing.T) {
	main()
}
