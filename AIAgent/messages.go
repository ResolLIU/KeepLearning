package AIAgent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/coze-dev/coze-go"
	"net/http"
	"time"
)

func CreateMessage(token, conversationID, content string, role coze.MessageRole, contentType coze.MessageContentType) *coze.CreateMessageResp {
	authCli := coze.NewTokenAuth(token)
	cozeCli := coze.NewCozeAPI(authCli, coze.WithBaseURL("https://api.coze.cn"))
	//
	req := &coze.CreateMessageReq{
		ConversationID: conversationID,
		Role:           role,
		Content:        content,
		ContentType:    contentType,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := cozeCli.Conversations.Messages.Create(ctx, req)
	if err != nil {
		fmt.Println("create fail!")
	}
	return resp
}

func QueryMessageList(token, conversationID string) (list []coze.Message) {
	authCli := coze.NewTokenAuth(token)
	// Init the Coze client through the access_token.
	cozeCli := coze.NewCozeAPI(authCli, coze.WithBaseURL("https://api.coze.cn"))
	req := &coze.ListConversationsMessagesReq{
		ConversationID: conversationID,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := cozeCli.Conversations.Messages.List(ctx, req)
	if err != nil {
		fmt.Println("list fail!")
	}
	for _, i := range resp.Items() {
		list = append(list, *i)
	}
	return
}
func RetrieveMessage(token, conversationId, messageId string) *coze.RetrieveConversationsMessagesResp {
	authCli := coze.NewTokenAuth(token)
	cozeCli := coze.NewCozeAPI(authCli, coze.WithBaseURL("https://api.coze.cn"))
	//
	req := &coze.RetrieveConversationsMessagesReq{
		ConversationID: conversationId,
		MessageID:      messageId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := cozeCli.Conversations.Messages.Retrieve(ctx, req)
	if err != nil {
		fmt.Println("create fail!")
	}
	return resp
}
func DeleteMessagenById(token, conversationId, messageId string) error {
	authCli := coze.NewTokenAuth(token)
	// Init the Coze client through the access_token.
	cozeCli := coze.NewCozeAPI(authCli, coze.WithBaseURL("https://api.coze.cn"))
	req := &coze.DeleteConversationsMessagesReq{
		ConversationID: conversationId,
		MessageID:      messageId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := cozeCli.Conversations.Messages.Delete(ctx, req)
	if err != nil {
		fmt.Println("list fail!")
	}

	return err
}

type ModifyMessageRequest struct {
	Content     string `json:"content"`      // 消息内容
	ContentType string `json:"content_type"` // 内容类型，如"text"
}

// ModifyMessageResponse 定义API响应的结构（根据实际返回调整）
type ModifyMessageResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"` // 可根据实际返回定义更具体的结构
	Message string      `json:"message"`
}

func ModifyCozeMessage(token, conversationID, messageID, content, contentType string) (*ModifyMessageResponse, error) {
	// 构建请求URL，包含查询参数
	url := fmt.Sprintf(
		"https://api.coze.cn/v1/conversation/message/modify?conversation_id=%s&message_id=%s",
		conversationID,
		messageID,
	)
	// 构建请求体 todo：可能会有额外字段
	reqBody := ModifyMessageRequest{
		Content:     content,
		ContentType: contentType,
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("序列化请求体失败: %w", err)
	}

	// 创建POST请求
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
	}

	// 解析响应
	var result ModifyMessageResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return &result, nil
}
