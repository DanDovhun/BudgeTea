package datamng

import (
	"fmt"
	"time"
)

type Expense struct {
	Name         string
	Category     string
	Price        float64
	Denomination string
	Day          int
	Date         string
}

// Creates a new expense, it's essentially a constructor for the Expense struct
func NewExpense(title, category, denomination string, price float64) Expense {
	currentTime := time.Now()              // Get current time
	year, month, day := currentTime.Date() // Get current date

	// Setup and return a new expense
	return Expense{
		Name:         title,
		Category:     category,
		Price:        price,
		Denomination: denomination,
		Day:          day,
		Date:         fmt.Sprintf("%v-%v-%v", day, month, year),
	}
}

// Formats the Expense into a string so it could be added to a CSV file
func (ex Expense) FormatToCSV() string {
	return fmt.Sprintf("%v,%v,%v,%v,%v", ex.Name, ex.Category, ex.Price, ex.Denomination, ex.Date)
}
