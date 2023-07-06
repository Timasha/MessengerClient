package utils

import (
	"encoding/json"
	"log"
	"time"
)

type AuthData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func LoginMessage(login string, password string) (data []byte, err error) {
	authData := AuthData{
		Login:    login,
		Password: password,
	}
	data, err = json.Marshal(authData)
	if err != nil {
		return data, err
	}
	return
}
func MessegeUnmarshal(data []byte) (msg Message, err error) {
	err = json.Unmarshal(data, &msg)
	return
}
func MessegeMarshal(msg string, msgType int) (data []byte) {
	var marshErr error
	data, marshErr = json.Marshal(Message{
		MsgType: msgType,
		Login:   CurrentSessionData.CurrentLogin,
		Time:    time.Now().Unix(),
		Msg:     msg,
		JWT:     CurrentSessionData.JWT,
		Channel: CurrentSessionData.CurrentChan,
	})
	if marshErr != nil {
		log.Fatalf("Marshal messege json error: %v", marshErr)
	}
	return
}
