package datamng

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

// Adds a new expense to the database
func Add(expense Expense) error {
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
		months.Months = append(months.Months, Month{
			Budget:        months.Budget, // Make that months budget equal to the current budget
			TotalSpending: 0,             // Set total spending to 0

			Year: currentYear,
			Moon: currentMonth,

			Groceries:    []Expense{},
			Hobbies:      []Expense{},
			Rent:         []Expense{},
			OtherBills:   []Expense{},
			Travel:       []Expense{},
			Miscelanious: []Expense{},
		})

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
