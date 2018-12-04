package main

import (
	"github.com/buchenglei/service-template/config"
	router "github.com/buchenglei/service-template/service/http/router"
)

func main() {
	config := config.Get()
	// demo
	config.WebAddress = ":8080"

	httpService := router.GetRouters()
	err := httpService.Run(config.WebAddress)
	if err != nil {
		panic(err)
	}
}
