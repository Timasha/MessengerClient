package main

import (
	backendconnection "MessengerClient/internal/backendConnection"
	ui "MessengerClient/internal/cli"
	"MessengerClient/internal/utils"
	"fmt"
)

func main() {
	utils.ServerIp = ""
	fmt.Println(
		`										Messenger by Timasha
Commmands:
For create channel use: /create [channel name]
For connect to existent channel: /connect [channel name]
For disconnect: /disconnect
If you forget commands: /list
For delete messege with given ID: /delete [messege id]
For send messege just type something
		`)
	for {
		ui.Auth(utils.ServerIp)
		ui.CLI(
			backendconnection.CreateChannelFunc,
			backendconnection.ConnectChannelFunc,
			backendconnection.DisconnectFunc,
			func() {
				println(`Commmands:
				For create channel use: /create [channel name]
				For connect to existent channel: /connect [channel name]
				For disconnect: /disconnect
				If you forget commands: /list
				For delete messege with given ID: /delete [messege id]
				For send messege just type something`)
			}, backendconnection.SendMessageFunc,
			backendconnection.DeleteMessegeFunc,
			utils.ServerIp)
		utils.CurrentSessionData.CurrentWSConn.Close()
	}
}
