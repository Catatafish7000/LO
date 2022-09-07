package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const OrdersTable = "\"Orders\""

type ConfigDB struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

type Exchange struct {
	C  *Cache
	DB *sqlx.DB
}

func NewPsqlDB(cfg ConfigDB) (*sqlx.DB, error) {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName)
	db, err := sqlx.Open("postgres", psqlConn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (ex *Exchange) AddOrderToDB(key string) error {
	tx, err := ex.DB.Begin()
	if err != nil {
		return err
	}
	data, err := json.Marshal(ex.C.cache[key])
	ans := string(data)
	fmt.Println(ans)
	psqlQ := fmt.Sprintf("INSERT INTO Orders (key,order_data) VALUES ('%s','%v');", ex.C.cache[key].UID, ans)
	tx.QueryRow(psqlQ)
	tx.Commit()
	return nil
}
