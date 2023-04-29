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

			Year: currentYear,
			Moon: currentMonth,

			Groceries:    []Expense{},
			Hobbies:      []Expense{},
			Rent:         []Expense{},
			OtherBills:   []Expense{},
			Travel:       []Expense{},
			Miscelanious: []Expense{},
		})

		monthIndex = months.FindMonthByYear(currentMonth, currentYear)
	}

	months.Months[monthIndex].TotalSpending += expense.Price
	switch expense.Category {
	case "Groceries":
		months.Months[monthIndex].Groceries = append(months.Months[monthIndex].Groceries, expense)

	case "Hobbies":
		months.Months[monthIndex].Hobbies = append(months.Months[monthIndex].Hobbies, expense)

	case "Rent":
		months.Months[monthIndex].Rent = append(months.Months[monthIndex].Rent, expense)

	case "Other bills":
		months.Months[monthIndex].OtherBills = append(months.Months[monthIndex].OtherBills, expense)

	case "Travel":
		months.Months[monthIndex].Travel = append(months.Months[monthIndex].Travel, expense)

	case "Miscelanious":
		months.Months[monthIndex].Miscelanious = append(months.Months[monthIndex].Miscelanious, expense)
	}

	j, _ := json.MarshalIndent(months, "", "    ")
	_ = ioutil.WriteFile("expenses.json", j, 0644)

	return nil
}
