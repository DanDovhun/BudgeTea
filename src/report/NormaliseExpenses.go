package report

import (
	"BudgeTea/datamng"
	"BudgeTea/forex"
)

func NormaliseExpenses(month datamng.Month, currency string) (datamng.Month, error) {
	conversionRates, err := forex.GetRates(currency)

	if err != nil {
		return datamng.Month{}, err
	}

	// Normalise all grocery expenses to be the same price
	for _, i := range month.Groceries {
		expense := i

		// If it's different currency than current user's preference
		if expense.Denomination != currency {
			// Convert it to the user's preferred currency
			newPrice := calculatePrice(currency, expense, conversionRates)

			// Set new price
			expense.SetPrice(newPrice)
		}
	}

	// Normalise all hobbies expenses to be the same price
	for _, i := range month.Hobbies {
		expense := i

		// If it's different currency than current user's preference
		if expense.Denomination != currency {
			// Convert it to the user's preferred currency
			newPrice := calculatePrice(currency, expense, conversionRates)

			// Set new price
			expense.SetPrice(newPrice)
		}
	}

	// Normalise all miscelanious expenses to be the same price
	for _, i := range month.Miscelanious {
		expense := i

		// If it's different currency than current user's preference
		if expense.Denomination != currency {
			// Convert it to the user's preferred currency
			newPrice := calculatePrice(currency, expense, conversionRates)

			// Set new price
			expense.SetPrice(newPrice)
		}
	}

	// Normalise all other expenses to be the same price
	for _, i := range month.OtherBills {
		expense := i

		// If it's different currency than current user's preference
		if expense.Denomination != currency {
			// Convert it to the user's preferred currency
			newPrice := calculatePrice(currency, expense, conversionRates)

			// Set new price
			expense.SetPrice(newPrice)
		}
	}

	// Normalise all rent expenses to be the same price
	for _, i := range month.Rent {
		expense := i

		// If it's different currency than current user's preference
		if expense.Denomination != currency {
			// Convert it to the user's preferred currency
			newPrice := calculatePrice(currency, expense, conversionRates)

			// Set new price
			expense.SetPrice(newPrice)
		}
	}

	// Normalise all travel expenses to be the same price
	for _, i := range month.Travel {
		expense := i

		// If it's different currency than current user's preference
		if expense.Denomination != currency {
			// Convert it to the user's preferred currency
			newPrice := calculatePrice(currency, expense, conversionRates)

			// Set new price
			expense.SetPrice(newPrice)
		}
	}

	return month, nil
}
