package domain

import "time"

type Task struct {
    ld           uint64 
	Userld       uint64
    Title        string 
    Description  string     
    Status       TaskStatus  
    Date         time.Time
	CreatedDate  time. Time
	UpdatedDate   time. Time
	DeletedDate  *time. Time
}

	type TaskStatus string
const (
NewTaskStatus          TaskStatus = "NEW"
DoneTaskStatus         TaskStatus ="DONE"
ImportantTaskStatus    TaskStatus = "IMPORTANT"
ExpiredTaskStatus      TaskStatus =" EXPIRED "
);