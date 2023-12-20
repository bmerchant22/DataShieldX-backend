package web

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
	"github.com/bmerchant22/DataShieldX-backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	openai "github.com/sashabaranov/go-openai"
	"context"
	"strconv"
	"strings"
	// "os"
)

//* Get ENV variables 
var mongoURI 				= "mongodb+srv://deven:neved@cluster0.h5zvxkm.mongodb.net/?retryWrites=true&w=majority"			//os.Getenv("MONGO_URI")
var mongoDBName 			= 	"trumio-backend"		//os.Getenv("MONGO_DB_NAME")
var projectsCollectionName 	= "projects-collection"			//os.Getenv("PROJECTS_COLLECTION_NAME")
// var usersCollectionName := 			//os.Getenv("USERS_COLLECTION_NAME")
var appsCollectionName = "apps-collection" 			//os.Getenv("APPS_COLLECTION_NAME")
const openaiAPIKey = "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"
const YYYYMMDD = "2006/01/02"

var MongoClient *mongo.Client

//* Create MongoDB client and connect to the database
func CreateMongoClient() (*mongo.Client, error) {
	// Create a new client
// 	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Connect to the database
// 	err = client.Connect(context.Background())
// 	if err != nil {
// 		return nil, err
// 	}
// 	return client, nil
		return MongoClient, nil
}

func ActuallyCreateMongoClient() (error) {
	// Create a new client
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}
	// Connect to the database
	err = client.Connect(context.Background())
	if err != nil {
		return err
	}
	//set mongoClient
	MongoClient = client
	return  nil
}

//* Generic function to check if a document with a key exists in the collection
//* Returns True if the document exists, False otherwise
func CheckIfDocumentExists(client *mongo.Client, databaseName string, collectionName string, key string, value string) bool {
	// Get the collection
	collection := client.Database(databaseName).Collection(collectionName)
	// Create a filter
	filter := bson.D{{key, value}}
	// Check if the document exists
	if err := collection.FindOne(context.Background(), filter).Err(); err != nil {
		return false
	}
	return true
}

//* Generic function to insert a document into the collection
//* Returns True if the document was inserted successfully, False otherwise
func InsertDocument(client *mongo.Client, databaseName string, collectionName string, document interface{}) (bool,  error) {
	// Get the collection
	collection := client.Database(databaseName).Collection(collectionName)
	// Insert the document
	_, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		return false, err
	}
	return true, nil
}

//* Generic function to delete a document from the collection
//* Returns True if the document was deleted successfully, False otherwise
func DeleteDocument(client *mongo.Client, databaseName string, collectionName string, key string, value string) (bool, error) {
	// Get the collection
	collection := client.Database(databaseName).Collection(collectionName)
	// Create a filter
	filter := bson.D{{key, value}}
	// Delete the document
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return false, err
	}
	return true, nil
}

//* Generic function to update a document in the collection
//* Returns True if the document was updated successfully, False otherwise
func UpdateDocument(client *mongo.Client, databaseName string, collectionName string, id_key string, id_value string, upsert_doc interface{}) (bool, error) {
	// Get the collection
	collection := client.Database(databaseName).Collection(collectionName)
	// Create a filter
	filter := bson.D{{id_key, id_value}}
	// Create an update
	
	update := bson.D{{"$set", upsert_doc}}
	// Update the document
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return false, err
	}
	return true, nil
}


// Create Project Handler
// TODO : Fix the Milestones, Team and Apps fields when pushed to MongoDb
func (srv *Server) CreateProjectHandler(c *gin.Context) {

	fmt.Printf("Mongo URI: %s\n", mongoURI)
	// Get the JSON request body & parse it as model Project
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if project.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project name is required"})
		return
	}

	// Print the project details
	fmt.Printf("Project ID: %s\n", project.Project_ID)
	fmt.Printf("Project Name: %s\n", project.Name)
	fmt.Printf("Project Description: %s\n", project.Project_Desc)
	fmt.Printf("Milestones: %s\n", project.Milestones)
	fmt.Printf("Team: %s\n", project.Team)
	fmt.Printf("Apps: %s\n", project.Apps)

	// Create a new client and connect to the server
	mongoClient, err := CreateMongoClient()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the project already exists
	if CheckIfDocumentExists(mongoClient, mongoDBName, projectsCollectionName, "project_id", project.Project_ID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project already exists"})
		return
	}

	// Insert the project into the database
	inserted, err := InsertDocument(mongoClient, mongoDBName, projectsCollectionName, project)
	if err != nil || !inserted {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Return the response
	c.JSON(http.StatusOK, gin.H{"message": "Project created successfully"})
}

// Delete Project Handler
func (srv *Server) DeleteProjectHandler(c *gin.Context) {

	// Get the JSON request body & parse it as model Project
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Print the project details
	// fmt.Printf("Project ID: %s\n", project.Project_ID)

	// Create a new client and connect to the server
	mongoClient, err := CreateMongoClient()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the project exists
	if !CheckIfDocumentExists(mongoClient, mongoDBName, projectsCollectionName, "project_id", project.Project_ID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project doesn't exist"})
		return
	}

	// Delete the project from the database
	deleted, err := DeleteDocument(mongoClient, mongoDBName, projectsCollectionName, "project_id", project.Project_ID)
	if err != nil || !deleted {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Return the response
	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}

// Update Project Handler using Upsert functionality
func (srv *Server) UpdateProjectHandler(c *gin.Context) {
	// Get the JSON request body & parse it as model Project
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "project": project})
		return
	}
	// Create a new client and connect to the server
	mongoClient, err := CreateMongoClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Check if the project exists
	if !CheckIfDocumentExists(mongoClient, mongoDBName, projectsCollectionName, "project_id", project.Project_ID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project doesn't exist"})
		return
	}

	// Update the project in the database
	updated, err := UpdateDocument(mongoClient, mongoDBName, projectsCollectionName, "project_id", project.Project_ID, project)
	if err != nil || !updated {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// Return the response
	c.JSON(http.StatusOK, gin.H{"message": "Project updated successfully"})

}


// Get Projects Handler
func (srv *Server) GetProjectsHandler(c *gin.Context) {
	
	// Create a new client and connect to the server
	mongoClient, err := CreateMongoClient()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
		return
	}

	// Get all the projects from the database
	var projects []models.Project
	collection := mongoClient.Database(mongoDBName).Collection(projectsCollectionName)
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error2": err.Error()})
		return
	}
	if err = cursor.All(context.Background(), &projects); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error3": err.Error()})
		return
	}

	// Return the response
// 	ret, err2 := json.Marshal(gin.H{"projects": projects})
	
// 	if (err2 != nil) {c.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})}
	c.JSON(http.StatusOK, gin.H{"projects": projects})

}

// Get Project Handler
func (srv *Server) GetProjectHandler(c *gin.Context) {
	
	//get requests aren't supposed to have bodies so we query by url params (the param is id)
	var project models.Project
	var ok bool
	// if err := c.ShouldBindJSON(&project); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
// 		return
// 	}
	defer func() {
		if r := recover(); r != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": r})
			return
		}
	}()
	project.Project_ID, ok = c.GetQuery("id");
	if (!ok) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing id"});
	}

	// Print the project details
	// fmt.Printf("Project ID: %s\n", project.Project_ID)

	// Create a new client and connect to the server
	mongoClient, err := CreateMongoClient()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error2": err.Error()})
		return
	}

	// Check if the project exists
	if !CheckIfDocumentExists(mongoClient, mongoDBName, projectsCollectionName, "project_id", project.Project_ID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project doesn't exist"})
		return
	}

	// Get the project from the database
	var projectFromDB models.Project
	collection := mongoClient.Database(mongoDBName).Collection(projectsCollectionName)
	filter := bson.D{{"project_id", project.Project_ID}}
	if err := collection.FindOne(context.Background(), filter).Decode(&projectFromDB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error3": err.Error()})
		return
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{"project": projectFromDB})

}


// Get Apps Handler
func (srv *Server) GetAppsHandler(c *gin.Context) {
	
	// Create a new client and connect to the server
	mongoClient, err := CreateMongoClient()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get all the apps from the database
	var apps []models.App
	collection := mongoClient.Database(mongoDBName).Collection(appsCollectionName)
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = cursor.All(context.Background(), &apps); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the response
	c.JSON(http.StatusOK, apps)
}

func (srv *Server) GenerateMilestoneHandler (c *gin.Context) {
	//in request: a json with fields name and project_desc - unmarshal and send to model, receive milestones in json format and return
	var request struct {
		ProjectName string `json:"project_name"`
		ProjectDesc string `json:"project_desc"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	gpt := openai.NewClient(openaiAPIKey)

	resp, err := gpt.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
					Model: openai.GPT3Dot5Turbo,
					Messages: []openai.ChatCompletionMessage{
							{
									Role:    "user",
									Content: fmt.Sprintf(`You are a project milestone generator assistant. And you stick to the point & provide very concise milestone breakdown of an idea overview given to you. You first plan the roadmap based on the idea overview provided and then you always output the milestones in the following format as a list. The best thing about you is that you strictly follow the format given below: 

Milestone 1
- Description: <fill-here>
- Number of days required: <fill-here>
Milestone 2
- Description: <fill-here>
- Number of days required: <fill-here>
.
.
.
Milestone n
- Description: <fill-here>
- Number of days required: <fill-here>
------

Following is the idea overview. Please make sure to stick to the format and provide me with concise milestones that are only one line and one line only.
Project Name: %s
Project Description: %s`,request.ProjectName, request.ProjectDesc),
							},
					},
			},
	)

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error calling OpenAI API"})
        return
    }

    // Extract the response from the OpenAI API
    var response string
    if len(resp.Choices) > 0 {
        response = resp.Choices[0].Message.Content
    } else {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No choices found in OpenAI response"})
        return
    }
		//split response
		// c.JSON(http.StatusOK, gin.H{"data": response})
		dirtyAllItems := strings.Split(response,"\n")
		var allItems []string //remove empty strings due to \n\n
		for _,item := range dirtyAllItems {
			if (item != "") {
				allItems = append(allItems, item)
			}
		}
// 		c.JSON(http.StatusOK, gin.H{"data": allItems})
		ret := make([]models.Milestone, 0)
		currDate := time.Now()
		var newMilestone *models.Milestone
		for idx, item := range allItems {
			fmt.Println(idx)
			fmt.Println(item)
			switch (idx%3) {
				case 0:
					newMilestone = new(models.Milestone)
					newMilestone.Tasks = make([]models.Task, 0)
					newMilestone.Milestone_ID = fmt.Sprintf("%v",idx/3 + 1)
				case 1:
					newMilestone.Milestone_Desc = strings.Split(item, ": ")[1]
				case 2:
					days,_ := strconv.Atoi(strings.Split(item,": ")[1])
					currDate = currDate.AddDate(0,0,days)
					newMilestone.Completion_Date = currDate.Format(YYYYMMDD)
					ret = append(ret, *newMilestone)
			}
		}
		
    // Return the response to the frontend
    c.JSON(http.StatusOK, gin.H{"milestones": ret})
}

func (srv *Server) GenerateTasksHandler (c *gin.Context) {
	//in request: a json with fields name and project_desc - unmarshal and send to model, receive milestones in json format and return
	var request struct {
		MilestoneDesc string `json:"milestone_desc"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	gpt := openai.NewClient(openaiAPIKey)

	resp, err := gpt.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
					Model: openai.GPT3Dot5Turbo,
					Messages: []openai.ChatCompletionMessage{
							{
									Role:    "user",
									Content: fmt.Sprintf(`You are a task generator assistant to achieve a particular milestone. And you stick to the point & provide very concise task breakdown of an idea overview given to you. You first plan the roadmap based on the milestone provided and then you always output the tasks in the following format as a list. The best thing about you is that you strictly follow the format given below: 

Task 1
- Description: <fill-here>
- Start-date: <fill-here>
- End-date: <fill-here>
Task 2
- Description: <fill-here>
- Start-date: <fill-here>
- End-date: <fill-here>
.
.
.
Task n
- Description: <fill-here>
- Start-date: <fill-here>
- End-date: <fill-here>
------

Following is the task overview. Please make sure to stick to the format and provide me with concise tasks that are only one line and one line only.
Milestone Description: %s`,request.MilestoneDesc),
							},
					},
			},
	)

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error calling OpenAI API"})
        return
    }

    // Extract the response from the OpenAI API
    var response string
    if len(resp.Choices) > 0 {
        response = resp.Choices[0].Message.Content
    } else {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No choices found in OpenAI response"})
        return
    }
		//split response
		// c.JSON(http.StatusOK, gin.H{"data": response})
		dirtyAllItems := strings.Split(response,"\n")
		var allItems []string //remove empty strings due to \n\n
		for _,item := range dirtyAllItems {
			if (item != "") {
				allItems = append(allItems, item)
			}
		}
// 		c.JSON(http.StatusOK, gin.H{"data": allItems})
		ret := make([]models.Task, 0)
		var newTask *models.Task
		for idx, item := range allItems {
			fmt.Println(idx)
			fmt.Println(item)
			switch (idx%4) {
				case 0:
					newTask = new(models.Task)
					newTask.Status = "pending"
					newTask.Assignees = make([]models.User, 0)
					newTask.Task_ID = fmt.Sprintf("%v",idx/4 + 1)
				case 1:
					newTask.Task_Desc = strings.Split(item, ": ")[1]
				case 2:
					start,_ := strconv.Atoi(strings.Split(item,": ")[1])
					newTask.Start_Time = time.Now().AddDate(0,0,start).Format(YYYYMMDD)
				case 3:
					end,_ := strconv.Atoi(strings.Split(item,": ")[1])
					newTask.End_Time = time.Now().AddDate(0,0,end).Format(YYYYMMDD)
					ret = append(ret, *newTask)
			}
		}
		
    // Return the response to the frontend
    c.JSON(http.StatusOK, gin.H{"tasks": ret})
}