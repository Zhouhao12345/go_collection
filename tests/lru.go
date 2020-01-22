package tests

import (
	"fmt"
	"testing"
	"zhouhao.com/collections/list"
)

func TestLru(t *testing.T) {
	var (
		initPairs map[string]interface{}
		err error
		lru *list.LRU
		value interface{}
	)
	initPairs = make(map[string]interface{})
	initPairs["db_1"] = "db_1"
	initPairs["db_2"] = "db_2"
	initPairs["db_3"] = "db_3"
	initPairs["db_4"] = "db_4"
	initPairs["db_5"] = "db_5"
	initPairs["db_6"] = "db_6"
	initPairs["db_7"] = "db_7"
	if lru, err = list.NewLRU(9, initPairs); err != nil {
		panic("init lru error")
	}
	if err = lru.Insert("db_8", "db_8"); err != nil {
		panic("Insert error")
	}
	if err = lru.Insert("db_9", "db_9"); err != nil {
		panic("Insert error")
	}
	if value, err = lru.Get("db_4"); err != nil {
		panic("Get error")
	} else {
		fmt.Print(value.(string))
	}
	if err = lru.Insert("db_10", "db_10"); err != nil {
		panic("Insert error")
	}
	if err = lru.Insert("db_11", "db_10"); err != nil {
		panic("Insert error")
	}
}
