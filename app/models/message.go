package models

import (
	"log"
	"time"
)

type Message struct {
	Text      string
	UserID    int
	UserName  string
	WorkID    string
	ChatUUID  string
	CreatedAt string
	When      string
	AvatarURL string
	Gravar    string
}

func (m *Message) CreateMessage() error {
	cmd := `insert into messages (
		uuid,
		text,
		user_id,
		user_name,
		work_id,
		chat_uuid,
		created_at) values(?, ?, ?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd, createUUID(), m.Text, m.UserID, m.UserName, m.WorkID, m.ChatUUID, time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetChatUUIDByWorkID(id int) (m Message, err error) {
	m = Message{}
	cmd := `select chat_uuid from messages where work_id = ? group by chat_uuid`
	err = Db.QueryRow(cmd, id).Scan(&m.ChatUUID)
	return m, err
}

func GetChatUUIDByUserID(id int) (s_m []Message, err error) {
	cmd := `select chat_uuid from messages where user_id = ? group by chat_uuid`
	rows, err := Db.Query(cmd, id)
	for rows.Next() {
		var m Message
		err = rows.Scan(&m.ChatUUID)
		if err != nil {
			log.Println(err)
		}
		s_m = append(s_m, m)
	}
	rows.Close()

	return s_m, err
}

func GetMessagesByUUID(uuid string) (messages []Message, err error) {
	cmd := `select text, user_name, work_id, strftime('%m/%d %H:%M', created_at, 'localtime') from messages where chat_uuid = ?`
	rows, err := Db.Query(cmd, uuid)
	if err != nil {
		log.Println(err)
	}
	var m Message

	for rows.Next() {
		err = rows.Scan(&m.Text, &m.UserName, &m.WorkID, &m.CreatedAt)
		if err != nil {
			log.Println(err)
		}
		messages = append(messages, m)
	}
	rows.Close()

	return messages, err
}
