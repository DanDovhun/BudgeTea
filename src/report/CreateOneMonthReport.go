package report

import (
	"BudgeTea/datamng"
	"fmt"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"github.com/xuri/excelize/v2"
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
	spending += fmt.Sprintf("Budget:;%v %v\n", month.Budget, currency)
	spending += fmt.Sprintf("Total spending:;%v %v\n", month.TotalSpending, currency)

	if month.TotalSpending < month.Budget { // If the spending is less than budget
		under := month.Budget - month.TotalSpending
		percentage := (under / month.Budget) * 100
		spending += fmt.Sprintf("Under budget:;%v %v;%v", under, currency, percentage) + " %\n\n"
	} else { // If over budget
		over := month.TotalSpending - month.Budget
		percentage := (over / month.Budget) * 100
		spending += fmt.Sprintf("Over budget:;%v %v;%v", over, currency, percentage) + " %\n\n"
	}

	// Collect groceries expenses
	groceries = "Title;Cost;Date\n"
	for _, i := range month.Groceries {
		groceries += fmt.Sprintf("%v;%v %v; %v\n", i.Name, i.Price, i.Denomination, i.Date)
		grocerySpending += i.Price
	}

	// Add groceries
	spending += fmt.Sprintf("Groceries:;%v %v\n", grocerySpending, currency) + groceries + "\n\n"

	// Collect rent expenses
	rent = "Title;Cost;Date\n"
	for _, i := range month.Rent {
		rent += fmt.Sprintf("%v;%v %v; %v\n", i.Name, i.Price, i.Denomination, i.Date)
		rentSpending += i.Price
	}

	// Add rent
	spending += fmt.Sprintf("Rent:;%v %v\n", rentSpending, currency) + rent + "\n\n"

	// Collect other bills expenses
	others = "Title;Cost;Date\n"
	for _, i := range month.OtherBills {
		others += fmt.Sprintf("%v;%v %v; %v\n", i.Name, i.Price, i.Denomination, i.Date)
		otherSpending += i.Price
	}

	// Add other bills
	spending += fmt.Sprintf("Other bills:;%v %v\n", otherSpending, currency) + others + "\n\n"

	// Adds hobbies
	hobbies = "Title;Cost;Date\n"
	for _, i := range month.Hobbies {
		hobbies += fmt.Sprintf("%v;%v %v; %v\n", i.Name, i.Price, i.Denomination, i.Date)
		hobbySpending += i.Price
	}

	// Add other bills
	spending += fmt.Sprintf("Hobbies:;%v %v\n", hobbySpending, currency) + hobbies + "\n\n"

	// Adds travel
	hobbies = "Title;Cost;Date\n"
	for _, i := range month.Travel {
		travel += fmt.Sprintf("%v;%v %v; %v\n", i.Name, i.Price, i.Denomination, i.Date)
		travelSpending += i.Price
	}

	// Add other bills
	spending += fmt.Sprintf("Travel:;%v %v\n", travelSpending, currency) + hobbies + "\n\n"

	// Adds travel
	misc = "Title;Cost;Date\n"
	for _, i := range month.Miscelanious {
		misc += fmt.Sprintf("%v;%v %v; %v\n", i.Name, i.Price, i.Denomination, i.Date)
		miscSpending += i.Price
	}

	// Add other bills
	spending += fmt.Sprintf("Miscelanious:;%v %v\n", travelSpending, currency) + hobbies + "\n\n"

	return spending
}

func CreateLastMonthExcel(file fyne.URIWriteCloser, data datamng.Month, budget float64, currency, name string) error {
	excelFile, err := excelize.OpenFile(fmt.Sprintf("%v", file.URI()))

	if err != nil {
		return err
	}

	excelFile.SetCellValue("Sheet1", "A1", "Budget:")
	excelFile.SetCellValue("Sheet1", "B1", fmt.Sprintf("%v", budget))

	excelFile.SetCellValue("Sheet1", "A2", "Total Spending")
	excelFile.SetCellValue("Sheet1", "B2", fmt.Sprintf("%v %v", data.TotalSpending, currency))

	if data.TotalSpending < data.Budget {
		under := data.Budget - data.TotalSpending
		percentage := (under / data.Budget) * 100

		excelFile.SetCellValue("Sheet1", "A3", "Under budget:")
		excelFile.SetCellValue("Sheet1", "B3", fmt.Sprintf("%v %v", under, currency))
		excelFile.SetCellValue("Sheet1", "C3", fmt.Sprintf("%v", percentage)+"%")
	} else {
		over := data.TotalSpending - data.Budget
		percentage := (over / data.Budget) * 100

		excelFile.SetCellValue("Sheet1", "A3", "Over budget:")
		excelFile.SetCellValue("Sheet1", "B3", fmt.Sprintf("%v %v", over, currency))
		excelFile.SetCellValue("Sheet1", "C3", fmt.Sprintf("%v", percentage)+"%")
	}

	excelFile.SetCellValue("Sheet1", "A6", "Title")
	excelFile.SetCellValue("Sheet1", "B6", "Cost")
	excelFile.SetCellValue("Sheet1", "C6", "Date")

	var spending float64
	groceriesRow := 5
	rows := 7

	for _, i := range data.Groceries {
		excelFile.SetCellValue("Sheet1", fmt.Sprintf("A%v", rows), i.Name)
		excelFile.SetCellValue("Sheet1", fmt.Sprintf("B%v", rows), fmt.Sprintf("%v %v", i.Price, i.Denomination))
		excelFile.SetCellValue("Sheet1", fmt.Sprintf("C%v", rows), i.Date)

		spending += i.Price
		rows++
	}

	excelFile.SetCellValue("Sheet1", fmt.Sprintf("A%v", groceriesRow), "Groceries:")
	excelFile.SetCellValue("Sheet1", fmt.Sprintf("B%v", groceriesRow), fmt.Sprintf("%v %v", spending, currency))

	spending = 0
	rentRow := rows + 2
	rentTitleRow := rentRow + 1
	rows += 4

	excelFile.SetCellValue("Sheet1", fmt.Sprintf("A%v", rentTitleRow), "Title")
	excelFile.SetCellValue("Sheet1", fmt.Sprintf("B%v", rentTitleRow), "Cost")
	excelFile.SetCellValue("Sheet1", fmt.Sprintf("C%v", rentTitleRow), "Date")

	for _, i := range data.Rent {
		excelFile.SetCellValue("Sheet1", fmt.Sprintf("A%v", rows), i.Name)
		excelFile.SetCellValue("Sheet1", fmt.Sprintf("B%v", rows), fmt.Sprintf("%v %v", i.Price, i.Denomination))
		excelFile.SetCellValue("Sheet1", fmt.Sprintf("C%v", rows), i.Date)

		spending += i.Price
		rows++
	}

	excelFile.SetCellValue("Sheet1", fmt.Sprintf("A%v", rentRow), "Rent:")
	excelFile.SetCellValue("Sheet1", fmt.Sprintf("B%v", rentRow), fmt.Sprintf("%v %v", spending, currency))

	spending = 0
	billsRow := rows + 2
	billsTitleRow := billsRow + 1
	rows += 4

	excelFile.SetCellValue("Sheet1", fmt.Sprintf("A%v", billsTitleRow), "Title")
	excelFile.SetCellValue("Sheet1", fmt.Sprintf("B%v", billsTitleRow), "Cost")
	excelFile.SetCellValue("Sheet1", fmt.Sprintf("C%v", billsTitleRow), "Date")

	for _, i := range data.OtherBills {
		excelFile.SetCellValue("Sheet1", fmt.Sprintf("A%v", rows), i.Name)
		excelFile.SetCellValue("Sheet1", fmt.Sprintf("B%v", rows), fmt.Sprintf("%v %v", i.Price, i.Denomination))
		excelFile.SetCellValue("Sheet1", fmt.Sprintf("C%v", rows), i.Date)

		spending += i.Price
		rows++
	}

	excelFile.SetCellValue("Sheet1", fmt.Sprintf("A%v", billsRow), "Other Bills:")
	excelFile.SetCellValue("Sheet1", fmt.Sprintf("B%v", billsRow), fmt.Sprintf("%v %v", spending, currency))

	spending = 0
	hobbiesRow := rows + 2
	hobbiesTitleRow := hobbiesRow + 1
	rows += 4

	excelFile.SetCellValue("Sheet1", fmt.Sprintf("A%v", hobbiesTitleRow), "Title")
	excelFile.SetCellValue("Sheet1", fmt.Sprintf("B%v", hobbiesTitleRow), "Cost")
	excelFile.SetCellValue("Sheet1", fmt.Sprintf("C%v", hobbiesTitleRow), "Date")

	for _, i := range data.Hobbies {
		excelFile.SetCellValue("Sheet1", fmt.Sprintf("A%v", rows), i.Name)
		excelFile.SetCellValue("Sheet1", fmt.Sprintf("B%v", rows), fmt.Sprintf("%v %v", i.Price, i.Denomination))
		excelFile.SetCellValue("Sheet1", fmt.Sprintf("C%v", rows), i.Date)

		spending += i.Price
		rows++
	}

	excelFile.SetCellValue("Sheet1", fmt.Sprintf("A%v", hobbiesRow), "Hobbies:")
	excelFile.SetCellValue("Sheet1", fmt.Sprintf("B%v", hobbiesRow), fmt.Sprintf("%v %v", spending, currency))

	spending = 0
	travelRows := rows + 2
	travelTitleRows := travelRows + 1
	rows += 4

	excelFile.SetCellValue("Sheet1", fmt.Sprintf("A%v", travelTitleRows), "Title")
	excelFile.SetCellValue("Sheet1", fmt.Sprintf("B%v", travelTitleRows), "Cost")
	excelFile.SetCellValue("Sheet1", fmt.Sprintf("C%v", travelTitleRows), "Date")

	for _, i := range data.Travel {
		excelFile.SetCellValue("Sheet1", fmt.Sprintf("A%v", rows), i.Name)
		excelFile.SetCellValue("Sheet1", fmt.Sprintf("B%v", rows), fmt.Sprintf("%v %v", i.Price, i.Denomination))
		excelFile.SetCellValue("Sheet1", fmt.Sprintf("C%v", rows), i.Date)

		spending += i.Price
		rows++
	}

	excelFile.SetCellValue("Sheet1", fmt.Sprintf("A%v", travelRows), "Travel:")
	excelFile.SetCellValue("Sheet1", fmt.Sprintf("B%v", travelRows), fmt.Sprintf("%v %v", spending, currency))

	spending = 0
	miscRows := rows + 2
	miscTitleRows := miscRows + 1
	rows += 4

	excelFile.SetCellValue("Sheet1", fmt.Sprintf("A%v", miscTitleRows), "Title")
	excelFile.SetCellValue("Sheet1", fmt.Sprintf("B%v", miscTitleRows), "Cost")
	excelFile.SetCellValue("Sheet1", fmt.Sprintf("C%v", miscTitleRows), "Date")

	for _, i := range data.Miscelanious {
		excelFile.SetCellValue("Sheet1", fmt.Sprintf("A%v", rows), i.Name)
		excelFile.SetCellValue("Sheet1", fmt.Sprintf("B%v", rows), fmt.Sprintf("%v %v", i.Price, i.Denomination))
		excelFile.SetCellValue("Sheet1", fmt.Sprintf("C%v", rows), i.Date)

		spending += i.Price
		rows++
	}

	excelFile.SetCellValue("Sheet1", fmt.Sprintf("A%v", miscRows), "Miscelanious:")
	excelFile.SetCellValue("Sheet1", fmt.Sprintf("B%v", miscRows), fmt.Sprintf("%v %v", spending, currency))

	err = excelFile.Save()

	if err != nil {

		return err
	}

	//move := fmt.Sprintf("%v", file.URI())

	//err = os.Rename(name, move)
	os.Create(file.URI().Fragment())

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
