package main

import (
	"log"

	"github.com/anonutopia/gowaves"
	"gorm.io/gorm"
)

var conf *Config

var db *gorm.DB

var wnc *gowaves.WavesNodeClient

var wm *WavesMonitor

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = initConfig()

	db = initDb()

	wnc = initWavesNodeClient()

	initWavesMonitor()
}
