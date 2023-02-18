package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger(), gindump.Dump())

	server.Run(":8016")

}
