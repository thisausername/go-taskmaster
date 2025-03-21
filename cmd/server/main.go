package main

import (
	"fmt"
	"log"

	"github.com/thisausername/go-taskmaster/internal/config"
)

func main() {

	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("配置初始化失败", err)
	}
	fmt.Printf("当前端口", cfg.Server.Port)

	fmt.Println("Hello TakeMaster!")
}
