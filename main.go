package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

var log = fmt.Println

const (
	port = ":8887"
	web  = "./web"
)

type Qchat struct {
	Conns  []Conn
	Counts int32
}

type Conn struct {
	Id   string
	Conn *websocket.Conn
	User User
}

type User struct {
	Name string `json:"_init_user_name"`
}

type Message struct {
	Content string `json:"content"`
	Group   string `json:"group"`
}

func NewQchat() *Qchat {
	return &Qchat{
		Counts: 0,
	}
}

func (c *Qchat) handleWs(ws *websocket.Conn) {
	c.Counts += 1
	conn := Conn{
		Id:   uuid.New().String(),
		Conn: ws,
	}
	c.Conns = append(c.Conns, conn)

	log("users: ", c.Counts)

	c.readListener(conn)

}

func (c *Qchat) readListener(ws Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := ws.Conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log("client disconnect")
				break
			}
			log("ws read failed")
			continue
		}
		msg := buf[:n]

		if IsFirstMsg(msg) { // if init message with user name
			var user User
			err = json.Unmarshal(buf[:n], &user)
			if err != nil {
				log("marshal user failed")
				continue
			}
			ws.User.Name = user.Name

			log("id: ", ws.Id)
			log("name: ", ws.User.Name)
			c.BroadCast(ws.User.Name, " entered")
		} else { // else deal with other message with data and channel
			var m Message
			err = json.Unmarshal(msg, &m)
			if err != nil {
				log("marshal message failed")
				continue
			}
			log("group: ", m.Group)
			log("content: ", m.Content)
			c.BroadCast(ws.User.Name, m.Content)
		}
	}
}

func (c *Qchat) BroadCast(name string, content string) {
	for _, conn := range c.Conns {
		log(conn.User.Name)
		log("in broadcast loop")
		conn.Conn.Write([]byte(name + ": " + content))
	}
}

func main() {

	chat := NewQchat()

	http.Handle("/", http.FileServer(http.Dir(web)))
	http.Handle("/ws", websocket.Handler(chat.handleWs))
	log("listening: ", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}

}
