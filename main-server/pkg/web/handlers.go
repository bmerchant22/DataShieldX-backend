package web

import (
	"fmt"
	"github.com/bmerchant22/server-controller/pkg/store"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type Server struct {
	r *gin.Engine
}

func (srv *Server) StartServerHandler(c *gin.Context) {
	teamStr := c.Param("team")
	portStr := c.Param("port")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		zap.S().Errorf("Error while converting portStr to port : %v", err)
	}
	team, err := strconv.Atoi(teamStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team parameter"})
		return
	}

	// Define root directory based on the team
	rootDir := ""
	switch team {
	case 1:
		rootDir = "Desktop/advent-of-code"
	case 2:
		rootDir = "Desktop/DataShieldX"
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team"})
		return
	}

	// Resolve the root directory path
	absRootDir, err := store.ResolvePath(rootDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	zap.S().Infof("Working root directory : %v", absRootDir)

	// Start the command
	containerID, err := store.StartDockerContainer(absRootDir, team, port)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Docker container started successfully for team %d", team), "containerID": containerID, "port": port})

}

//func (srv *Server) StopServerHandler(c *gin.Context) {
//	err := store.StopCodeServer()
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "Code server stopped successfully"})
//}
