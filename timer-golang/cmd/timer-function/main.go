package main

import (
	"os"

	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

var logger *slog.Logger

func getLogger() *slog.Logger {
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}
	return logger
}

func postTick(c *gin.Context) {
	getLogger().Info("This HTTP triggered function executed successfully.", slog.Int("status_code", 200), slog.String("status", "OK"))
	// do work here
}

func getPort() string {
	port := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = ":" + val
	}
	return port
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/TimerFunction", postTick)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})
	r.Run(getPort())
}
