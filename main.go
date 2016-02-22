package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/TemplateTalk/first"
	"github.com/TemplateTalk/master"
	"github.com/TemplateTalk/paragraph"
	"github.com/TemplateTalk/second"
)

func main() {
	go listen(":8100")
	waitForSignal()
	os.Exit(0)
}

func listen(hostport string) {
	log.Println("Server Started at", hostport)
	master.AddContent("/first/", first.GetName(), nil, first.GetFile())
	master.AddContent("/second/", second.GetName(), nil, second.GetFiles()...)
	master.AddContent("/paragraph/", paragraph.GetName(), paragraph.GetContent(), paragraph.GetFile())
	http.ListenAndServe(hostport, master.Mux())
}

func waitForSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	_ = <-c
}
