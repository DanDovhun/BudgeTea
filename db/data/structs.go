package data

type Item struct {
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Year     int     `json:"year"`
	Month    string  `json:"month"`
	Day      int     `json:"day"`
}

type MonthlyExpense struct {
	Year              int     `json:"year"`
	Month             string  `json:"month"`
	TotalExpenses     float64 `json:"total_expenses"`
	GroceriesExpenses float64 `json:"groceries_expenses"`
	TravelExpenses    float64 `json:"travel_expenses"`
	OtherExpenses     float64 `json:"other_expenses"`
}

type Category struct {
	Price float64 `json:"total_price"`
	Items []Item  `json:"items"`
}

type Data struct {
	TotalPrice float64          `json:"total_price"`
	Groceries  Category         `json:"groceries"`
	Travel     Category         `json:"travel"`
	Others     Category         `json:"others"`
	Expenses   []MonthlyExpense `json:"monthly_expenses"`
}
