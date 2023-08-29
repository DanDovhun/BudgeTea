package datamng

import (
	"BudgeTea/forex"
	"time"
)

// Stores information about month's spending
type Month struct {
	// Months budget and
	Budget        float64 `json:"budget"`
	TotalSpending float64 `json:"total_spending"`

	// Stores the month's name and year
	Year int        `json:"year"`
	Moon time.Month `json:"month"`

	// Stores month's expenses in specific categories
	Groceries    []Expense `json:"groceries"`
	Hobbies      []Expense `json:"hobbies"`
	Rent         []Expense `json:"rent"`
	OtherBills   []Expense `json:"other_bills"`
	Travel       []Expense `json:"travel"`
	Miscelanious []Expense `json:"miscalenious"`
}

// Month Constructor
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

// When changing prefered currency, the spending will be changed to the value of the new one
func (month *Month) ChangeSpendingToCurrency(old, new string) error {
	spending, err := forex.Convert(old, new, month.TotalSpending) // Convert spending to the new currency

	if err != nil { // If error happens
		return err // ret7urn error
	}

	month.TotalSpending = spending // Set spending to the new value

	return nil // Return null error
}

// Gets Budget
func (month Month) GetBudget() float64 {
	return month.Budget
}

// Will get total spending
func (month Month) GetTotalSpending() float64 {
	return month.TotalSpending
}
