package models

type User struct {
    Team     int    `bson:"team"`
    Username string `bson:"username"`
    Password string `bson:"password"`
}

type Project struct {
    Project_ID    string      `bson:"project_id" binding:"required"`
    Name          string      `bson:"name"`
    Project_Desc  string      `bson:"project_desc"`
    Milestones    []Milestone `bson:"milestones"`
    Team          []User      `bson:"team"`
    Apps          []App         `bson:"apps"`
}

type Milestone struct {
    Milestone_ID    string `bson:"milestone_id"`
    Milestone_Desc  string `bson:"milestone_desc"`
    Completion_Date string `bson:"completion_date"`
    Tasks           []Task `bson:"tasks"`
}

type Task struct {
    Task_ID     string  `bson:"task_id"`
    Task_Desc   string  `bson:"task_desc"`
    Start_Time  string  `bson:"start_time"`
    End_Time    string  `bson:"end_time"`
    Status      string  `bson:"status"`
    Assignees   []User  `bson:"assignees"`
}

type App struct {
    App_Name         string `bson:"app_name"`
    App_Desc         string `bson:"app_desc"`
    Approval_Status  string `bson:"approval_status"`
}