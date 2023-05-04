package report

import (
	"BudgeTea/datamng"
	"BudgeTea/forex"
	"fmt"
)

func (er ExpenseReport) GetBudgetAndSpending() string {
	var result string

	budget := fmt.Sprintf("Budget:,%v", er.Budget)
	spending := fmt.Sprintf("Spending:,%v", er.TotalSpending)

	result = fmt.Sprintf("%v\n%v\n", budget, spending)

	return result
}

func (er ExpenseReport) getGroceries() (string, error) {
	var report, items string

	data, err := datamng.GetData()

	if err != nil {
		return "", err
	}

	currency, err := datamng.GetCurrency()

	if err != nil {
		return "", err
	}

	rates, err := forex.GetRates(data.Denomination)

	if err != nil {
		return "", err
	}

	month := data.GetLastMonth().Groceries
	spending := fmt.Sprintf("Groceries:,%v", er.Groceries)

	items = "Name,Price,Category,Date Added\n"

	for _, i := range month {
		var price float64

		if i.Denomination != data.Denomination {
			price = calculatePrice(currency, i, rates)
		} else {
			price = i.Price
		}

		items += fmt.Sprintf("%v,%v %v,%v,%v\n", i.Name, price, currency, i.Category, i.Date)
	}

	report += fmt.Sprintf("%v\n%v\n", spending, items)

	return report, nil
}

func (er ExpenseReport) getRent() (string, error) {
	var report, items string

	data, err := datamng.GetData()

	if err != nil {
		return "", err
	}

	currency, err := datamng.GetCurrency()

	if err != nil {
		return "", err
	}

	rates, err := forex.GetRates(data.Denomination)

	if err != nil {
		return "", err
	}

	month := data.GetLastMonth().Rent
	spending := fmt.Sprintf("Rent:,%v", er.Rent)

	items = "Name,Price,Category,Date Added\n"

	for _, i := range month {
		var price float64

		if i.Denomination != data.Denomination {
			price = calculatePrice(currency, i, rates)
		} else {
			price = i.Price
		}

		items += fmt.Sprintf("%v,%v %v,%v,%v\n", i.Name, price, currency, i.Category, i.Date)
	}

	report += fmt.Sprintf("%v\n%v\n", spending, items)

	return report, nil
}

func (er ExpenseReport) getRent() (string, error) {
	var report, items string

	data, err := datamng.GetData()

	if err != nil {
		return "", err
	}

	currency, err := datamng.GetCurrency()

	if err != nil {
		return "", err
	}

	rates, err := forex.GetRates(data.Denomination)

	if err != nil {
		return "", err
	}

	month := data.GetLastMonth().Rent
	spending := fmt.Sprintf("Rent:,%v", er.Rent)

	items = "Name,Price,Category,Date Added\n"

	for _, i := range month {
		var price float64

		if i.Denomination != data.Denomination {
			price = calculatePrice(currency, i, rates)
		} else {
			price = i.Price
		}

		items += fmt.Sprintf("%v,%v %v,%v,%v\n", i.Name, price, currency, i.Category, i.Date)
	}

	report += fmt.Sprintf("%v\n%v\n", spending, items)

	return report, nil
}

func (er ExpenseReport) getRent() (string, error) {
	var report, items string

	data, err := datamng.GetData()

	if err != nil {
		return "", err
	}

	currency, err := datamng.GetCurrency()

	if err != nil {
		return "", err
	}

	rates, err := forex.GetRates(data.Denomination)

	if err != nil {
		return "", err
	}

	month := data.GetLastMonth().Rent
	spending := fmt.Sprintf("Rent:,%v", er.Rent)

	items = "Name,Price,Category,Date Added\n"

	for _, i := range month {
		var price float64

		if i.Denomination != data.Denomination {
			price = calculatePrice(currency, i, rates)
		} else {
			price = i.Price
		}

		items += fmt.Sprintf("%v,%v %v,%v,%v\n", i.Name, price, currency, i.Category, i.Date)
	}

	report += fmt.Sprintf("%v\n%v\n", spending, items)

	return report, nil
}

func (er ExpenseReport) getRent() (string, error) {
	var report, items string

	data, err := datamng.GetData()

	if err != nil {
		return "", err
	}

	currency, err := datamng.GetCurrency()

	if err != nil {
		return "", err
	}

	rates, err := forex.GetRates(data.Denomination)

	if err != nil {
		return "", err
	}

	month := data.GetLastMonth().Rent
	spending := fmt.Sprintf("Rent:,%v", er.Rent)

	items = "Name,Price,Category,Date Added\n"

	for _, i := range month {
		var price float64

		if i.Denomination != data.Denomination {
			price = calculatePrice(currency, i, rates)
		} else {
			price = i.Price
		}

		items += fmt.Sprintf("%v,%v %v,%v,%v\n", i.Name, price, currency, i.Category, i.Date)
	}

	report += fmt.Sprintf("%v\n%v\n", spending, items)

	return report, nil
}

func (er ExpenseReport) getRent() (string, error) {
	var report, items string

	data, err := datamng.GetData()

	if err != nil {
		return "", err
	}

	currency, err := datamng.GetCurrency()

	if err != nil {
		return "", err
	}

	rates, err := forex.GetRates(data.Denomination)

	if err != nil {
		return "", err
	}

	month := data.GetLastMonth().Rent
	spending := fmt.Sprintf("Rent:,%v", er.Rent)

	items = "Name,Price,Category,Date Added\n"

	for _, i := range month {
		var price float64

		if i.Denomination != data.Denomination {
			price = calculatePrice(currency, i, rates)
		} else {
			price = i.Price
		}

		items += fmt.Sprintf("%v,%v %v,%v,%v\n", i.Name, price, currency, i.Category, i.Date)
	}

	report += fmt.Sprintf("%v\n%v\n", spending, items)

	return report, nil
}
