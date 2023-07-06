package backendconnection

import (
	ui "MessengerClient/internal/cli"
	"MessengerClient/internal/utils"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func recieveMessege(conn *websocket.Conn, serverIp string) {
	for {
		_, msg, readMsgErr := conn.ReadMessage()
		if readMsgErr != nil {
			log.Printf("Server disconnected")
			return
		}
		messege, msgUnmarshalErr := utils.MessegeUnmarshal(msg)
		if msgUnmarshalErr != nil {
			log.Fatalf("JSON messege unmarshal error: %v", msgUnmarshalErr)
		}
		if messege.MsgType == 3 {
			ui.Work = false
			return
		} else {
			fmt.Printf("%v:<%v> %v:%v\n", messege.ID, time.Unix(messege.Time, 0).Format("2006-01-0215:04:05"), messege.Login, strings.Trim(messege.Msg, "\n"))
		}
	}
}
