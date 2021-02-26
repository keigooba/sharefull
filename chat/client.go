package chat

import "github.com/gorilla/websocket"

type client struct {
	socket *websocket.Conn //websocketの取得
	send   chan []byte
	room   *room
}



func (c *client) read() {
	for { //呼び出し後無限ループ
		// c.socket(websocket)からデータを読み込み、c.room.forwardチャネルに送る
		if _, msg, err := c.socket.ReadMessage(); err == nil {
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
		if err := c.socket.WriteMessage(websocket.TextMessage, msg);
		err != nil {
			break
		}
	}
	c.socket.Close()
}
