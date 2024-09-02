package main

import (
	"time"

	"github.com/TRAD3R/tlog"
)

func main() {
	log := tlog.GetLogger(true)
	log.Info("Starting...")

	ch := make(chan int)

	go print(log)
	<-ch
}

func print(log *tlog.Logger) {
	for {
		time.Sleep(1 * time.Second)
		log.Info("printing")
	}
}
