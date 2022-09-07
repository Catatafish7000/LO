package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"sync"
)

type Cache struct {
	cache map[string]Order
	mx    *sync.Mutex
}

func NewCache() *Cache {
	db := new(Cache)
	db.cache = make(map[string]Order)
	db.mx = new(sync.Mutex)
	return db
}

func (ex *Exchange) GetCache() error {
	psqlQ := fmt.Sprintf("SELECT order_data FROM %s ;", "Orders")
	rows, err := ex.DB.Queryx(psqlQ)
	if err != nil {
		return err
	}
	for rows.Next() {
		order := new(Order)
		var info []byte
		err = rows.Scan(&info)
		if err != nil {
			return err
		}
		err = json.Unmarshal(info, order)
		if err != nil {
			return err
		}
		ex.C.Add(order.UID, *order)
	}
	return nil
}

func (c *Cache) Add(key string, value Order) {
	c.mx.Lock()
	c.cache[key] = value
	c.mx.Unlock()
}

func (ex *Exchange) SaveInformation(msg *stan.Msg) {
	ord := new(Order)
	err := json.Unmarshal(msg.MsgProto.Data, ord)
	if err != nil {
		fmt.Printf("invalid format data %s \n", string(msg.MsgProto.Data))
		return
	}
	ex.C.mx.Lock()
	ex.C.cache[ord.UID] = *ord
	ex.C.mx.Unlock()
	if err = ex.AddOrderToDB(ord.UID); err != nil {
		fmt.Printf("invalid format data %s \n", string(msg.MsgProto.Data))
		return
	}
}

func (c *Cache) GetInformation(uid string) ([]byte, error) {
	ord := c.cache[uid]
	ans, err := json.Marshal(ord)
	if err != nil {
		return nil, err
	}
	if ord.UID == "" {
		return []byte("no such order"), nil
	}
	return ans, nil
}

type Saver interface {
	SaveInformation(msg *stan.Msg)
}

type Getter interface {
	GetInformation(uid string) ([]byte, error)
}
