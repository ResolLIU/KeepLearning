package AIAgent

import (
	"fmt"
	"testing"
)

func TestCreateConversation(t *testing.T) {
	token, botID := "cztei_hWDbwkcfvg0JZwMzjSYV7vwcLFJSKTTyUdGq2GBNxFJnpBOWWt5xOc8F1bHPRzbFg", "7541816661469888562"
	cov := CreateConversation(token, botID)
	if cov.ID == "" {
		t.Errorf("fail to create")
	} else {
		fmt.Println("covID is ", cov.ID)
	}
}
func TestCreateConversationAndDeleteIt(t *testing.T) {
	token, botID := "cztei_qNPnAxqht64ZOsoxxSuedlQkQwPsg9giBdhRILktysaaUqdgMhVlvva5tmh9VRRoG", "7541816661469888562"
	cov := CreateConversation(token, botID)
	if cov.ID == "" {
		t.Errorf("fail to create")
		return
	} else {
		fmt.Println("covID is ", cov.ID)
	}
	// 创建了一个新的conversation
	list := QueryConversationList(token, botID)
	var flag bool
	for i := range list {
		if list[i].ID == cov.ID {
			flag = true
		}
	}
	if !flag {
		t.Errorf("no such coversation")
	} else {
		fmt.Println("create OK")
	}
	// 删除之后不应该还有这条信息
	err := DeleteConversationById(token, cov.ID)
	if err != nil {
	}
	list = QueryConversationList(token, botID)
	for i := range list {
		if list[i].ID == cov.ID {
			flag = false
		}
	}
	if !flag {
		t.Errorf("the coversation delete failed")
	} else {
		fmt.Println("delete OK")
	}
}
