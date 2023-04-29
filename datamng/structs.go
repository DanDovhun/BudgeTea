package datamng

import (
	"time"
)

// Month struct
type Month struct {
	// Months budget and
	Budget        float64
	TotalSpending float64

	// Stores the month's name and year
	Year int
	Moon time.Month

	// Stores month's expenses in specific categories
	Groceries    []Expense
	Hobbies      []Expense
	Rent         []Expense
	OtherBills   []Expense
	Travel       []Expense
	Miscelanious []Expense
}

// Data struct
// The database is loaded into this struct
type Data struct {
	Budget       float64
	Denomination float64
	Months       []Month
}

// Searches for a specific month
func (data Data) FindMonthByYear(month time.Month, year int) int {
	// Iterate through months
	for i, j := range data.Months {
		// If a month and and its year match
		if j.Moon == month && j.Year == year {
			return i // Return its index
		}
	}

	return -1 // If not found, return -1
}

func (month *Month) SetBudget(budget float64) {
	month.Budget = budget
}

func (data *Data) SetBudget(budget float64) {
	data.Budget = budget
}
