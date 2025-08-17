package models

type ButtonType string

const (
	NextButtonType ButtonType = "nextTask"
	PrevButtonType ButtonType = "prevTask"
)

type Command string

const (
	StartCommand    Command = "/start"
	TaskListCommand Command = "/tasks"
)
