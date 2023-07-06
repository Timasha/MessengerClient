package backendconnection

import (
	"MessengerClient/internal/utils"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

var DeleteMessegeFunc func(string) = func(id string) {
	_, err := strconv.ParseUint(strings.Trim(id, "\n"), 10, 64)
	if err != nil {
		fmt.Println("This is not number. Try again.")
	} else if utils.CurrentSessionData.CurrentChan != "" && utils.CurrentSessionData.CurrentWSConn != nil {
		data := utils.MessegeMarshal(strings.Trim(id, "\n"), 0)
		err := utils.CurrentSessionData.CurrentWSConn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Fatalf("Write websocket messege error: %v", err)
		}
	} else {
		fmt.Println("Please connect to some channel")
	}
}
