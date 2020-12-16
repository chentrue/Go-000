package main

import (
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
	"os"
	"os/signal"
	"service-notification/logic/push"
	"service-notification/models"
	"service-notification/pkg/logging"
	"service-notification/pkg/setting"
	"service-notification/routers"
	"time"
	_ "net/http/pprof"
)

func main()  {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	//set up
	logging.Setup()
	setting.Setup()
	models.Setup()

	//cron
	c := cron.New(cron.WithSeconds(),cron.WithChain(cron.DelayIfStillRunning(cron.DefaultLogger)))
	fmt.Println("crontal starting.....")
	c.AddFunc("*/60 * * * * *", func() {
		//推送完成的回调
		push.Callback()
	})
	c.Start()
	go push.EmailConsume()
	go push.AppConsume()
	go push.SmsConsume()

	r := routers.InitRouter()
	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler: r,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil{
			log.Printf("Listen: %s \n", err)
		}
	}()

	gracefulShutdown(s)
}

func gracefulShutdown(server *http.Server)  {
	//监听中断信号，等待运行中的进程执行完成后再关闭，同时关闭nsq连接
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit
	log.Println("Shutdown Server...")

	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	if err := server.Shutdown(ctx);err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	//关闭消费和生产链接
	push.GracefullShutdown()

	log.Println("Server Exiting")
}
