package main

import (
	"log"
	"time"
)

type WavesMonitor struct {
}

func (wm *WavesMonitor) start() {
	for {
		log.Println("tick")
		time.Sleep(time.Second)
	}
}

func initWavesMonitor() {
	wm = &WavesMonitor{}
	wm.start()
}
