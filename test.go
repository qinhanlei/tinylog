package main

import log "github.com/qinhanlei/tinylog"

func init() { log.Init(".") }

func main() {
	log.Debug("Hello Debug")
	log.Info("Hello Info")
	log.Warn("Hello Warn")
	log.Error("Hello Error")
	log.Fatal("Hello Fatal")
}
