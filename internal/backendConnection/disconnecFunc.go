package backendconnection

import (
	"MessengerClient/internal/utils"
	"fmt"
	"log"
)

var DisconnectFunc func() = func() {
	var err error
	if utils.CurrentSessionData.CurrentWSConn != nil {
		err = utils.CurrentSessionData.CurrentWSConn.Close()
	} else {
		fmt.Println("You are not connected.")
		return
	}
	utils.CurrentSessionData.CurrentWSConn = nil
	utils.CurrentSessionData.CurrentChan = ""
	if err == nil {
		fmt.Println("Disconnect succesful")
	} else {
		log.Fatalf("Cant disconnect error: %v", err)
	}
}
