package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func init() {
	// ローカル
	// Db, err = sql.Open("postgres", "host=ec2-54-242-43-231.compute-1.amazonaws.com port=5432 user=xbvbdmyqrzdjno password=0627faa9325b1a74827905a36003c023334a4ea5072ebef0114eae4ca979bfa1 dbname=d1kg7305bvf6n6 sslmode=require")
	// 本番
	Db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("db接続でエラー %v", err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(id SERIAL PRIMARY KEY,
		Uuid CHAR(64) NOT NULL UNIQUE,
		name CHAR(255),
		email CHAR(255),
		password CHAR(255),
		avatar_url CHAR(255),
		avatar_id CHAR(255),
		created_at TIMESTAMP)`, tableNameUser)
	_, err = Db.Exec(cmdU)
	if err != nil {
		log.Fatalln(err)
	}

	cmdW := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(id SERIAL PRIMARY KEY,
		date CHAR(255),
		title CHAR(255),
		money CHAR(255),
		job_id CHAR(255),
		evaluation CHAR(255),
		user_id INTEGER,
		created_at TIMESTAMP)`, tableNameWork)
	_, err = Db.Exec(cmdW)
	if err != nil {
		log.Fatalln(err)
	}

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(id SERIAL PRIMARY KEY,
		uuid CHAR(64) NOT NULL UNIQUE,
		email CHAR(255),
		user_id INTEGER,
		created_at TIMESTAMP)`, tableNameSession)
	_, err = Db.Exec(cmdS)
	if err != nil {
		log.Fatalln(err)
	}

	cmdA := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(id SERIAL PRIMARY KEY,
		uuid CHAR(255) NOT NULL UNIQUE,
		work_id INTEGER,
		user_id INTEGER,
		created_at TIMESTAMP)`, tableNameApplyUser)
	_, err = Db.Exec(cmdA)
	if err != nil {
		log.Fatalln(err)
	}

	cmdM := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(id SERIAL PRIMARY KEY,
		uuid CHAR(64) NOT NULL UNIQUE,
		text CHAR(255),
		user_id INTEGER,
		user_name CHAR(255),
		work_id INTEGER,
		chat_uuid CHAR(64) NOT NULL,
		created_at TIMESTAMP)`, tableNameMessage)
	_, err = Db.Exec(cmdM)
	if err != nil {
		log.Fatalln(err)
	}
}
