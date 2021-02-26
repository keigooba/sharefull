package chat

type room struct {
	forward chan []byte //メッセージの送信に使う
}
