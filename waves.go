package main

import (
	"log"
	"time"

	"github.com/anonutopia/gowaves"
)

func initWavesNodeClient() *gowaves.WavesNodeClient {
	wnc := &gowaves.WavesNodeClient{
		Host: WavesNodeURL,
		Port: 80,
	}

	return wnc
}

type WavesMonitor struct {
	StartedTime int64
}

func (wm *WavesMonitor) start() {
	wm.StartedTime = time.Now().Unix() * 1000
	mylog()
	for {
		// todo - make sure that everything is ok with 100 here
		pages, err := wnc.TransactionsAddressLimit(conf.Address, 100)
		if err != nil {
			log.Println(err)
		}

		if len(pages) > 0 {
			for _, t := range pages[0] {
				wm.checkTransaction(&t)
			}
		}

		time.Sleep(time.Second)
	}
}

func (wm *WavesMonitor) checkTransaction(talr *gowaves.TransactionsAddressLimitResponse) {
	tr := &Transaction{TxID: talr.ID}
	db.FirstOrCreate(tr, tr)
	if !tr.Processed {
		wm.processTransaction(tr, talr)
	}
}

func (wm *WavesMonitor) processTransaction(tr *Transaction, talr *gowaves.TransactionsAddressLimitResponse) {
	tr.Processed = true
	db.Save(tr)
}

func initWavesMonitor() {
	wm = &WavesMonitor{}
	wm.start()
}
