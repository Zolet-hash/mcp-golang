
package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func NewRouter(mcpServer *mcp.Server) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	handler := mcp.NewStreamableHTTPHandler(
		func(r *http.Request) *mcp.Server {
			return mcpServer
		},
		&mcp.StreamableHTTPOptions{
			JSONResponse: true,
			Stateless:    true,
		},
	)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.Any("/mcp", func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	})

	return router
}