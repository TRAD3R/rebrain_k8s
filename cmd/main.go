package main

import "github.com/TRAD3R/tlog"

func main() {
	log := tlog.GetLogger(true)
	log.Info("Starting...")
}
