package main

import "log"

var wm *WavesMonitor

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	initWavesMonitor()
}
