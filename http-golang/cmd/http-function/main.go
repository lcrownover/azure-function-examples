package main

import (
	"os"

	"log/slog"

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
	r.POST("/TimerTrigger1", postTick)
	r.Run(getPort())
}
