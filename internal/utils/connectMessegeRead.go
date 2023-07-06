package utils

import (
	"encoding/json"
	"log"
)

type ConnectJson struct {
	ChannelStatus string           `json:"channelStatus"`
	MsgHistory    []HistoryMessage `json:"msgHistory"`
}

func ReadConnectJSON(data []byte) (result ConnectJson) {
	jsonUnmarshErr := json.Unmarshal(data, &result)
	if jsonUnmarshErr != nil {
		log.Fatalf("Json unmarshal error: %v", jsonUnmarshErr)
	}
	return
}
