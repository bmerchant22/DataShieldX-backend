package web

import (
	"github.com/bmerchant22/pkg/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	r *gin.Engine
}

func (srv *Server) StartServerHandler(c *gin.Context) {
	err := store.StartCodeServer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Code server started successfully"})
}

func (srv *Server) StopServerHandler(c *gin.Context) {
	err := store.StopCodeServer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Code server stopped successfully"})
}
