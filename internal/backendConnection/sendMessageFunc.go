package backendconnection

import (
	"MessengerClient/internal/utils"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

var SendMessageFunc func(string) = func(msg string) {
	if utils.CurrentSessionData.CurrentChan != "" && utils.CurrentSessionData.CurrentWSConn != nil {
		data := utils.MessegeMarshal(msg, 1)
		err := utils.CurrentSessionData.CurrentWSConn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Fatalf("Write websocket messege error: %v", err)
		}
	} else {
		fmt.Println("Please connect to some channel")
	}
}
