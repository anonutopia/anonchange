package main

import (
	"log"

	"gorm.io/gorm"
)

var conf *Config

var db *gorm.DB

var wm *WavesMonitor

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	db = initDb()

	initWavesMonitor()
}
