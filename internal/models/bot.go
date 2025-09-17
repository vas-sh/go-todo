package models

type ButtonType string

const (
	NextButtonType     ButtonType = "nextTask"
	PrevButtonType     ButtonType = "prevTask"
	CancelButtonType   ButtonType = "cancel"
	SkipButtonType     ButtonType = "skip"
	CreateButtonType   ButtonType = "create"
	YearButtonType     ButtonType = "year"
	NextYearButtonType ButtonType = "nextYear"
	PrevYearButtonType ButtonType = "prevYear"
	MonthButtonType    ButtonType = "month"
	DayButtonType      ButtonType = "day"
	HourButtonType     ButtonType = "hour"
	MinutesButtonType  ButtonType = "minutes"
	NextPageButtonType ButtonType = "nextPage"
	PrevPageButtonType ButtonType = "prevPage"
)

type UserStatus string

const (
	WaitForTaskTitleUserStatus       UserStatus = "waitForTitle"
	WaitForTaskDescriptionUserStatus UserStatus = "waitForDescription"
	WaitForNothingUserStatus         UserStatus = "waitForNothing"
)

type Command string

const (
	StartCommand      Command = "/start"
	TaskListCommand   Command = "/tasks"
	CreateTaskCommand Command = "/create"
)
