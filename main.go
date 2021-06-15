package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func init() {
	//f, _ := os.OpenFile("./logs/flowapp.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetLevel(log.TraceLevel)

	//log.SetOutput(f)
}

var quit = make(chan struct{})

func main() {
	fmt.Println("Ola mundo Digital CI-CD")
	log.Info("Ola mundo Digital CI-CD")
	done := make(chan bool, 1)

	handleSigterm(func() {
		log.Info("App encerrado")
	})

	<-done
}

func handleSigterm(handleExit func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		handleExit()
		os.Exit(1)
	}()
}
