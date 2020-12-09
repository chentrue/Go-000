package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//用于通知主程序全部关闭
	stop := make(chan int)
	// 创建带有cancel的父context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//创建errgroup
	g, _ := errgroup.WithContext(ctx)
	server := http.Server{
		Addr: "127.0.0.1:8090",
		Handler: nil,
	}
	g.Go(func() error {
		return server.ListenAndServe()
	})
	//并发执行监听signal信号，如果接收到信号则关闭全部程序
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		<- c
		cancel()
	}()
	// context cancel后， 关闭http server， 全部关闭后通知主goroutine
	go func() {
		<- ctx.Done()
		//关闭
		go func() {
			fmt.Println(ctx.Err())
			if err := server.Shutdown(context.Background()); err != nil{
				fmt.Printf("server shutdown failed: %s", err)
			}
			close(stop)
			return
		}()
	}()
	if err := g.Wait(); err != nil{
		cancel()        //收到错误后，关闭
		fmt.Println(err)
	}
	<-stop
}
