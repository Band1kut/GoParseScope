package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"sync"
	"time"
)

var tableList = []string{"iPoker"}

type Base struct {
	base *sql.DB
	mux  *sync.Mutex
}

func New() (*Base, error) {
	db, err := sql.Open("sqlite3", "stats.db")
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия базы: %w", err)
	}

	err = createDB(db)
	if err != nil {
		return nil, err
	}

	return &Base{db, &sync.Mutex{}}, nil
}

func createDB(db *sql.DB) error {
	for _, table := range tableList {
		_, err := db.Exec(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
    											id INTEGER PRIMARY KEY AUTOINCREMENT,
												date TEXT,
												name TEXT UNIQUE,
												json TEXT)`, table))
		if err != nil {
			return fmt.Errorf("ошибка создания таблицы %s: %w", table, err)
		}
		_, err = db.Exec(fmt.Sprintf(`CREATE INDEX IF NOT EXISTS idx_%s_data_name ON %s (date, name)`, table, table))
		if err != nil {
			return fmt.Errorf("ошибка создания индекса idx_%s_data_name: %w", table, err)
		}
	}
	return nil
}

func (db *Base) Get(name, network string) (string, error) {
	db.mux.Lock()
	defer db.mux.Unlock()

	date := time.Now().Format("2006-01-02")
	var result string

	err := db.base.QueryRow(
		fmt.Sprintf("SELECT json FROM %s WHERE date = ? AND name = ?", network),
		date, name).Scan(&result)

	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	return result, nil
}

func (db *Base) Set(name, network, json string) error {
	db.mux.Lock()
	defer db.mux.Unlock()

	date := time.Now().Format("2006-01-02")

	_, err := db.base.Exec(
		fmt.Sprintf("INSERT OR REPLACE INTO %s (date, name, json) VALUES (?, ?, ?)", network),
		date, name, json)

	if err != nil {
		return err
	}

	return nil
}

func (db *Base) Close() error {
	db.mux.Lock()
	defer db.mux.Unlock()
	err := db.base.Close()

	if err != nil {
		return err
	}
	return nil
}

//import (
//	"database/sql"
//	_ "github.com/mattn/go-sqlite3"
//)
//
//func InitDB(filepath string) *sql.DB {
//	database, err := sql.Open("sqlite3", filepath)
//
//	if err != nil {
//		panic(err)
//	}
//
//	if database == nil {
//		panic("database nil")
//	}
//
//	return database
//}
//
//func CreateTable(database *sql.DB) {
//	// создание таблицы ips
//	sqlTable := `
//    CREATE TABLE IF NOT EXISTS ips(
//        IP TEXT NOT NULL PRIMARY KEY,
//        Count INTEGER);
//    `
//
//	_, err := database.Exec(sqlTable)
//
//	if err != nil {
//		panic(err)
//	}
//}
//
//func InsertIp(database *sql.DB, ip string) {
//	sqlAddItem := `
//    INSERT OR REPLACE INTO ips(IP, Count)
//    VALUES (?, COALESCE((SELECT Count FROM ips WHERE IP = ?), 0) + 1);
//    `
//
//	stmt, err := database.Prepare(sqlAddItem)
//
//	if err != nil {
//		panic(err)
//	}
//
//	_, err2 := stmt.Exec(ip, ip)
//
//	if err2 != nil {
//		panic(err2)
//	}
//}
