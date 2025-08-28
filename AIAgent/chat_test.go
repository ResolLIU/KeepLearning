package AIAgent

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/coze-dev/coze-go"
	"io"
	"os"
	"strings"
	"testing"
)

func TestChatWithRobot(t *testing.T) {
	botID := "7541825499077885979"
	token := "cztei_l26GTiLKaiw8VCpZNF15gwW9NEUVPoywHkpcKuHhQhlEFOojT5X3HheUijeHsTMcz"
	userID := "丁真"
	scanner := bufio.NewScanner(os.Stdin)
	cov := CreateConversation(token, botID)
	for {
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Println("对话结束，再见！")
			break
		}
		if input == "" {
			continue
		}
		resp := StreamChat(token, botID, cov.ID, userID, input)
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
	}

}
