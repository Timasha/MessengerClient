package utils

import "github.com/gorilla/websocket"

type SessionData struct {
	CurrentLogin  string
	CurrentChan   string
	JWT           string
	CurrentWSConn *websocket.Conn
}

var ServerIp string
var CurrentSessionData SessionData
