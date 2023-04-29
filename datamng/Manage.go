package datamng

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

func Add(expense Expense) error {
	data, err := os.Open("expenses.json")
	var months Data

	if err != nil {
		return err
	}

	defer data.Close()

	byteValue, _ := ioutil.ReadAll(data)
	json.Unmarshal(byteValue, &months)

	currentYear, currentMonth, _ := time.Now().Date()

	monthIndex := months.FindMonthByYear(currentMonth, currentYear)

	if monthIndex == -1 {
		months.Months = append(months.Months, Month{
			Budget:        months.Budget,
			TotalSpending: 0,
			Year:          currentYear,
			Moon:          currentMonth,
			Expenses:      []Expense{},
		})

		monthIndex = months.FindMonthByYear(currentMonth, currentYear)
	}

	months.Months[monthIndex].TotalSpending += expense.Price
	months.Months[monthIndex].Expenses = append(months.Months[monthIndex].Expenses, expense)

	j, _ := json.MarshalIndent(months, "", "    ")
	_ = ioutil.WriteFile("expenses.json", j, 0644)

	return nil
}
