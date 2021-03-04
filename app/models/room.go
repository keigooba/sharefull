package models

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/keigooba/sharefull/app/test"
)

type room struct {
	forward chan *Message    //メッセージの送信に使う
	join    chan *Client     //参加
	leave   chan *Client     //退室
	clients map[*Client]bool //クライアント保持
	tracer  test.Tracer
	avatar  Avatar
}

func NewRoom() *room {
	return &room{
		forward: make(chan *Message),
		join:    make(chan *Client),
		leave:   make(chan *Client),
		clients: make(map[*Client]bool),
		tracer:  test.Off(), //値を返さない
		avatar:  UseAvatar,
	}
}

func (r *room) Run() {
	for { //無限ループ 強制終了までループし続ける。
		select {
		case client := <-r.join:
			// 参加
			r.clients[client] = true //mapで指定しメモリ消費を抑える
			r.tracer.Trace("新しいクライアントが参加しました")
		case client := <-r.leave:
			// 退室
			delete(r.clients, client) //メモリ削除
			close(client.send)
			r.tracer.Trace("クライアントが退室しました")
		case msg := <-r.forward:
			r.tracer.Trace("メッセージを受信しました:", string(msg.Text))
			// すべてのクライアントにメッセージを送信
			for client := range r.clients {
				select {
				case client.send <- msg:
					// msgを送信
					r.tracer.Trace("-- クライアントに送信されました")
				default:
					// client.sendの受信に空きがない時、送信に失敗
					delete(r.clients, client)
					close(client.send)
					r.tracer.Trace("-- 送信に失敗しました。クライアントをクリーンアップします")
				}
			}
		}
	}
}

// websocket用のシグネチャー（署名）を追加する
const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize,
	WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil) //websocket使用のためアップグレード
	if err != nil {
		log.Fatalln("ServerHTTP:", err)
		return
	}

	cookie, err := req.Cookie("_cookie")
	if err != nil {
		http.Redirect(w, req, "/login", 302)
	} else {
		if cookie.Value == "" {
			http.Redirect(w, req, "/", 302) //空の時リダイレクト
		}
		sess := Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
		u, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}

		client := &Client{
			socket: socket,
			send:   make(chan *Message, messageBufferSize),
			room:   r,
			User:   u,
		}
		r.join <- client
		defer func() { r.leave <- client }()
		go client.write()
		client.read()
	}
}
