package datamng

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Stores information about an expense
type Expense struct {
	Name         string
	Category     string
	Price        float64
	Denomination string
	Day          int
	Date         string
}

// Creates a new expense, it's essentially a constructor for the Expense struct
func NewExpense(title, category string, price float64) (Expense, error) {
	currentTime := time.Now()              // Get current time
	year, month, day := currentTime.Date() // Get current date

	data, err := os.Open("expenses.json") // Open the json file
	var months Data                       // The loaded data will be stored here

	// If the file couldn't be opened
	if err != nil {
		return Expense{}, err // Return the error
	}

	defer data.Close() // close the file when the function ends

	byteValue, _ := ioutil.ReadAll(data) //Read the file as []bytes
	json.Unmarshal(byteValue, &months)   // Store the bytes in the months

	// Setup and return a new expense
	return Expense{
		Name:         title,
		Category:     category,
		Price:        price,
		Denomination: months.Denomination,
		Day:          day,
		Date:         fmt.Sprintf("%v-%v-%v", day, month, year),
	}, nil
}

// Formats the Expense into a string so it could be added to a CSV file
func (ex Expense) FormatToCSV() string {
	return fmt.Sprintf("%v,%v,%v,%v,%v", ex.Name, ex.Category, ex.Price, ex.Denomination, ex.Date)
}

// Adds a new expense to the database
func (ex Expense) Add(expense Expense) error {
	data, err := os.Open("expenses.json") // Open the json file
	var months Data                       // The loaded data will be stored here

	// If the file couldn't be opened
	if err != nil {
		return err // Return the error
	}

	defer data.Close() // close the file when the function ends

	byteValue, _ := ioutil.ReadAll(data) //Read the file as []bytes
	json.Unmarshal(byteValue, &months)   // Store the bytes in the months

	currentYear, currentMonth, _ := time.Now().Date() // Get current year and month

	monthIndex := months.FindMonthByYear(currentMonth, currentYear) // Check where current month is and if it exists

	// If the month doesn't yet exist (the new expense is first that month)
	if monthIndex == -1 {
		// Add the month to the database
		months.Months = append(months.Months, NewMonth(months.Budget))

		// The index of that month will be the last in the database
		monthIndex = len(months.Months) - 1
	}

	//Increment that month's spending by the expense's cost
	months.Months[monthIndex].TotalSpending += expense.Price

	// Sort the expense into its category
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

	// Format the update database and write it to the json file
	j, _ := json.MarshalIndent(months, "", "    ")
	_ = ioutil.WriteFile("expenses.json", j, 0644)

	// Do not return errors
	return nil
}

// Sets the cost of the expense
func (ex *Expense) SetPrice(price float64) {
	ex.Price = price
}
