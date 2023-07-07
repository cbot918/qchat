package main

import (
	"net/http"

	"github.com/cbot918/qchat/logger"
	"golang.org/x/net/websocket"
)

func main() {

	cfg, err := GetConfig()
	if err != nil {
		log("get env config error")
		panic(err)
	}

	l := logger.NewLogger()

	s := NewStorage(cfg)
	s.InitPsql()

	chat := NewQchat(s, l)

	http.Handle("/", http.FileServer(http.Dir(cfg.Web)))
	http.Handle("/ws", websocket.Handler(chat.handleWs))
	http.HandleFunc("/auth/login", chat.H.LoginHandler)
	http.HandleFunc("/friend/list", chat.H.ListFriend)
	log("listening: ", cfg.Port)
	err = http.ListenAndServe(cfg.Port, nil)
	if err != nil {
		panic(err)
	}

}
