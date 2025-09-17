package bot

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vas-sh/todo/internal/models"
)

func (s *srv) createTaskDruft(ctx context.Context, chatID int64) error {
	userID, err := s.userSrv.GetUserID(ctx, chatID)
	if err != nil {
		return err
	}
	err = s.taskSrv.DeleteTaskDruft(ctx, userID)
	if err != nil {
		return err
	}
	body := models.TaskDruft{
		UserID:     userID,
		UserStatus: string(models.WaitForTaskTitleUserStatus),
		Status:     models.NewStatus,
	}
	err = s.taskSrv.CreateTaskDruft(ctx, body)
	if err != nil {
		return err
	}
	keyboard := s.optionsKeyboard([]models.ButtonType{
		models.CancelButtonType,
	})
	return s.sendTextMassageWithKeyboard(chatID, "Please enter task title", keyboard)
}

func (s *srv) manageTextMassage(ctx context.Context, message tgbotapi.Message) error {
	chatID := message.From.ID
	userID, err := s.userSrv.GetUserID(ctx, chatID)
	if err != nil {
		return err
	}
	status, err := s.taskSrv.GetTaskDruftStatus(ctx, userID)
	if err != nil {
		return err
	}
	switch status {
	case models.WaitForTaskTitleUserStatus:
		return s.addTitleToDruft(ctx, userID, chatID, message.Text)
	case models.WaitForTaskDescriptionUserStatus:
		return s.addDescriptionToDruft(ctx, userID, chatID, message.Text)
	}
	return nil
}

func (s *srv) addTitleToDruft(ctx context.Context, userID, chatID int64, data string) error {
	body := models.TaskDruft{
		UserID:     userID,
		Title:      data,
		UserStatus: string(models.WaitForTaskDescriptionUserStatus),
	}
	err := s.updateDruft(ctx, body)
	if err != nil {
		return err
	}
	keyboard := s.optionsKeyboard([]models.ButtonType{
		models.CancelButtonType,
		models.CreateButtonType,
	})
	return s.sendTextMassageWithKeyboard(chatID, "Please enter task description", keyboard)
}

func (s *srv) addDescriptionToDruft(ctx context.Context, userID, chatID int64, data string) error {
	body := models.TaskDruft{
		UserID:      userID,
		Description: data,
		UserStatus:  string(models.WaitForNothingUserStatus),
	}
	err := s.updateDruft(ctx, body)
	if err != nil {
		return err
	}
	keyboard := s.statusKeyboard()
	return s.sendTextMassageWithKeyboard(chatID, "Please select task status", keyboard)
}

func (s *srv) addStatusToDruft(ctx context.Context, chatID int64, data string, messageID int) error {
	userID, err := s.userSrv.GetUserID(ctx, chatID)
	if err != nil {
		return err
	}
	body := models.TaskDruft{
		UserID: userID,
		Status: models.Status(data),
	}
	err = s.updateDruft(ctx, body)
	if err != nil {
		return err
	}
	return s.selectYear(ctx, chatID, messageID, "")
}

func (s *srv) selectYear(ctx context.Context, chatID int64, messageID int, data string) error {
	var year int
	if data == "" {
		year = time.Now().Year()
	} else {
		year = s.getYearFromCallback(ctx, data)
	}
	if year == 0 {
		year = time.Now().Year()
	}
	keyboard := s.yearKeyboard(year)
	return s.refreshTextMassageWithKeyboard(chatID, keyboard, messageID, "Please select year")
}

func (s *srv) getYearFromCallback(ctx context.Context, data string) int {
	parts := strings.Split(data, ":")
	if len(parts) != 2 {
		s.logger.ErrorContext(ctx, "failed get year from callback")
		return 0
	}
	dateStr := parts[1]
	date, err := strconv.Atoi(dateStr)
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error())
		return 0
	}
	return date
}

func (s *srv) getDateFromCallback(ctx context.Context, data string) string {
	parts := strings.Split(data, ":")
	if len(parts) != 2 {
		s.logger.ErrorContext(ctx, "failed get date from callback")
		return ""
	}
	return parts[1]
}

func (s *srv) selectMonth(ctx context.Context, chatID int64, messageID int, data string) error {
	year := s.getYearFromCallback(ctx, data)
	keyboard := s.monthKeyboard(year)
	return s.refreshTextMassageWithKeyboard(chatID, keyboard, messageID, "Please select year")
}

func (s *srv) selectDay(ctx context.Context, chatID int64, messageID int, data string) error {
	date := s.getDateFromCallback(ctx, data)
	keyboard := s.dayKeyboard(ctx, date)
	return s.refreshTextMassageWithKeyboard(chatID, keyboard, messageID, "Please select day")
}

func (s *srv) selectHour(ctx context.Context, chatID int64, messageID int, data string) error {
	date := s.getDateFromCallback(ctx, data)
	keyboard := s.hourKeyboard(date)
	return s.refreshTextMassageWithKeyboard(chatID, keyboard, messageID, "Please select hour")
}

func (s *srv) selectMinutes(ctx context.Context, chatID int64, messageID int, data string) error {
	date := s.getDateFromCallback(ctx, data)
	keyboard := s.minutesKeyboard(date, 1)
	return s.refreshTextMassageWithKeyboard(chatID, keyboard, messageID, "Please select minutes")
}

func (s *srv) addEstimateTimeToDruft(ctx context.Context, chatID int64, messageID int, data string) error {
	dateStr := s.getDateFromCallback(ctx, data)
	t, err := time.ParseInLocation("2006/1/2/15/04", dateStr, time.Local)
	if err != nil {
		return err
	}
	userID, err := s.userSrv.GetUserID(ctx, chatID)
	if err != nil {
		return err
	}
	err = s.taskSrv.UpdateTaskDruft(ctx,
		models.TaskDruft{
			UserID:       userID,
			EstimateTime: &t,
		})
	if err != nil {
		return err
	}
	return s.createTask(ctx, chatID, userID, messageID)
}

func (s *srv) createTask(ctx context.Context, chatID, userID int64, messageID int) error {
	err := s.taskSrv.CreateFromDruft(ctx, userID)
	if err != nil {
		return err
	}
	return s.refreshMassage(chatID, messageID, "Task created!")
}

func (s *srv) updateDruft(ctx context.Context, body models.TaskDruft) error {
	return s.taskSrv.UpdateTaskDruft(ctx, body)
}

func (s *srv) deleteTaskDruft(ctx context.Context, chatID int64) error {
	userID, err := s.userSrv.GetUserID(ctx, chatID)
	if err != nil {
		return err
	}
	return s.taskSrv.DeleteTaskDruft(ctx, userID)
}

func (s *srv) createFromDruft(ctx context.Context, chatID int64) error {
	userID, err := s.userSrv.GetUserID(ctx, chatID)
	if err != nil {
		return err
	}
	return s.taskSrv.CreateFromDruft(ctx, userID)
}

func (*srv) yearKeyboard(currentYear int) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"Prev", fmt.Sprintf("%s:%d", models.PrevYearButtonType, currentYear-1)),
			tgbotapi.NewInlineKeyboardButtonData(
				strconv.Itoa(currentYear), fmt.Sprintf("%s:%d",
					string(models.YearButtonType), currentYear)),
			tgbotapi.NewInlineKeyboardButtonData(
				"Next", fmt.Sprintf("%s:%d", models.NextYearButtonType, currentYear+1)),
		),
	)
}

func (*srv) monthKeyboard(year int) tgbotapi.InlineKeyboardMarkup {
	months := []string{
		"Jan", "Feb", "Mar",
		"Apr", time.May.String(), "Jun",
		"Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec",
	}
	rows := make([][]tgbotapi.InlineKeyboardButton, 0, (len(months)+3)/4)
	for i := 0; i < len(months); i += 4 {
		var row []tgbotapi.InlineKeyboardButton
		for j := 0; j < 4 && i+j < len(months); j++ {
			monthIndex := i + j
			row = append(row,
				tgbotapi.NewInlineKeyboardButtonData(
					months[monthIndex],
					fmt.Sprintf("%s:%d/%d", models.MonthButtonType, year, monthIndex+1),
				),
			)
		}
		rows = append(rows, row)
	}
	return tgbotapi.InlineKeyboardMarkup{InlineKeyboard: rows}
}

func (*srv) hourKeyboard(date string) tgbotapi.InlineKeyboardMarkup {
	rows := make([][]tgbotapi.InlineKeyboardButton, 0, (24+7)/8)
	for i := 0; i < 24; i += 8 {
		var row []tgbotapi.InlineKeyboardButton
		for j := 0; j < 8 && i+j < 24; j++ {
			hour := i + j
			row = append(row,
				tgbotapi.NewInlineKeyboardButtonData(
					fmt.Sprintf("%02d", hour),
					fmt.Sprintf("%s:%s/%d", models.HourButtonType, date, hour),
				),
			)
		}
		rows = append(rows, row)
	}
	return tgbotapi.InlineKeyboardMarkup{InlineKeyboard: rows}
}

func (s *srv) updateMinutesMessage(ctx context.Context, chatID int64, messageID int, data string) error {
	date, page := s.getDateAndPageFromCallback(ctx, data)
	keyboard := s.minutesKeyboard(date, page)
	return s.refreshTextMassageWithKeyboard(chatID, keyboard, messageID, "Please select minutes")
}

func (s *srv) getDateAndPageFromCallback(ctx context.Context, data string) (date string, page int) {
	parts := strings.Split(data, ":")
	if len(parts) != 3 {
		s.logger.ErrorContext(ctx, "failed get date and page")
		return
	}
	date = parts[2]
	pageStr := parts[1]
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error())
		return
	}
	return
}

func (*srv) minutesKeyboard(date string, page int) tgbotapi.InlineKeyboardMarkup {
	minMinutes, maxMinutes := (page-1)*20, page*20

	var rows [][]tgbotapi.InlineKeyboardButton
	var row []tgbotapi.InlineKeyboardButton
	for i := minMinutes; i < maxMinutes; i++ {
		button := tgbotapi.NewInlineKeyboardButtonData(
			fmt.Sprintf("%02d", i),
			fmt.Sprintf("%s:%s/%02d", models.MinutesButtonType, date, i),
		)
		row = append(row, button)
		if (i+1)%5 == 0 {
			rows = append(rows, row)
			row = []tgbotapi.InlineKeyboardButton{}
		}
	}
	if len(row) > 0 {
		rows = append(rows, row)
	}

	var optionsBtns []tgbotapi.InlineKeyboardButton

	if page > 1 {
		optionsBtns = append(optionsBtns, tgbotapi.NewInlineKeyboardButtonData(
			"Prev",
			fmt.Sprintf("%s:%d:%s", models.PrevPageButtonType, page-1, date)))
	}
	if page < 3 {
		optionsBtns = append(optionsBtns, tgbotapi.NewInlineKeyboardButtonData(
			"Next",
			fmt.Sprintf("%s:%d:%s", models.NextPageButtonType, page+1, date)))
	}

	rows = append(rows, optionsBtns)
	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func (s *srv) dayKeyboard(ctx context.Context, date string) tgbotapi.InlineKeyboardMarkup {
	year, month := s.getMonthAndYear(ctx, date)
	m := time.Month(month)
	firstOfNextMonth := time.Date(year, m+1, 1, 0, 0, 0, 0, time.UTC)
	lastOfMonth := firstOfNextMonth.AddDate(0, 0, -1)
	days := lastOfMonth.Day()
	rows := make([][]tgbotapi.InlineKeyboardButton, 0, (days+7)/8)
	for i := 1; i <= days; i += 8 {
		var row []tgbotapi.InlineKeyboardButton
		for j := 0; j < 8 && i+j <= days; j++ {
			day := i + j
			row = append(row,
				tgbotapi.NewInlineKeyboardButtonData(
					strconv.Itoa(day), fmt.Sprintf("%s:%s/%d", models.DayButtonType, date, day),
				),
			)
		}
		rows = append(rows, row)
	}
	return tgbotapi.InlineKeyboardMarkup{InlineKeyboard: rows}
}

func (s *srv) getMonthAndYear(ctx context.Context, date string) (year, month int) {
	parts := strings.Split(date, "/")
	if len(parts) != 2 {
		s.logger.ErrorContext(ctx, "incorrect format")
		return 0, 0
	}
	yearStr := parts[0]
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error())
		return 0, 0
	}
	monthStr := parts[1]
	month, err = strconv.Atoi(monthStr)
	if err != nil {
		s.logger.ErrorContext(ctx, err.Error())
		return 0, 0
	}
	return
}

func (*srv) optionsKeyboard(buttonTypes []models.ButtonType) tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton
	var buttons []tgbotapi.InlineKeyboardButton
	for _, item := range buttonTypes {
		switch item {
		case models.CancelButtonType:
			buttons = append(buttons,
				tgbotapi.NewInlineKeyboardButtonData("Cancel", string(item)))
		case models.CreateButtonType:
			buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData("Finish", string(models.CreateButtonType)))
		}
	}
	if len(buttons) > 0 {
		rows = append(rows, tgbotapi.NewInlineKeyboardRow(buttons...))
	}
	return tgbotapi.InlineKeyboardMarkup{InlineKeyboard: rows}
}

func (*srv) statusKeyboard() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🟡 New", string(models.NewStatus)),
			tgbotapi.NewInlineKeyboardButtonData("🟡 In progress", string(models.InProgressStatus)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🟢 Done", string(models.DoneStatus)),
			tgbotapi.NewInlineKeyboardButtonData("🔴 Canceled", string(models.CanceledStatus)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Cancel", string(models.CancelButtonType)),
			tgbotapi.NewInlineKeyboardButtonData("Finish", string(models.CreateButtonType)),
		),
	)
	return tgbotapi.NewInlineKeyboardMarkup(keyboard.InlineKeyboard...)
}
