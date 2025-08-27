package AIAgent

import (
	"context"
	"fmt"
	"github.com/coze-dev/coze-go"
	"time"
)

func CreateChat(token, botID, userId string, stream bool) coze.Chat {
	authCli := coze.NewTokenAuth(token)
	// Init the Coze client through the access_token.
	cozeCli := coze.NewCozeAPI(authCli, coze.WithBaseURL("https://api.coze.cn"))
	//
	req := &coze.CreateChatsReq{
		BotID:  botID,
		UserID: userId,
		Stream: &stream,
		Messages: []*coze.Message{
			coze.BuildUserQuestionText("我刚问过什么问题?", nil),
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := cozeCli.Chat.Create(ctx, req)
	if err != nil {
		fmt.Println("create fail!")
	}
	return resp.Chat
}
