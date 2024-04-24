package api

import (
	"backend/global"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/husterL9/kvserver/api/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CommandRequest 用于解析请求体中的命令
type CommandRequest struct {
	Command string `json:"command"`
}

func HandleCommand(context *gin.Context) {
	// 获取请求体中的命令字符串
	var cmdReq CommandRequest
	if err := context.ShouldBindJSON(&cmdReq); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "Invalid request data",
		})
		return
	}
	fmt.Println("Received command:", cmdReq.Command)
	// 解析命令字符串
	parts := strings.Fields(cmdReq.Command)
	if len(parts) < 1 {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "Invalid command format",
		})
		return
	}

	command := parts[0]
	var key, value string
	if len(parts) > 1 {
		key = parts[1]
	}
	if len(parts) > 2 {
		value = strings.Join(parts[2:], " ")
	}

	// 处理不同类型的命令
	switch command {
	case "set":
		// 处理 set 命令
		global.KVStoreClient.Set(key, value, &protobuf.MetaData{})
		context.JSON(http.StatusOK, gin.H{
			"status":  "Success",
			"message": "Key set successfully",
		})
	case "get":
		// 处理 get 命令
		fmt.Println("get=========", key)
		value, err := global.KVStoreClient.Get(key)
		fmt.Println("err=========", err)
		if err != nil {
			// 解析错误以查看是否是 NotFound
			st, ok := status.FromError(err)
			if ok && st.Code() == codes.NotFound {
				// 键不存在的情况下，依然返回200 OK，但是值为空或者自定义的默认值
				context.JSON(http.StatusOK, gin.H{
					"status":  "Success",
					"message": "Key not found: " + key,
				})
				return
			}
			// 其他类型的错误，返回错误信息
			context.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Error",
				"message": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"status":  "Success",
			"message": value,
		})
	case "append":
		success, err := global.KVStoreClient.Append(key, value, &protobuf.MetaData{})
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Error",
				"message": "Failed to append data: " + err.Error(),
			})
			return
		}
		if success {
			context.JSON(http.StatusOK, gin.H{
				"status":  "Success",
				"message": "Data appended successfully to key: " + key,
			})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Error",
				"message": "Append failed for unknown reasons",
			})
		}
	default:
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": "Unsupported command",
		})
	}
}
