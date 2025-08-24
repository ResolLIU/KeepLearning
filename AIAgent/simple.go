package AIAgent

import (
	"context"
	"errors"
	"fmt"
	"github.com/coze-dev/coze-go"
	"io"
	"time"
)

func SingleChat() {
	token := "cztei_lfwNhm5iwvBBB7kMmTsAYbdfuMQMonCCsBQ1XlMYPHIm41BUG11CnUIt9zLLU4zyx"
	botID := "7541825499077885979"
	userID := "llq"
	authCli := coze.NewTokenAuth(token)
	// Init the Coze client through the access_token.
	cozeCli := coze.NewCozeAPI(authCli, coze.WithBaseURL("https://api.coze.cn"))
	//
	// Step one, create chats
	// Call the coze.chats().stream() method to create a chats. The create method is a streaming
	// chats and will return a Flowable ChatEvent. Developers should iterate the iterator to get
	// chats event and handle them.
	// //
	req := &coze.CreateChatsReq{
		BotID:  botID,
		UserID: userID,
		Messages: []*coze.Message{
			coze.BuildUserQuestionText("我刚问过什么问题?", nil),
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := cozeCli.Chat.Stream(ctx, req)
	if err != nil {
		fmt.Printf("Error starting chats: %v\n", err)
		return
	}

	defer resp.Close()
	for {
		event, err := resp.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("Stream finished")
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		if event.Event == coze.ChatEventConversationMessageDelta {
			fmt.Print(event.Message.Content)
		} else if event.Event == coze.ChatEventConversationChatCompleted {
			fmt.Printf("Token usage:%d\n", event.Chat.Usage.TokenCount)
		} else {
			fmt.Printf("\n")
		}
	}

	fmt.Printf("done, log:%s\n", resp.Response().LogID())
}
