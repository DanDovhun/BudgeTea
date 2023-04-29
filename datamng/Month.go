package datamng

import "time"

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

func (month *Month) SetBudget(budget float64) {
	month.Budget = budget
}
