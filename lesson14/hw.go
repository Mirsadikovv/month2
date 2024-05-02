package main

import (
	"fmt"
	"time"
)

func main() {
	db := &InMemoryDB{
		Data:   make(map[string]interface{}),
		GetCh:  make(chan string),
		SetCh:  make(chan KV),
		DoneCh: make(chan struct{}),
	}

	go db.GetWorker()

	go db.SetWorker()

	db.SetCh <- KV{Key: "key1", Value: "привет"}

	db.GetCh <- "key1"
	value := <-db.GetCh
	fmt.Println("Value:", value)
	time.Sleep(time.Second)
}

type InMemoryDB struct {
	Data   map[string]interface{}
	GetCh  chan string
	SetCh  chan KV
	DoneCh chan struct{}
}

type KV struct {
	Key   string
	Value interface{}
}

func (db *InMemoryDB) Get(key string) (interface{}, bool) {
	db.GetCh <- key

	value := <-db.GetCh

	return value, true
}

func (db *InMemoryDB) Set(kv KV) (bool, error) {
	db.SetCh <- kv

	select {
	case <-db.DoneCh:
		return true, nil
	}
}

func (db *InMemoryDB) GetWorker() {
	for {
		select {
		case key := <-db.GetCh:
			value, ok := db.Data[key]
			if !ok {
				db.GetCh <- "not found"
			} else {
				db.GetCh <- fmt.Sprintf("%v", value)
			}
		}
	}
}

func (db *InMemoryDB) SetWorker() {
	for {
		select {
		case kv := <-db.SetCh:
			db.Data[kv.Key] = kv.Value
			db.DoneCh <- struct{}{}
		}
	}
}
