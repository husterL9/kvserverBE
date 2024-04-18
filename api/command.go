package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// CommandRequest 用于解析请求体中的命令
type CommandRequest struct {
	Command string `json:"command"` // 命令类型：set, get, append
	Key     string `json:"key"`     // 键
	Value   string `json:"value"`   // 值，对于 get 命令，此字段可忽略
}

func HandleCommand(context *gin.Context) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	// conn2, err2 := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// defer conn.Close()
	kvstore.NewKVStoreClient(conn)
	c := kvstore.NewKVStoreClient(conn)
	// 设置键值对
	// success, err := c.Set("3", "2", meta)
	// if err != nil {
	// 	log.Fatalf("could not set key-value: %v", err)
	// }
	// fmt.Printf("Set result: %v\n", success)

	// 获取键值对
	c.Set("2", "3", &protobuf.MetaData{})
	var cmd CommandRequest

	// 解析请求体中的 JSON
	if err := context.ShouldBindJSON(&cmd); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "Invalid request data",
		})
		return
	}

	// 处理不同类型的命令
	switch cmd.Command {
	case "set":
		// 处理 set 命令
		SetKey(cmd.Key, cmd.Value)
		context.JSON(http.StatusOK, gin.H{
			"status":  "Success",
			"message": "Key set successfully",
		})
	case "get":
		// 处理 get 命令
		value, err := GetKey(cmd.Key)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{
				"status":  "Error",
				"message": "Key not found",
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"status":  "Success",
			"message": "Key retrieved successfully",
			"value":   value,
		})
	case "append":
		// 处理 append 命令
		AppendKey(cmd.Key, cmd.Value)
		context.JSON(http.StatusOK, gin.H{
			"status":  "Success",
			"message": "Value appended successfully",
		})
	default:
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "Unsupported command",
		})
	}
}
