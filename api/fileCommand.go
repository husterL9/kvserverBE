package api

import (
	"backend/global"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LsFileRequest struct {
	CurrentPath string `json:"currentPath"`
}
type CdDirRequest struct {
	Path string `json:"path"`
}
type MkdirRequest struct {
	Path string `json:"path"`
}
type CreateRequest struct {
	Path string `json:"filename"`
}

// LsFile
func LsFile(context *gin.Context) {

	var req LsFileRequest
	// 解析请求中的JSON数据
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "请求数据不正确: " + err.Error()})
		return
	}
	fmt.Println("Received command:", req.CurrentPath)
	// 调用KVStoreClient中的LsFile方法
	files, err := global.KVStoreClient.LsFile(req.CurrentPath)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": "error: 无法获取文件列表: " + err.Error(),
		})
		return
	}
	//将files转成字符串
	var filesStr string
	for _, file := range files {
		filesStr += file + "  "
	}

	// 返回成功响应
	context.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": filesStr,
	})
}

// CdDir
func CdDir(context *gin.Context) {
	var req CdDirRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "请求数据不正确: " + err.Error()})
		return
	}
	fmt.Println("Received command:", req.Path)
	err := global.KVStoreClient.CdDir(req.Path)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": "error: 无法切换目录: " + err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "成功切换目录",
	})
}

// Mkdir
func Mkdir(context *gin.Context) {
	var req MkdirRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "请求数据不正确: " + err.Error()})
		return
	}
	fmt.Println("Received command:", req.Path)
	err := global.KVStoreClient.MakeDir(req.Path)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": "error: 无法创建目录: " + err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "成功创建目录",
	})
}

// CreateFile
func CreateFile(context *gin.Context) {
	var req CreateRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "请求数据不正确: " + err.Error()})
		return
	}
	fmt.Println("Received command:", req.Path)
	err := global.KVStoreClient.CreateFile(req.Path)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  "Error",
			"message": "error: 无法创建文件: " + err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "成功创建文件",
	})
}
