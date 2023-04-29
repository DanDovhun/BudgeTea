package datamng

import (
	"time"
)

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

// Sets current budget
func (data *Data) SetBudget(budget float64) {
	data.Budget = budget
}
