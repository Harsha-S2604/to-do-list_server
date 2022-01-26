package models

type Task struct {
	User		User
	TaskId		int
	TaskName 	string
	IsCompleted bool
}