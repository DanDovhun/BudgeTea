package datamng

import (
	"time"
)

type Month struct {
	Budget        float64
	TotalSpending float64

	Year int
	Moon time.Month

	Groceries    []Expense
	Hobbies      []Expense
	Rent         []Expense
	OtherBills   []Expense
	Travel       []Expense
	Miscelanious []Expense
}

type Data struct {
	Budget       float64
	Denomination float64
	Months       []Month
}

func (data Data) FindMonthByYear(month time.Month, year int) int {
	for i, j := range data.Months {
		if j.Moon == month && j.Year == year {
			return i
		}
	}

	return -1
}
