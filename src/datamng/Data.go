package datamng

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

// Data struct
// The database is loaded into this struct
type Data struct {
	Budget       float64 `json:"budget"`
	Denomination string  `json:"denomination"`
	Months       []Month `json:"months"`
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

// Sets preferred currency
func (data *Data) SetCurrency(currency string) {
	data.Denomination = currency
}

// Gets budget from the db
func GetBudget() (float64, error) {
	data, err := os.Open("expenses.json") // Open the json file
	var months Data                       // The loaded data will be stored here

	// If the file couldn't be opened
	if err != nil {
		return 0, err // Return the error
	}

	byteValue, _ := ioutil.ReadAll(data) //Read the file as []bytes
	json.Unmarshal(byteValue, &months)   // Store the bytes in the months

	data.Close()

	return months.Budget, nil
}
