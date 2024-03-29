package report

import (
	"BudgeTea/datamng"
	"BudgeTea/forex"
	"fmt"
)

// Holds spending report data
type ExpenseReport struct {
	Budget        float64
	TotalSpending float64
	Groceries     float64
	Hobbies       float64
	Rent          float64
	Travel        float64
	Miscelanious  float64
	OtherBills    float64
}

// Converts expense's currency to user's preferred currency
func convert(from string, amount, rate float64) float64 {
	return amount * rate
}

// Get expense report for current month
func GetThisMonthsExpenses() (ExpenseReport, error) {
	data, err := datamng.GetData() // Get months

	// If program cannot access db
	if err != nil {
		return ExpenseReport{}, err
	}

	//Get current month
	month := data.Months[len(data.Months)-1]

	// Get user's preferred currency
	currency, err := datamng.GetCurrency()

	// If cannot be accessed
	if err != nil {
		return ExpenseReport{}, nil
	}

	expenses := ExpenseReport{}
	expenses.Budget = data.Budget
	expenses.TotalSpending = month.TotalSpending

	conversionRates, err := forex.GetRates(data.Denomination)

	// Normalise all grocery expenses to be the same price
	for _, i := range month.Groceries {
		expense := i

		// If it's different currency than current user's preference
		if expense.Denomination != currency {
			// Convert it to the user's preferred currency
			var newPrice float64

			// If the currency is set to SEK and the price is in euros
			if expense.Denomination == "EUR" && currency == "SEK" {
				newPrice = convert("EUR", expense.Price, conversionRates.First)
			}

			// If the currency is set to SEK and the price is in us dollars
			if expense.Denomination == "USD" && currency == "SEK" {
				newPrice = convert("USD", expense.Price, conversionRates.Second)
			}

			// If the currency is set to EUR and the price is in swedish kronar
			if expense.Denomination == "SEK" && currency == "EUR" {
				newPrice = convert("SEK", expense.Price, conversionRates.First)
			}

			// If the currency is set to EUR and the price is in us dollars
			if expense.Denomination == "USD" && currency == "EUR" {
				newPrice = convert("USD", expense.Price, conversionRates.Second)
			}

			// If the currency is set to EUR and the price is in swedish kronar
			if expense.Denomination == "SEK" && currency == "USD" {
				newPrice = convert("SEK", expense.Price, conversionRates.Second)
			}

			// If the currency is set to EUR and the price is in us dollars
			if expense.Denomination == "EUR" && currency == "USD" {
				newPrice = convert("USD", expense.Price, conversionRates.First)
			}

			fmt.Println(newPrice)

			// Set new price
			expense.SetPrice(newPrice)
		}

		expenses.Groceries += expense.Price
	}

	// Normalise all hobbies expenses to be the same price
	for _, i := range month.Hobbies {
		expense := i

		// If it's different currency than current user's preference
		if expense.Denomination != currency {
			// Convert it to the user's preferred currency
			var newPrice float64

			// If the currency is set to SEK and the price is in euros
			if expense.Denomination == "EUR" && currency == "SEK" {
				newPrice = convert("EUR", expense.Price, conversionRates.First)
			}

			// If the currency is set to SEK and the price is in us dollars
			if expense.Denomination == "USD" && currency == "SEK" {
				newPrice = convert("USD", expense.Price, conversionRates.Second)
			}

			// If the currency is set to EUR and the price is in swedish kronar
			if expense.Denomination == "SEK" && currency == "EUR" {
				newPrice = convert("SEK", expense.Price, conversionRates.First)
			}

			// If the currency is set to EUR and the price is in us dollars
			if expense.Denomination == "USD" && currency == "EUR" {
				newPrice = convert("USD", expense.Price, conversionRates.Second)
			}

			// If the currency is set to EUR and the price is in swedish kronar
			if expense.Denomination == "SEK" && currency == "USD" {
				newPrice = convert("SEK", expense.Price, conversionRates.Second)
			}

			// If the currency is set to EUR and the price is in us dollars
			if expense.Denomination == "EUR" && currency == "USD" {
				newPrice = convert("USD", expense.Price, conversionRates.First)
			}

			// Set new price
			expense.SetPrice(newPrice)
		}

		expenses.Hobbies += expense.Price
	}

	// Normalise all miscelanious expenses to be the same price
	for _, i := range month.Miscelanious {
		expense := i

		// If it's different currency than current user's preference
		if expense.Denomination != currency {
			// Convert it to the user's preferred currency
			var newPrice float64

			// If the currency is set to SEK and the price is in euros
			if expense.Denomination == "EUR" && currency == "SEK" {
				newPrice = convert("EUR", expense.Price, conversionRates.First)
			}

			// If the currency is set to SEK and the price is in us dollars
			if expense.Denomination == "USD" && currency == "SEK" {
				newPrice = convert("USD", expense.Price, conversionRates.Second)
			}

			// If the currency is set to EUR and the price is in swedish kronar
			if expense.Denomination == "SEK" && currency == "EUR" {
				newPrice = convert("SEK", expense.Price, conversionRates.First)
			}

			// If the currency is set to EUR and the price is in us dollars
			if expense.Denomination == "USD" && currency == "EUR" {
				newPrice = convert("USD", expense.Price, conversionRates.Second)
			}

			// If the currency is set to EUR and the price is in swedish kronar
			if expense.Denomination == "SEK" && currency == "USD" {
				newPrice = convert("SEK", expense.Price, conversionRates.Second)
			}

			// If the currency is set to EUR and the price is in us dollars
			if expense.Denomination == "EUR" && currency == "USD" {
				newPrice = convert("USD", expense.Price, conversionRates.First)
			}

			// Set new price
			expense.SetPrice(newPrice)
		}

		expenses.Miscelanious += expense.Price
	}

	// Normalise all other expenses to be the same price
	for _, i := range month.OtherBills {
		expense := i

		// If it's different currency than current user's preference
		if expense.Denomination != currency {
			// Convert it to the user's preferred currency
			var newPrice float64

			// If the currency is set to SEK and the price is in euros
			if expense.Denomination == "EUR" && currency == "SEK" {
				newPrice = convert("EUR", expense.Price, conversionRates.First)
			}

			// If the currency is set to SEK and the price is in us dollars
			if expense.Denomination == "USD" && currency == "SEK" {
				newPrice = convert("USD", expense.Price, conversionRates.Second)
			}

			// If the currency is set to EUR and the price is in swedish kronar
			if expense.Denomination == "SEK" && currency == "EUR" {
				newPrice = convert("SEK", expense.Price, conversionRates.First)
			}

			// If the currency is set to EUR and the price is in us dollars
			if expense.Denomination == "USD" && currency == "EUR" {
				newPrice = convert("USD", expense.Price, conversionRates.Second)
			}

			// If the currency is set to EUR and the price is in swedish kronar
			if expense.Denomination == "SEK" && currency == "USD" {
				newPrice = convert("SEK", expense.Price, conversionRates.Second)
			}

			// If the currency is set to EUR and the price is in us dollars
			if expense.Denomination == "EUR" && currency == "USD" {
				newPrice = convert("USD", expense.Price, conversionRates.First)
			}

			// Set new price
			expense.SetPrice(newPrice)
		}

		expenses.OtherBills += expense.Price
	}

	// Normalise all rent expenses to be the same price
	for _, i := range month.Rent {
		expense := i

		// If it's different currency than current user's preference
		if expense.Denomination != currency {
			// Convert it to the user's preferred currency
			var newPrice float64

			// If the currency is set to SEK and the price is in euros
			if expense.Denomination == "EUR" && currency == "SEK" {
				newPrice = convert("EUR", expense.Price, conversionRates.First)
			}

			// If the currency is set to SEK and the price is in us dollars
			if expense.Denomination == "USD" && currency == "SEK" {
				newPrice = convert("USD", expense.Price, conversionRates.Second)
			}

			// If the currency is set to EUR and the price is in swedish kronar
			if expense.Denomination == "SEK" && currency == "EUR" {
				newPrice = convert("SEK", expense.Price, conversionRates.First)
			}

			// If the currency is set to EUR and the price is in us dollars
			if expense.Denomination == "USD" && currency == "EUR" {
				newPrice = convert("USD", expense.Price, conversionRates.Second)
			}

			// If the currency is set to EUR and the price is in swedish kronar
			if expense.Denomination == "SEK" && currency == "USD" {
				newPrice = convert("SEK", expense.Price, conversionRates.Second)
			}

			// If the currency is set to EUR and the price is in us dollars
			if expense.Denomination == "EUR" && currency == "USD" {
				newPrice = convert("USD", expense.Price, conversionRates.First)
			}

			// Set new price
			expense.SetPrice(newPrice)
		}

		expenses.Rent += expense.Price
	}

	// Normalise all travel expenses to be the same price
	for _, i := range month.Travel {
		expense := i

		// If it's different currency than current user's preference
		if expense.Denomination != currency {
			// Convert it to the user's preferred currency
			var newPrice float64

			// If the currency is set to SEK and the price is in euros
			if expense.Denomination == "EUR" && currency == "SEK" {
				newPrice = convert("EUR", expense.Price, conversionRates.First)
			}

			// If the currency is set to SEK and the price is in us dollars
			if expense.Denomination == "USD" && currency == "SEK" {
				newPrice = convert("USD", expense.Price, conversionRates.Second)
			}

			// If the currency is set to EUR and the price is in swedish kronar
			if expense.Denomination == "SEK" && currency == "EUR" {
				newPrice = convert("SEK", expense.Price, conversionRates.First)
			}

			// If the currency is set to EUR and the price is in us dollars
			if expense.Denomination == "USD" && currency == "EUR" {
				newPrice = convert("USD", expense.Price, conversionRates.Second)
			}

			// If the currency is set to EUR and the price is in swedish kronar
			if expense.Denomination == "SEK" && currency == "USD" {
				newPrice = convert("SEK", expense.Price, conversionRates.Second)
			}

			// If the currency is set to EUR and the price is in us dollars
			if expense.Denomination == "EUR" && currency == "USD" {
				newPrice = convert("USD", expense.Price, conversionRates.First)
			}

			// Set new price
			expense.SetPrice(newPrice)
		}

		expenses.Travel += expense.Price
	}

	return expenses, nil
}
