package datamng

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

// Write the new budget in the database
func SetBudget(budget float64) error {
	data, err := os.Open("expenses.json") // Open the json file
	var months Data                       // The loaded data will be stored here

	if err != nil {
		return err
	}

	defer data.Close() // close the file when the function ends

	byteValue, _ := ioutil.ReadAll(data) //Read the file as []bytes
	json.Unmarshal(byteValue, &months)   // Store the bytes in the months

	months.SetBudget(budget) // Set current budget

	currentYear, currentMonth, _ := time.Now().Date() // Get current year and month

	monthIndex := months.FindMonthByYear(currentMonth, currentYear) // Check where current month is and if it exists

	// If the month doesn't yet exist (the new expense is first that month)
	if monthIndex == -1 {
		// Add the month to the database
		months.Months = append(months.Months, NewMonth(months.Budget))

		// The index of that month will be the last in the database
		monthIndex = len(months.Months) - 1
	}

	months.Months[len(months.Months)-1].SetBudget(budget) // Set current month's budget

	j, _ := json.MarshalIndent(months, "", "    ")
	_ = ioutil.WriteFile("expenses.json", j, 0644)

	return nil
}

// Write the current currency in the database
func SetCurrency(currency string) error {
	data, err := os.Open("expenses.json") // Open the json file
	var months Data                       // The loaded data will be stored here

	if err != nil {
		return err
	}

	defer data.Close() // close the file when the function ends

	byteValue, _ := ioutil.ReadAll(data) //Read the file as []bytes
	json.Unmarshal(byteValue, &months)   // Store the bytes in the months

	oldCurrency, err := GetCurrency() // Get old currency

	if err != nil {
		return err
	}

	connectionError := months.Months[len(months.Months)-1].ChangeSpendingToCurrency(oldCurrency, currency)

	if connectionError != nil {
		return errors.New("Cannot connect to the forex API")
	}

	months.SetCurrency(currency)

	j, _ := json.MarshalIndent(months, "", "    ")
	_ = ioutil.WriteFile("expenses.json", j, 0644)

	return nil
}

// Fetches the preferred currency
func GetCurrency() (string, error) {
	var currency string
	data, err := os.Open("expenses.json") // Open the json file
	var months Data                       // The loaded data will be stored here

	if err != nil {
		return "", err
	}

	byteValue, _ := ioutil.ReadAll(data) //Read the file as []bytes
	json.Unmarshal(byteValue, &months)   // Store the bytes in the months

	currency = months.Denomination

	data.Close() // close the file

	return currency, nil
}
