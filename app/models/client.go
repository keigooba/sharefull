package models

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn //websocketの取得
	send   chan *Message
	room   *room
	user   User
}

func (c *client) read() {
	for { //呼び出し後無限ループ
		// c.socket(websocket)からデータを読み込み、c.room.forwardチャネルに送る
		var msg *Message
		if err := c.socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now().Format("15:04")
			msg.UserName = c.user.Name
			msg.UserID = c.user.ID
			//ここでメッセージを保存する
			if err := msg.CreateMessage(); err != nil {
				log.Fatalln(err)
			}
			c.room.forward <- msg //msgが受信されるまでブロック
		} else {
			break //websocketが異常終了した場合、すぐに閉じる
		}
	}
	c.socket.Close()
}

// read()と並列で処理する
func (c *client) write() {
	for msg := range c.send { //無限ループ バッファ（チャネル）が空の時はチャネルの受信をブロックする
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
