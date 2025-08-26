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
