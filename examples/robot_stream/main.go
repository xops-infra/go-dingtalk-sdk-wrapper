package main

import (
	"context"
	"fmt"
	"os"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/joho/godotenv"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/chatbot"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/client"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/utils"
)

var cred *client.AppCredentialConfig

const (
	topic = "/v1.0/im/bot/messages/get"
)

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

func OnChatReceive(ctx context.Context, data *chatbot.BotCallbackDataModel) ([]byte, error) {
	fmt.Println(tea.Prettify(data))
	msg := "hello world"
	return []byte(msg), nil
}

func main() {
	cli := client.NewStreamClient(
		client.WithAppCredential(cred),
		client.WithUserAgent(client.NewDingtalkGoSDKUserAgent()),
		client.WithSubscription(utils.SubscriptionTypeKCallback, topic, chatbot.NewDefaultChatBotFrameHandler(OnChatReceive).OnEventReceived),
	)

	err := cli.Start(context.Background())
	if err != nil {
		panic(err)
	}

	defer cli.Close()

	select {}
}
