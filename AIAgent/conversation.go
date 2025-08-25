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

func CreateConversation(token, botID string) coze.Conversation {
	authCli := coze.NewTokenAuth(token)
	// Init the Coze client through the access_token.
	cozeCli := coze.NewCozeAPI(authCli, coze.WithBaseURL("https://api.coze.cn"))
	//
	req := &coze.CreateConversationsReq{
		BotID: botID,
		Messages: []*coze.Message{
			coze.BuildUserQuestionText("我刚问过什么问题?", nil),
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := cozeCli.Conversations.Create(ctx, req)
	if err != nil {
		fmt.Println("create fail!")
	}
	return resp.Conversation
}

func QueryConversationList(token, botID string) (list []coze.Conversation) {
	authCli := coze.NewTokenAuth(token)
	// Init the Coze client through the access_token.
	cozeCli := coze.NewCozeAPI(authCli, coze.WithBaseURL("https://api.coze.cn"))
	req := &coze.ListConversationsReq{
		BotID:    botID,
		PageNum:  1,
		PageSize: 20,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := cozeCli.Conversations.List(ctx, req)
	if err != nil {
		fmt.Println("list fail!")
	}
	for _, i := range resp.Items() {
		list = append(list, *i)
	}
	return
}

func DeleteConversationById(token, conversationId string) error {
	url := fmt.Sprintf("https://api.coze.cn/v1/conversations/%s", conversationId)

	// 创建请求
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 处理响应
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
	}
	fmt.Printf("对话 %s 删除成功\n", conversationId)
	return nil
}

func ClearConversasionById(token, conversationId string) error {
	authCli := coze.NewTokenAuth(token)
	// Init the Coze client through the access_token.
	cozeCli := coze.NewCozeAPI(authCli, coze.WithBaseURL("https://api.coze.cn"))
	req := &coze.ClearConversationsReq{
		conversationId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := cozeCli.Conversations.Clear(ctx, req)
	if err != nil {
		fmt.Println("list fail!")
	}
	fmt.Println(resp.ID)
	return err
}

// 请求体结构
type updateConversationRequest struct {
	Name string `json:"name"`
}

// 响应体结构（根据实际API返回调整）
type updateConversationResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// 可以根据API文档添加其他字段
}

// 更新会话名
func updateCozeConversation(conversationID, token, newName string) error {
	// 构建请求URL
	url := fmt.Sprintf("https://api.coze.cn/v1/conversations/%s", conversationID)
	// 创建请求体
	reqBody := updateConversationRequest{
		Name: newName,
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("序列化请求体失败: %v", err)
	}
	// 创建PUT请求
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("创建请求失败: %v", err)
	}
	// 设置请求头
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
	}
	// 解析响应（可选，根据需要处理）
	var response updateConversationResponse
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("解析响应失败: %v", err)
	}

	fmt.Printf("对话 %s 更新成功，新名称: %s\n", response.ID, response.Name)
	return nil
}
