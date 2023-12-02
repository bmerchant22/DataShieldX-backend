package web

const (
	kStartServer = "/start/:team/:port"
	kStopServer  = "/stop"
	kLogs        = "/logs"
	kList        = "/list/*dir"
	kFile        = "/read/*path"
)
