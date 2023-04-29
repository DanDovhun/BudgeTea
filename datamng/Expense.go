package datamng

import (
	"fmt"
	"time"
)

type Expense struct {
	Name         string
	Category     string
	Price        float64
	Denomination string
	Day          int
	Date         string
}

func NewExpense(title, category, denomination string, price float64) Expense {
	currentTime := time.Now()
	year, month, day := currentTime.Date()

	return Expense{
		Name:         title,
		Category:     category,
		Price:        price,
		Denomination: denomination,
		Day:          day,
		Date:         fmt.Sprintf("%v-%v-%v", day, month, year),
	}
}

func (ex Expense) FormatToCSV() string {
	return fmt.Sprintf("%v,%v,%v,%v,%v", ex.Name, ex.Category, ex.Price, ex.Denomination, ex.Date)
}
