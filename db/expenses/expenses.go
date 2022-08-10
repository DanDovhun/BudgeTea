package expenses

import (
	"BudgetManager/db/data"
	"errors"
	"time"
)

func GetMonthlyExpenses() ([]data.MonthlyExpense, error) {
	info, err := data.GetAllMonthlyExpenses()

	if err != nil {
		return info, err
	}

	return info, nil
}

func GetThisMonthsExpenses() (data.MonthlyExpense, error) {
	info, err := data.GetAll()

	if err != nil {
		return data.MonthlyExpense{}, err
	}

	currentTime := time.Now()
	year, month, _ := currentTime.Date()

	var mon string

	switch month {
	case 1:
		mon = "Jan"

	case 2:
		mon = "Feb"

	case 3:
		mon = "Mar"

	case 4:
		mon = "Apr"

	case 5:
		mon = "May"

	case 6:
		mon = "Jun"

	case 7:
		mon = "Jul"

	case 8:
		mon = "Aug"

	case 9:
		mon = "Sep"

	case 10:
		mon = "Oct"

	case 11:
		mon = "Nov"

	case 12:
		mon = "Dec"
	}

	expenseIndex := info.FindMonthlyExpense(year, mon)

	if expenseIndex == -1 {
		return data.MonthlyExpense{}, errors.New("You didn't spend any money this month")
	}

	return info.Expenses[expenseIndex], nil
}
