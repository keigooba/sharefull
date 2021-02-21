package models

import (
	"log"
	"time"
)

type ApplyUser struct {
	ID        int
	UUID      string
	UserID    int
	WorkID    int
	CreatedAt time.Time
}

func (u *User) CreateApplyUser(id int) error {
	cmd := `insert into apply_users (
		uuid,
		work_id,
		user_id,
		created_at) values (?, ?, ?, ?)`

	_, err = Db.Exec(cmd, createUUID(), id, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetApplyUsersByUserID(id int) (works_id []int, err error) {
	cmd := `select work_id from apply_users where user_id = ?`
	rows, err := Db.Query(cmd, id)
	if err != nil {
		log.Fatalln(err)
	}
	var a_user ApplyUser

	for rows.Next() {
		err = rows.Scan(&a_user.WorkID)

		if err != nil {
			log.Fatalln(err)
		}
		work_id := a_user.WorkID
		works_id = append(works_id, work_id)
	}
	rows.Close()

	return works_id, err
}

func GetApplyUsersByWorkID(id int) (applys_id []int, users_id []int, err error) {
	cmd := `select id, user_id from apply_users where work_id = ?`
	rows, err := Db.Query(cmd, id)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var a_user ApplyUser
		err = rows.Scan(&a_user.ID, &a_user.UserID)

		if err != nil {
			log.Fatalln(err)
		}
		apply_id := a_user.ID
		applys_id = append(applys_id, apply_id)
		user_id := a_user.UserID
		users_id = append(users_id, user_id)
	}
	return applys_id, users_id, err
}

func ApplyUserDelete(id int) error {
	cmd := `delete from apply_users where id = ?`
	_, err := Db.Exec(cmd, id)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
