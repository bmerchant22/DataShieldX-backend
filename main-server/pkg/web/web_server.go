package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

func CreateWebServer() *Server {
	srv := new(Server)
	srv.r = gin.Default()
	srv.r.Use(cors.Default())

	outFolderPath := "out"
	if _, err := os.Stat(outFolderPath); os.IsNotExist(err) {
		// If "out" folder doesn't exist, try "../out" (adjust as needed)
		outFolderPath = "../out"
	}

	srv.r.POST(kStartServer, srv.StartServerHandler)
	srv.r.POST(kStopServer, srv.StopServerHandler)
	srv.r.GET(kLogs, srv.LogsHandler)
	srv.r.POST(kStudentLogin, srv.LoginHandler)
	//srv.r.StaticFS("/_next", http.Dir(outFolderPath+"/_next"))
	//srv.r.StaticFS("/student", http.Dir(outFolderPath+"/student"))
	//srv.r.StaticFS("/client", http.Dir(outFolderPath+"/client"))
	//srv.r.POST(kList, srv.ListingFilesHandler)
	//srv.r.POST(kFile, srv.DisplayFileContent)

	if err := srv.r.Run("0.0.0.0:8000"); err != nil {
		zap.S().Errorf("Error while running the server !")
	}

	zap.S().Infof("Web server created successfully !!")

	return srv
}
