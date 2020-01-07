package main

import (
	"echoserver/config"
	"echoserver/router"
	"fmt"
)

func main() {
	config := config.New()
	e := router.New()
	fmt.Println("Data =>", config.ConfigData)
	port := fmt.Sprintf("%s%d", ":", config.ConfigData.Server.Port)
	e.Logger.Fatal(e.Start(port))
}
