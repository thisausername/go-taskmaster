package main

import (
	"log"

	"github.com/thisausername/go-taskmaster/internal/config"
	"github.com/thisausername/go-taskmaster/internal/server"
)

func main() {

	_, err := config.Init()
	if err != nil {
		log.Fatalf("配置初始化失败", err)
	}

	server.StartServer()
}
