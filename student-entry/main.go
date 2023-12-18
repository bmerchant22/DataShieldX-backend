package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
	"encoding/json"
	"io"
	"errors"
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
	"github.com/gin-contrib/cors"
)

type Users struct {
	Team string		`json:"team"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type App struct {
    App_Name         string `bson:"app_name" json:"app_name"`
    App_Desc         string `bson:"app_desc" json:"app_desc"`
    Approval_Status  string `bson:"approval_status" json:"approval_status"`
}

type AppRequest struct {
	Project_Id string `json:"project_id"`
	Apps []App	`json:"apps"`
}

const (
	base_url = "0.0.0.0:8000" //where this will listen
	server_controller = "http://localhost:8001" //where this will request the server list from
)

var teamAddrs map[string]string //map of teams -> their servers
var prevReqTime time.Time
var apps []App

var allUsers = []Users{
	{Team: "1", Username: "user1", Password: "pass1"},
	{Team: "2", Username: "user2", Password: "pass2"},
	{Team: "3", Username: "user3", Password: "pass3"},
}
//for directing requests to correct server
func getLoadBalanceAddr(team string) (val string, ok bool) {
	val, ok = teamAddrs[team]
	return
}

func updateAddrs() (err error){
	log.Println("updating server list")
	// defer log.Println("exited update addrs", err)
	resp, err :=http.Get(fmt.Sprintf("%s/servers",server_controller))
	defer log.Println("err:", err)
// 	log.Println(err)
// 	log.Println(resp.StatusCode)
	if (err != nil) {return err}
	if (resp.StatusCode != http.StatusOK) {
		err = errors.New(resp.Status)
		return
	}
	defer resp.Body.Close()
	log.Println(resp.Status)
	
	var respBody []byte
	respBody, err = io.ReadAll(resp.Body)
	if (err != nil) {return err}
	//clear teamAddrs, but back it up first
	backup := make(map[string]string)
	for key, val := range teamAddrs {
		backup[key] = val
		delete(teamAddrs, key)
	}
	//set teamAddrs to response
	err = json.Unmarshal(respBody, &teamAddrs)
	if (err != nil) {
		//restore backup
		for key, val := range backup {
			teamAddrs[key] = val
		}
		return err
	}
	//done
	log.Printf("%v", teamAddrs)
	return nil
}

// func getUsersFromController() (err error) {
// 	resp, err := http.Get(fmt.Sprintf("%s/getUsers",server_controller))
// 	if (err != nil) {return}
// 	if resp.StatusCode != http.StatusOK {
// 		err = errors.New(resp.Status)
// 		return
// 	}
// 	defer resp.Body.Close()
	
// 	var respBody []byte
// 	respBody, err = io.ReadAll(resp.Body)
// 	if (err != nil) {return}
// 	//set users to response
// 	err = json.Unmarshal(respBody, &allUsers)
// 	log.Println(allUsers)
// 	return
// }

func getAppsFromController() (err error) {
	log.Println("Getting apps")
	resp, err := http.Get(fmt.Sprintf("%s/getApps",server_controller))
	if (err != nil) {return}
	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		return
	}
	defer resp.Body.Close()
	
	var respBody []byte
	respBody, err = io.ReadAll(resp.Body)
	log.Println(string(respBody))
	if (err != nil) {return}
	//set users to response
	err = json.Unmarshal(respBody, &apps)
	log.Println(err)
	log.Println("gotten apps")
	log.Println(apps)
	return
}

func getAppsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, apps);
}

func getProjectHandler(c *gin.Context) { //input: teamid in URL params, check teamid then forward straight to controller
	cred, err := c.Cookie("creds")
	if (err != nil) {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	creds := strings.Split(cred, ":")
	if (len(creds) != 3) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cookie invalid"})
		return
	}
	reqUser := Users{creds[0], creds[1], creds[2]}
	found := false
	for _, user := range allUsers {
		if (user == reqUser) {
			found = true
			break
		}
	}
	if (!found) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	//valid user, forward request
	log.Println(fmt.Sprintf("%s/%s?%s",server_controller,c.Request.URL.Path,c.Request.URL.RawQuery))
	resp, err := http.Get(fmt.Sprintf("%s/%s?%s",server_controller,"getProject",c.Request.URL.RawQuery))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var respbody []byte 
	respbody, err = io.ReadAll(resp.Body)
	if (err != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"error":err})
	}
	c.Data(resp.StatusCode, "application/json", respbody)
} 

func requestAppsHandler(c *gin.Context) { //input: teamid, array of apps to request
	//check creds
	
	cred, err := c.Cookie("creds")
	if (err != nil) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	creds := strings.Split(cred, ":")
	if (len(creds) != 3) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cookie invalid"})
		return
	}
	reqUser := Users{creds[0], creds[1], creds[2]}
	found := false
	for _, user := range allUsers {
		if (user == reqUser) {
			found = true
			break
		}
	}
	if (!found) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	//valid creds, validate request body
	var req AppRequest
	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err})
		return
	}
	if req.Project_Id != reqUser.Team {
		c.JSON(http.StatusBadRequest, gin.H{"error":"unauthorized"})
		return
	}
	for _, app := range req.Apps {
		if app.Approval_Status != "requested" {
			c.JSON(http.StatusBadRequest, gin.H{"error":"unauthorized"})
			return
		}
	}
	//valid request, send to server controller
	jsonbody, _ := json.Marshal(req)
	resp, _ := http.Post(fmt.Sprintf("%s/updateProject", server_controller),"application/json", bytes.NewBuffer(jsonbody))
	defer resp.Body.Close()
// 	if (resperr != nil) {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": resperr})
// 		return
// 	}
	// respbody := io.Reader(resp.Body)
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		// Handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(resp.StatusCode,"application/json",respBody)
}



func appReverseProxyHandler(c *gin.Context) {

	//first check cookie to see if valid
	cred, err := c.Cookie("creds")
	if (err != nil) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	creds := strings.Split(cred, ":")
	if (len(creds) != 3) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cookie invalid"})
		return
	}
	reqUser := Users{creds[0], creds[1], creds[2]}
	found := false
	for _, user := range allUsers {
		if (user == reqUser) {
			found = true
			break
		}
	}
	if (!found) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	
	//check time and update teamAddrs if necessary
	//update max every 10s
	if (time.Now().Sub(prevReqTime) > 10000) {
		err := updateAddrs()
		if (err != nil) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	//valid user, so set correct proxy
	addr, ok := getLoadBalanceAddr(reqUser.Team)
	if (!ok) {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"server not operating currently"})
		return
	}

	lastDigit := addr[len(addr)-1]
	newAddr := "http://team-"+string(lastDigit)+"-workspace:808"+string(lastDigit)
	log.Println(newAddr)

	proxy, err := url.Parse(newAddr)
	if err != nil {
		log.Printf("error in parse addr: %v", err)
		c.String(500, "error")
		return
	}


	log.Println("accessed path")
	log.Println(c.Request.URL)
	log.Println(c.Param("path"))
	// step 1: resolve proxy address, change scheme and host in requets
	req := c.Request
	req.URL.Scheme = proxy.Scheme
	req.URL.Host = proxy.Host
	req.URL.Path = c.Param("path")
	
	log.Println("request URL:", req.URL.Scheme, req.URL.Host, req.URL.Path, req.URL.RawQuery)
	
	rp := new(httputil.ReverseProxy)
	rp.Director = func(hreq *http.Request) {
		hreq = req
	 //  hreq.URL, err = url.Parse(fmt.Sprintf("%s%s",proxy,c.Param("path")))
	 //  log.Println("in URL:", c.Request.URL, "hreq URL:", hreq.URL)
		
		
	}
	rp.ServeHTTP(c.Writer, c.Request)
}

func authHandler(c *gin.Context) {
	//once logged in, set cookie
	var requestUser Users
	if err := c.ShouldBindJSON(&requestUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("user:",requestUser.Team,requestUser.Username,requestUser.Password)
	var foundUser *Users
	for _, user := range allUsers {
		if user == requestUser {
			foundUser = &user
			break
		}
	}
	// check if the user was found
	if foundUser == nil {
		log.Println("user not found:",requestUser.Team,requestUser.Username,requestUser.Password)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	
	//user found, so set cookie to their credentials (this is a prototype)
	c.SetCookie("creds", fmt.Sprintf("%s:%s:%s", requestUser.Team,requestUser.Username,requestUser.Password), 86400, "/", base_url, false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

var test int

func testt() {
	test = 5
}

func testtt() {
	log.Println(test)
}
func main() {
	test = 1
	testtt()
	log.Println(test)
	teamAddrs = make(map[string]string)
	// getUsersFromController()
	log.Println("Error msg ------------------")
	log.Println(allUsers)
	getAppsFromController()
	prevReqTime = time.Now()
	updateAddrs()
	r := gin.Default()
	
	r.Use(cors.Default())
	r.Use(static.Serve("/", static.LocalFile("./out", true)))
	//serve static files like /student/login, only check auth when handling apps
	
	r.GET("/app/code/*path", appReverseProxyHandler)
	r.GET("/api/getApps", getAppsHandler)
	r.GET("/api/getProject", getProjectHandler)
	r.POST("/api/requestApps", requestAppsHandler)
	r.POST("/api/auth", authHandler)
	
	
	
	if err := r.Run(base_url); err != nil {
		log.Printf("Error: %v", err)
	}
}