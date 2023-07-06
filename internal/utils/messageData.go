package utils

type Message struct {
	ID      uint   `json:"id"`
	MsgType int    `json:"msgtype"`
	Login   string `json:"login"`
	Time    int64  `json:"time"`
	Msg     string `json:"msg"`
	JWT     string `json:"jwt"`
	Channel string `json:"channel"`
}
type HistoryMessage struct {
	ID      uint64 `json:"id"`
	Login   string `json:"login"`
	Time    int64  `json:"time"`
	Msg     string `json:"msg"`
	Channel string `json:"channel"`
}
