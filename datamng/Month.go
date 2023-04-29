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

func NewMonth(budget float64) Month {
	currentYear, currentMonth, _ := time.Now().Date() // Get current year and month

	return Month{
		Budget:        budget, // Make that months budget equal to the current budget
		TotalSpending: 0,      // Set total spending to 0

		Year: currentYear,
		Moon: currentMonth,

		Groceries:    []Expense{},
		Hobbies:      []Expense{},
		Rent:         []Expense{},
		OtherBills:   []Expense{},
		Travel:       []Expense{},
		Miscelanious: []Expense{},
	}
}

// Sets month's budget
func (month *Month) SetBudget(budget float64) {
	month.Budget = budget
}
