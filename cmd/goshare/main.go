package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func withUI() {
	var c config
	err := loadConfig("configs/config.json", &c)
	if err != nil {
		panic(err)
	}
	go func() {
		time.Sleep(time.Second) // 等listen准备好,打开默认浏览器
		cmd := exec.Command("explorer", "http://localhost:9090")
		err := cmd.Start()
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	// 开启gin api router
	go func() {
		time.Sleep(time.Millisecond * 100)
		opt := &Options{}
		gw := NewGateway(opt)
		gw.Run(9090)
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	go func() {
		// for debug http://localhost:6160/debug/pprof
		log.Println(http.ListenAndServe(":6160", nil))
	}()
	var op = flag.String("op", "", "operation")
	flag.Parse()
	switch *op {
	case "bonus":
		bonus()
	default:
		withUI()
	}

}
