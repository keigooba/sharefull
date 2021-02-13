package data

import "time"

type Work struct {
	Id int
	Uuid string
	Date string
	Title string
	Money int
	Station string
    JobId int
	WorkingTimeStart string
	WorkingTimeFinish string
	BusinessDetail string
	Evaluation int
	CreatedAt time.Time
}

// 日付のフォーマット
func (work *Work) CreatedAtDate() string {
	return work.CreatedAt.Format("2006/01/02")
}

// すべての求人情報を取り出す
func Works() (works []Work, err error) {
	rows, err := Db.Query("SELECT date,title FROM works ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Work{}
		if err = rows.Scan(&conv.Title);
		err != nil {
			return
		}
		works = append(works, conv)
	}
	rows.Close()
	return
}
