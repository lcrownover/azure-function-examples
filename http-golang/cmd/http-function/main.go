package main

import (
	"database/sql"
	"net/http"
	"os"

	"log/slog"

	"github.com/gin-gonic/gin"
)

type App struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewApp() *App {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return &App{
		db:     nil, // TODO: connect to db
		logger: logger,
	}
}

func (a *App) getHandler(c *gin.Context) {
	msg := "This HTTP GET triggered function executed successfully."
	a.logger.Info(msg, slog.Int("status_code", 200), slog.String("status", "OK"))
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}

func (a *App) postHandler(c *gin.Context) {
	msg := "This HTTP POST triggered function executed successfully."
	a.logger.Info(msg, slog.Int("status_code", 200), slog.String("status", "OK"))
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}

func getPort() string {
	port := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = ":" + val
	}
	return port
}

func main() {
	app := NewApp()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/api/hello", app.postHandler)
	r.GET("/api/hello", app.getHandler)
	r.Run(getPort())
}
