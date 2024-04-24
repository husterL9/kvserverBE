package initialize

import (
	"fmt"
	"log"

	"github.com/husterL9/kvserver/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitKVClient() *client.KVStoreClient {
	fmt.Println("-===========")
	// 连接到gRPC服务
	//todo ,尝试重新连接
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// defer conn.Close()

	c := client.NewKVStoreClient(conn)
	return c
}
