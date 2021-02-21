package models

import (
	"log"
	"time"
)

type Work struct {
	ID             int
	Date           string
	Title          string
	Money          string
	JobID          string
	Evaluation     string
	UserID         int
	CreatedAt      time.Time
	User           User
	JobList        []string
	EvaluationList []string
}

func (w *Work) WorkList() error {
	w.JobList = []string{"エキストラ", "データ入力", "モニターバイト", "仕分けバイト", "工場バイト", "カフェ", "コールセンター", "イベントスタッフ", "試験監督"}
	w.EvaluationList = []string{"設定しない", "1以上", "2以上", "3以上"}
	return err
}

func (u *User) CreateWork(work *Work) (err error) {
	cmd := `insert into works (
		date,
		title,
		money,
		job_id,
		evaluation,
		user_id,
		created_at) values (?, ?, ?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd, work.Date, work.Title, work.Money, work.JobID, work.Evaluation, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetWork(id int) (work Work, err error) {
	cmd := `select id, date, title, money, job_id, evaluation, user_id, created_at from works
	where id = ?`
	work = Work{}

	err = Db.QueryRow(cmd, id).Scan(
		&work.ID,
		&work.Date,
		&work.Title,
		&work.Money,
		&work.JobID,
		&work.Evaluation,
		&work.UserID,
		&work.CreatedAt)

	return work, err
}

func GetWorks() (works []Work, err error) {
	cmd := `select id, strftime('%Y/%m/%d', date), title, money, job_id, evaluation, user_id, created_at from works where date >= date(CURRENT_TIMESTAMP) order by date asc`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var work Work
		err = rows.Scan(
			&work.ID,
			&work.Date,
			&work.Title,
			&work.Money,
			&work.JobID,
			&work.Evaluation,
			&work.UserID,
			&work.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}
		works = append(works, work)
	}
	rows.Close()

	return works, err
}

func (u *User) GetWorksByUser() (works []Work, err error) {
	cmd := `select id, date, title, money, job_id, evaluation, user_id, created_at from works where user_id = ?`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var work Work
		err = rows.Scan(
			&work.ID,
			&work.Date,
			&work.Title,
			&work.Money,
			&work.JobID,
			&work.Evaluation,
			&work.UserID,
			&work.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}
		works = append(works, work)
	}
	rows.Close()

	return works, err
}

func (w *Work) UpdateWork() (err error) {
	cmd := `update works set date = ?, title = ?, money = ?, job_id = ?, evaluation = ? where id = ?`
	_, err = Db.Exec(cmd, w.Date, w.Title, w.Money, w.JobID, w.Evaluation, w.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (w *Work) DeleteWork() (err error) {
	cmd := `delete from works where id = ?`
	_, err = Db.Exec(cmd, w.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
