package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/thisausername/go-taskmaster/internal/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartServer() {
	cfg := config.GetConfig()
	router := InitRouter()

	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: router,
	}
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		<-sigChan
		log.Println("收到终止信号，开始优雅关闭服务器...")

		//创建一个带有超时的上下文
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		//关闭服务器
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("强制关闭服务器:", err)
		}
	}()

	log.Printf("服务器启动,监听端口%s,环境:%s", cfg.Server.Port, cfg.Server.Env)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("服务器启动失败:%v", err)
	}
}
