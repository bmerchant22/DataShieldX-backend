package models

type User struct {
    Team     int    `bson:"team" json:"team"`
    Username string `bson:"username" json:"username"`
    Password string `bson:"password" json:"password"`
}

type Project struct {
    Project_ID    string      `bson:"project_id" json:"project_id" binding:"required"`
    Name          string      `bson:"name" json:"name"`
    Project_Desc  string      `bson:"project_desc" json:"project_desc"`
    Milestones    []Milestone `bson:"milestones" json:"milestones"`
    Team          []User      `bson:"team" json:"team"`
    Apps          []App       `bson:"apps" json:"apps"`
}

type Milestone struct {
    Milestone_ID    string `bson:"milestone_id" json:"milestone_id"`
    Milestone_Desc  string `bson:"milestone_desc" json:"milestone_desc"`
    Completion_Date string `bson:"completion_date" json:"completion_date"`
    Tasks           []Task `bson:"tasks" json:"tasks"`
}

type Task struct {
    Task_ID     string  `bson:"task_id" json:"task_id"`
    Task_Desc   string  `bson:"task_desc" json:"task_desc"`
    Start_Time  string  `bson:"start_time" json:"start_time"`
    End_Time    string  `bson:"end_time" json:"end_time"`
    Status      string  `bson:"status" json:"status"`
    Assignees   []User  `bson:"assignees" json:"assignees"`
}

type App struct {
    App_Name         string `bson:"app_name" json:"app_name"`
    App_Desc         string `bson:"app_desc" json:"app_desc"`
    Approval_Status  string `bson:"approval_status" json:"approval_status"`
}