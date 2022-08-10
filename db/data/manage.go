package data

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

func Add(price float64, category string) error {
	data, err := os.Open("db/data/data.json")

	if err != nil {
		return errors.New("Cannot open file")
	}

	byteValue, _ := ioutil.ReadAll(data)
	var items Data
	json.Unmarshal(byteValue, &items)

	currentTime := time.Now()
	year, month, day := currentTime.Date()

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

	item := Item{
		Category: category,
		Price:    price,
		Year:     year,
		Month:    mon,
		Day:      day,
	}

	monthlyExpenseIndex := items.FindMonthlyExpense(year, mon)

	if monthlyExpenseIndex == -1 {
		items.Expenses = append(items.Expenses, MonthlyExpense{
			Year:              year,
			Month:             mon,
			TotalExpenses:     0,
			GroceriesExpenses: 0,
			TravelExpenses:    0,
			OtherExpenses:     0,
		})

		monthlyExpenseIndex = items.FindMonthlyExpense(year, mon)
	}

	items.Expenses[monthlyExpenseIndex].TotalExpenses += price

	switch category {
	case "Groceries":
		items.Groceries.Items = append(items.Groceries.Items, item)
		items.Groceries.Price += price
		items.Expenses[monthlyExpenseIndex].GroceriesExpenses += price

	case "Travel":
		items.Travel.Items = append(items.Travel.Items, item)
		items.Travel.Price += price
		items.Expenses[monthlyExpenseIndex].TravelExpenses += price

	case "Others":
		items.Others.Items = append(items.Others.Items, item)
		items.Others.Price += price
		items.Expenses[monthlyExpenseIndex].TravelExpenses += price
	}

	items.TotalPrice += price

	j, _ := json.MarshalIndent(items, "", "    ")
	_ = ioutil.WriteFile("db/data/data.json", j, 0644)

	return nil
}

func GetAll() (Data, error) {
	data, err := os.Open("db/data/data.json")

	if err != nil {
		return Data{}, errors.New("Cannot open file")
	}

	defer data.Close()

	byteValue, _ := ioutil.ReadAll(data)

	var items Data
	json.Unmarshal(byteValue, &items)

	return items, nil
}

func GetAllMonthlyExpenses() ([]MonthlyExpense, error) {
	data, err := GetAll()

	if err != nil {
		return []MonthlyExpense{}, err
	}

	return data.Expenses, nil
}

func (dt Data) FindMonthlyExpense(year int, month string) int {
	for i, j := range dt.Expenses {
		if j.Year == year && j.Month == month {
			return i
		}
	}

	return -1
}
