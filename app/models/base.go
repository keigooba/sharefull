package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	"github.com/keigooba/sharefull/config"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

type Data struct {
	Works      []Work
	User       User
	ApplyUsers []ApplyUser
	NowDate    string
}

const (
	tableNameUser      = "users"
	tableNameWork      = "works"
	tableNameSession   = "sessions"
	tableNameApplyUser = "apply_users"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		Uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	Db.Exec(cmdU)

	cmdW := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date STRING,
		title STRING,
		money STRING,
		job_id STRING,
		evaluation STRING,
		user_id INTERGER,
		created_at DATETIME)`, tableNameWork)

	Db.Exec(cmdW)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		user_id INTEGER,
		created_at DATETIME)`, tableNameSession)

	Db.Exec(cmdS)

	cmdA := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		work_id INTEGER,
		user_id INTEGER,
		created_at DATETIME)`, tableNameApplyUser)

	Db.Exec(cmdA)
}

func Migration() {
	u := User{}
	u.Name = "test"
	u.Email = "test@email.com"
	u.PassWord = "testtest"
	u.CreateUser()

	work := &Work{
		Date:       "02/20",
		Title:      "テスト",
		Money:      "3000",
		JobID:      "1",
		Evaluation: "1",
	}
	u.CreateWork(work)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
