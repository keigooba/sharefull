package models

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	socket *websocket.Conn //websocketの取得
	send   chan *Message
	room   *room
	User   User
}

func (c *Client) read() {
	for { //呼び出し後無限ループ
		// c.socket(websocket)からデータを読み込み、c.room.forwardチャネルに送る
		var msg *Message
		if err := c.socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now().Format("15:04")
			msg.UserName = c.User.Name
			msg.UserID = c.User.ID
			if msg.Gravar == "Gravatar送信" { //Gravatarから画像受け取り
				msg.AvatarURL, _ = c.room.avatar.GravatarAvatarURL(c)
			} else {
				msg.AvatarURL, _ = c.room.avatar.AvatarURL(c)
			}
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
func (c *Client) write() {
	for msg := range c.send { //無限ループ バッファ（チャネル）が空の時はチャネルの受信をブロックする
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
