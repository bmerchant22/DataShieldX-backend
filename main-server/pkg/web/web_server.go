package web

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateWebServer() *Server {
	srv := new(Server)
	srv.r = gin.Default()

	srv.r.POST(kStartServer, srv.StartServerHandler)
	//srv.r.POST(kStopServer, srv.StopServerHandler)
	//srv.r.POST(kLogs, srv.LogsHandler)
	//srv.r.POST(kList, srv.ListingFilesHandler)
	//srv.r.POST(kFile, srv.DisplayFileContent)

	if err := srv.r.Run("localhost:8000"); err != nil {
		zap.S().Errorf("Error while running the server !")
	}

	zap.S().Infof("Web server created successfully !!")

	return srv
}
