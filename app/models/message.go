package models

import (
	"log"
	"time"
)

type Message struct {
	Text      string
	UserID    int
	WorkID    string
	ChatUUID  string
	CreatedAt time.Time
	Name      string
	When      string
}

func (m *Message) CreateMessage() error {
	cmd := `insert into messages (
		uuid,
		text,
		user_id,
		work_id,
		chat_uuid,
		created_at) values(?, ?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd, createUUID(), m.Text, m.UserID, m.WorkID, m.ChatUUID, time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetChatUUIDByWorkID(id int) (m Message, err error) {
	m = Message{}
	cmd := `select chat_uuid from messages where work_id = ? group by chat_uuid`
	err = Db.QueryRow(cmd, id).Scan(&m.ChatUUID)
	if err != nil {
		log.Println(err)
	}
	return m, err
}

func GetMessageByWorkID(id int) (messages []Message, err error) {
	cmd := `select text, user_id, chat_uuid from messages where work_id = ?`
	rows, err := Db.Query(cmd, id)
	if err != nil {
		log.Println(err)
	}
	var m Message

	for rows.Next() {
		err = rows.Scan(&m.Text, &m.UserID, &m.ChatUUID)
		if err != nil {
			log.Println(err)
		}
		messages = append(messages, m)
	}
	rows.Close()

	return messages, err
}
