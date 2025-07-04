package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reagin/double-ratchet/config"
	"github.com/reagin/double-ratchet/server/handlers"
	"github.com/reagin/double-ratchet/server/websocket"
)

var ServiceServer *http.Server

func init() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()
	router.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Content-Length")
		if ctx.Request.Method == http.MethodOptions {
			ctx.Status(http.StatusNoContent)
			return
		}
		ctx.Next()
	})

	// Don't Need Authorization Header
	routerGroup := router.Group("/api")
	// Upgrade GET method to WebSocket Connect
	routerGroup.GET("/websocket", websocket.HandleWebSocket)
	// Router methods below
	routerGroup.POST("/auth/valid", handlers.HandleAuthValid)
	routerGroup.POST("/auth/login", handlers.HandleAuthLogin)
	routerGroup.POST("/auth/forgot", handlers.HandleAuthForgot)
	routerGroup.POST("/auth/register", handlers.HandleAuthRegister)

	ServiceServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.GIN_SERVER, config.GIN_PORT),
		Handler: router,
	}
}
