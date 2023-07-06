package backendconnection

import (
	"MessengerClient/internal/utils"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var ConnectChannelFunc func([]string, string) = func(args []string, serverIp string) {
	if utils.CurrentSessionData.CurrentChan != "" {
		DisconnectFunc()
	}
	var connErr error
	var connectResponse utils.ConnectJson = connectChannelRequest(args, serverIp)
	if connectResponse.ChannelStatus == "channel_not_exist" {
		fmt.Printf("Channel %v not exist. Please create channel with create command.\n", strings.Trim(args[0], "\n"))
		return
	} else if connectResponse.ChannelStatus == "channel_exist" {
		caCertPool := x509.NewCertPool()
		var dialer websocket.Dialer = websocket.Dialer{
			TLSClientConfig: &tls.Config{
				RootCAs:            caCertPool,
				InsecureSkipVerify: true,
			},
		}
		utils.CurrentSessionData.CurrentWSConn, _, connErr = dialer.Dial(("wss://" + serverIp + "/ws/" + strings.Trim(args[0], "\n")), nil)
		if connErr != nil {
			log.Fatalf("Websocket connection error: %v", connErr)
		}
		utils.CurrentSessionData.CurrentChan = strings.Trim(args[0], "\n")
		fmt.Printf("Succesful connection to %v\n", utils.CurrentSessionData.CurrentChan)
		go recieveMessege(utils.CurrentSessionData.CurrentWSConn, serverIp)
		for i := 0; i < len(connectResponse.MsgHistory); i++ {
			fmt.Printf("%v:<%v> %v:%v\n", connectResponse.MsgHistory[i].ID, time.Unix(connectResponse.MsgHistory[i].Time, 0).Format("2006-01-0215:04:05"), connectResponse.MsgHistory[i].Login, strings.Trim(connectResponse.MsgHistory[i].Msg, "\n"))
		}
		data := utils.MessegeMarshal("", 2)
		err := utils.CurrentSessionData.CurrentWSConn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Fatalf("Connection to channel messege error: %v", err)
		}
	} else {
		log.Fatalf("Send http connect request error: %v", connectResponse)
	}
}
