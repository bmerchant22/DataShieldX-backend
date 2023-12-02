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
	logs, err := store.StartCodeServer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Code server started successfully", "logs": logs})
}

func (srv *Server) StopServerHandler(c *gin.Context) {
	err := store.StopCodeServer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Code server stopped successfully"})
}

func (srv *Server) LogsHandler(c *gin.Context) {
	logs := store.GetLatestLogs()
	c.String(http.StatusOK, logs)
}

func (srv *Server) ListingFilesHandler(c *gin.Context) {
	dir := c.Param("dir")
	absPath, err := store.ResolvePath(dir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Read the files in the directory
	files, err := store.GetDirDetails(absPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send the list of files as JSON response
	c.JSON(http.StatusOK, files)
}

func (srv *Server) DisplayFileContent(c *gin.Context) {
	// Resolve the file path
	path := c.Param("path")
	absPath, err := store.ResolvePath(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Read the contents of the file
	content, err := store.GetFileContent(absPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send the file content as JSON response
	c.JSON(http.StatusOK, gin.H{"content": string(content)})
}
