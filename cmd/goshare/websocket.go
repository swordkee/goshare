package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	pb "github.com/mineralres/protos/src/go/goshare"
	"google.golang.org/grpc"
)

type wsFront struct {
	grpcConn     *grpc.ClientConn
	grpcEndPoint string
	connLock     sync.RWMutex
	port         int
}

func makeWsFront(c config) *wsFront {
	var front wsFront
	front.grpcEndPoint = fmt.Sprintf("localhost:%d", c.Port)
	front.port = c.Port
	return &front
}

// 因为grpc-gateway对stream 的实现不是特别成熟,所以此处先用websocket完成stream方式的推送
func (front *wsFront) run() {
	r := gin.New()
	r.GET("/ws/uploadTick", front.uploadTick)
	s := &http.Server{
		Addr:    ":" + strconv.Itoa(front.port),
		Handler: r,
	}
	s.SetKeepAlivesEnabled(false)
	log.Printf("websocket listen on [%d] ", front.port)
	s.ListenAndServe()
}

// 推送tick到自建数据库
func (front *wsFront) uploadTick(c *gin.Context) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	for {
		conn.SetReadDeadline(time.Now().Add(15 * time.Second))
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("wsUploadTick exit", err)
			break
		} else {
			if messageType == websocket.BinaryMessage {
				var mds pb.MarketDataSnapshot
				if err := proto.Unmarshal(p, &mds); err == nil {
					if mds.Symbol == "" {
						continue
					}
				}
			}
		}
	}
}
