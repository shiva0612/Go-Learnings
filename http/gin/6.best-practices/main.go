package main

import (
	"context"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
}

func best_practice(c *gin.Context) {
	r := gin.Default()
	r.Use(gin.Logger())

	logfile, _ := os.Open("someifle")
	gin.DefaultWriter = io.MultiWriter(logfile, os.Stdout)

	// when using concurrent handler make sure to copy the context
	ctx := c.Copy()
	//call the service layer
	serviceLayer(ctx)

	//always use server.shutdown for graceful shutdown in http servers
}

func serviceLayer(ctx context.Context) {
	//some businesss logic implementation
}
