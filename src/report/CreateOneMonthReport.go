package report

import (
	"BudgeTea/datamng"
	"fmt"

	"fyne.io/fyne/v2"
)

// Formats the month object into a string that can be exported into a csv
func CreateOneMonthReport(month datamng.Month, currency string) string {
	var spending string

	var (
		groceries string
		rent      string
		others    string
		travel    string
		hobbies   string
		misc      string

		grocerySpending float64
		rentSpending    float64
		otherSpending   float64
		travelSpending  float64
		hobbySpending   float64
		miscSpending    float64
	)

	// Report on budget and total spending
	spending += fmt.Sprintf("Budget:,%v %v\n", month.Budget, currency)
	spending += fmt.Sprintf("Total spending:,%v %v\n", month.TotalSpending, currency)

	if month.TotalSpending < month.Budget { // If the spending is less than budget
		under := month.Budget - month.TotalSpending
		percentage := (under / month.Budget) * 100
		spending += fmt.Sprintf("Under budget:,%v %v,%v", under, currency, percentage) + " %\n\n"
	} else { // If over budget
		over := month.TotalSpending - month.Budget
		percentage := (over / month.Budget) * 100
		spending += fmt.Sprintf("Over budget:,%v %v,%v", over, currency, percentage) + " %\n\n"
	}

	// Collect groceries expenses
	groceries = "Title,Cost,Date\n"
	for _, i := range month.Groceries {
		groceries += fmt.Sprintf("%v,%v %v, %v\n", i.Name, i.Price, i.Denomination, i.Date)
		grocerySpending += i.Price
	}

	// Add groceries
	spending += fmt.Sprintf("Groceries:,%v %v\n", grocerySpending, currency) + groceries + "\n\n"

	// Collect rent expenses
	rent = "Title,Cost,Date\n"
	for _, i := range month.Rent {
		rent += fmt.Sprintf("%v,%v %v, %v\n", i.Name, i.Price, i.Denomination, i.Date)
		rentSpending += i.Price
	}

	// Add rent
	spending += fmt.Sprintf("Rent:,%v %v\n", rentSpending, currency) + rent + "\n\n"

	// Collect other bills expenses
	others = "Title,Cost,Date\n"
	for _, i := range month.OtherBills {
		others += fmt.Sprintf("%v,%v %v, %v\n", i.Name, i.Price, i.Denomination, i.Date)
		otherSpending += i.Price
	}

	// Add other bills
	spending += fmt.Sprintf("Other bills:,%v %v\n", otherSpending, currency) + others + "\n\n"

	// Adds hobbies
	hobbies = "Title,Cost,Date\n"
	for _, i := range month.Hobbies {
		hobbies += fmt.Sprintf("%v,%v %v, %v\n", i.Name, i.Price, i.Denomination, i.Date)
		hobbySpending += i.Price
	}

	// Add other bills
	spending += fmt.Sprintf("Hobbies:,%v %v\n", hobbySpending, currency) + hobbies + "\n\n"

	// Adds travel
	hobbies = "Title,Cost,Date\n"
	for _, i := range month.Travel {
		travel += fmt.Sprintf("%v,%v %v, %v\n", i.Name, i.Price, i.Denomination, i.Date)
		travelSpending += i.Price
	}

	// Add other bills
	spending += fmt.Sprintf("Travel:,%v %v\n", travelSpending, currency) + hobbies + "\n\n"

	// Adds travel
	misc = "Title,Cost,Date\n"
	for _, i := range month.Miscelanious {
		misc += fmt.Sprintf("%v,%v %v, %v\n", i.Name, i.Price, i.Denomination, i.Date)
		miscSpending += i.Price
	}

	// Add other bills
	spending += fmt.Sprintf("Miscelanious:,%v %v\n", travelSpending, currency) + hobbies + "\n\n"

	return spending
}

func CreateLastMonthsReport(file fyne.URIWriteCloser, fileName string) error {
	data, err := datamng.GetData()

	if err != nil {
		return err
	}

	month := data.GetLastMonth()

	normal, err := NormaliseExpenses(month, data.Denomination)

	if err != nil {
		return err
	}

	list := CreateOneMonthReport(normal, data.Denomination)

	file.Write([]byte(list))

	return nil
}
