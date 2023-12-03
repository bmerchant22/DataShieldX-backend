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
	team := c.Param("team")
	portStr := c.Param("port")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		zap.S().Errorf("Error while converting portStr to port : %v", err)
	}

	// Define root directory based on the team
	rootDir := ""
	switch team {
	case "1":
		rootDir = "Desktop/advent-of-code"
	case "2":
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

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Docker container started successfully for team %d", team), "containerID": containerID, "port": port, "rootDir": absRootDir})

}

func (srv *Server) StopServerHandler(c *gin.Context) {
	team := c.Param("team")
	containerName := store.NameContainer(team)

	err := store.StopDockerContainer(containerName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Docker container stopped successfully", "containerName": containerName})
}

func (srv *Server) LogsHandler(c *gin.Context) {
	team := c.Param("team")

	teamName := store.NameContainer(team)
	// Get container ID by team name
	containerID, err := store.GetContainerIDByTeamName(teamName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get logs for the container
	logs, err := store.GetContainerLogs(containerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"team": teamName, "logs": logs})
}
