package main

import (
	"entry/src/api"
	"entry/src/api/rest"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var service api.IAPI = rest.New()
	err := service.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	wait()

}
func wait() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

}
