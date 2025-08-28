package AIAgent

import (
	"context"
	"fmt"
	"github.com/coze-dev/coze-go"
	"time"
)

func CreateChat(token, botID, userId, content string) coze.Chat {
	authCli := coze.NewTokenAuth(token)
	// Init the Coze client through the access_token.
	cozeCli := coze.NewCozeAPI(authCli, coze.WithBaseURL("https://api.coze.cn"))
	//
	req := &coze.CreateChatsReq{
		BotID:  botID,
		UserID: userId,
		Messages: []*coze.Message{
			coze.BuildUserQuestionText(content, nil),
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
func StreamChat(token, botID, conversationId, userId, content string) coze.Stream[coze.ChatEvent] {
	authCli := coze.NewTokenAuth(token)
	// Init the Coze client through the access_token.
	cozeCli := coze.NewCozeAPI(authCli, coze.WithBaseURL("https://api.coze.cn"))
	//
	req := &coze.CreateChatsReq{
		BotID:          botID,
		UserID:         userId,
		ConversationID: conversationId,
		//ConversationID: conversationId,
		Messages: []*coze.Message{
			coze.BuildUserQuestionText(content, nil),
		},
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := cozeCli.Chat.Stream(ctx, req)
	if err != nil {
		fmt.Println("create fail!")
	}
	return resp
}
func ListChat(token, conversationId, chatId string) []*coze.Message {
	authCli := coze.NewTokenAuth(token)
	// Init the Coze client through the access_token.
	cozeCli := coze.NewCozeAPI(authCli, coze.WithBaseURL("https://api.coze.cn"))
	//
	req := &coze.ListChatsMessagesReq{
		ConversationID: conversationId,
		ChatID:         chatId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := cozeCli.Chat.Messages.List(ctx, req)
	if err != nil {
		fmt.Println("create fail!")
	}
	return resp.Messages
}

func ShowChatInfo(token, conversationId, chatId string) *coze.RetrieveChatsResp {
	authCli := coze.NewTokenAuth(token)
	// Init the Coze client through the access_token.
	cozeCli := coze.NewCozeAPI(authCli, coze.WithBaseURL("https://api.coze.cn"))
	//
	req := &coze.RetrieveChatsReq{
		ConversationID: conversationId,
		ChatID:         chatId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := cozeCli.Chat.Retrieve(ctx, req)
	if err != nil {
		fmt.Println("create fail!")
	}
	return resp
}

// 提交工具执行结果

func SubmitTools(token, conversationId, chatId string) *coze.SubmitToolOutputsChatResp {
	authCli := coze.NewTokenAuth(token)
	// Init the Coze client through the access_token.
	cozeCli := coze.NewCozeAPI(authCli, coze.WithBaseURL("https://api.coze.cn"))
	//
	req := &coze.SubmitToolOutputsChatReq{
		ConversationID: conversationId,
		ChatID:         chatId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := cozeCli.Chat.SubmitToolOutputs(ctx, req)
	if err != nil {
		fmt.Println("create fail!")
	}
	return resp
}

// 取消进行中的对话

func CancelChat(token, conversationId, chatId string) *coze.CancelChatsResp {
	authCli := coze.NewTokenAuth(token)
	// Init the Coze client through the access_token.
	cozeCli := coze.NewCozeAPI(authCli, coze.WithBaseURL("https://api.coze.cn"))
	//
	req := &coze.CancelChatsReq{
		ConversationID: conversationId,
		ChatID:         chatId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := cozeCli.Chat.Cancel(ctx, req)
	if err != nil {
		fmt.Println("create fail!")
	}
	return resp
}
