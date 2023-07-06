package main

import (
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {

	cfg, err := GetConfig()
	if err != nil {
		log("get env config error")
		panic(err)
	}

	chat := NewQchat()

	h := NewHandler()

	http.Handle("/", http.FileServer(http.Dir(cfg.Web)))
	http.Handle("/ws", websocket.Handler(chat.handleWs))
	http.HandleFunc("/auth/login", h.Login)
	log("listening: ", cfg.Port)
	err = http.ListenAndServe(cfg.Port, nil)
	if err != nil {
		panic(err)
	}

}
