package datamng

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func SetBudget(budget float64) error {
	data, err := os.Open("expenses.json") // Open the json file
	var months Data                       // The loaded data will be stored here

	if err != nil {
		return err
	}

	defer data.Close() // close the file when the function ends

	byteValue, _ := ioutil.ReadAll(data) //Read the file as []bytes
	json.Unmarshal(byteValue, &months)   // Store the bytes in the months

	months.SetBudget(budget)                              // Set current budget
	months.Months[len(months.Months)-1].SetBudget(budget) // Set current month's budget

	j, _ := json.MarshalIndent(months, "", "    ")
	_ = ioutil.WriteFile("expenses.json", j, 0644)

	return nil
}
