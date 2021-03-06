package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/google/uuid"
	// _ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

type Data struct {
	Works        []Work
	ApplyWorks   []Work
	WorkID       interface{}
	User         User
	ApplyUsers   []User
	ApplysID     []int
	ChatUUID     string
	Messages     []Message
	SendMessages []Message
	NowDate      string
	Host         string
}

const (
	tableNameUser      = "users"
	tableNameWork      = "works"
	tableNameSession   = "sessions"
	tableNameApplyUser = "apply_users"
	tableNameMessage   = "messages"
)

func init() {
	// ローカル sqlite3 (処理の修正も必要有り)
	// Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)

	// ローカル postgres
	// Db, err = sql.Open("postgres", "host=ec2-54-242-43-231.compute-1.amazonaws.com port=5432 user=xbvbdmyqrzdjno password=password dbname=d1kg7305bvf6n6 sslmode=require")
	// 本番 postgres
	Db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatalf("db接続でエラー %v", err)
	}

	// 	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 		Uuid STRING NOT NULL UNIQUE,
	// 		name STRING,
	// 		email STRING,
	// 		password STRING,
	// 		avatar_url STRING,
	// 		avatar_id STRING,
	// 		created_at DATETIME)`, tableNameUser)

	// 	Db.Exec(cmdU)

	// 	cmdW := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 		date STRING,
	// 		title STRING,
	// 		money STRING,
	// 		job_id STRING,
	// 		evaluation STRING,
	// 		user_id INTERGER,
	// 		created_at DATETIME)`, tableNameWork)

	// 	Db.Exec(cmdW)

	// 	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 		uuid STRING NOT NULL UNIQUE,
	// 		email STRING,
	// 		user_id INTEGER,
	// 		created_at DATETIME)`, tableNameSession)

	// 	Db.Exec(cmdS)

	// 	cmdA := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 		uuid STRING NOT NULL UNIQUE,
	// 		work_id INTEGER,
	// 		user_id INTEGER,
	// 		created_at DATETIME)`, tableNameApplyUser)

	// 	Db.Exec(cmdA)

	// 	cmdM := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 		uuid STRING NOT NULL UNIQUE,
	// 		text STRING,
	// 		user_id INTEGER,
	// 		user_name STRING,
	// 		work_id INTEGER,
	// 		chat_uuid STRING NOT NULL,
	// 		created_at DATETIME)`, tableNameMessage)

	// 	Db.Exec(cmdM)
}

// func Migration() {
// 	u := User{}
// 	u.Name = "test"
// 	u.Email = "test@email.com"
// 	u.PassWord = "testtest"
// 	u.CreateUser()

// 	work := &Work{
// 		Date:       "02/20",
// 		Title:      "テスト",
// 		Money:      "3000",
// 		JobID:      "1",
// 		Evaluation: "1",
// 	}
// 	u.CreateWork(work)
// }

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}

const rsLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = rsLetters[rand.Intn(len(rsLetters))]
	}
	return string(b)
}
