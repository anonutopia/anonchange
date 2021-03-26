package main

import "gorm.io/gorm"

// KeyValue model is used for storing key/values
type KeyValue struct {
	gorm.Model
	Key      string `sql:"size:255;unique_index"`
	ValueInt uint64 `sql:"type:int"`
	ValueStr string `sql:"type:string"`
}

// Transaction represents node's transaction
type Transaction struct {
	gorm.Model
	TxID      string `sql:"size:255"`
	Processed bool   `sql:"DEFAULT:false"`
}

// Token represents exchange token data
type Token struct {
	gorm.Model
	ID      string `sql:"size:255"`
	Address string `sql:"size:255"`
	Price   uint64 `sql:"type:int"`
}
