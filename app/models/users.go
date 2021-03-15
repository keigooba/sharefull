package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	AvatarURL string
	AvatarID  string
	CreatedAt time.Time
	ApplyID   int
}

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		avatar_url,
		avatar_id,
		created_at) values($1, $2, $3, $4, $5, $6, $7)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		u.AvatarURL,
		u.AvatarID,
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, avatar_id, created_at from users where id = $1`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.AvatarID,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) UpdateUser() error {
	cmd := `update users set name = $1, email = $2, password = $3 where id = $4`
	_, err = Db.Exec(cmd, u.Name, u.Email, Encrypt(u.PassWord), u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() error {
	cmd := `delete from users where id = $1`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions (uuid, email, user_id, created_at) values($1, $2, $3, $4)`
	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	cmd2 := `select id, uuid, email, user_id, created_at from sessions where email = $1 and user_id = $2`
	err = Db.QueryRow(cmd2, u.Email, u.ID).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)

	return session, err
}

func (login_user *User) GetUserLogin() (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at from users where email = $1 and password = $2`
	err = Db.QueryRow(cmd, login_user.Email, Encrypt(login_user.PassWord)).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.PassWord, &user.CreatedAt)

	return user, err
}

func (sess *Session) CheckSession() (valid bool, err error) {
	cmd := `select id, uuid, email, user_id, created_at from sessions where uuid = $1`
	err = Db.QueryRow(cmd, sess.UUID).Scan(&sess.ID, &sess.UUID, &sess.Email, &sess.UserID, &sess.CreatedAt)

	if err != nil {
		valid = false
		return valid, err
	}

	if sess.ID != 0 {
		valid = true
	}

	return valid, err
}

func (sess *Session) DeleteSessionByUUID() error {
	cmd := `delete from sessions where uuid = $1`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (sess *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, avatar_url, avatar_id, created_at from users where id = $1`
	err = Db.QueryRow(cmd, sess.UserID).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.PassWord, &user.AvatarURL,
		&user.AvatarID, &user.CreatedAt)
	if err != nil {
		log.Println(err)
	}
	return user, err
}

// Google認証
func (u *User) AuthGetUser() (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at from users where email = $1`
	err = Db.QueryRow(cmd, u.Email).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.PassWord, &user.CreatedAt)

	//errorも返すことがある為、errorハンドリングなし
	return user, err
}
