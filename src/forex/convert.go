package forex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Value struct {
	Euro    float64 `json:"EUR"`
	Dollars float64 `json:"USD"`
	Kronar  float64 `json:"SEK"`
}

type Response struct {
	Data Value `json:"data"`
}

const APi_KEY string = "g6cMYaCD3g8QIpqp9z8pQK5piOd82ROCY16uSU7P"

func Convert(from, to string, amount float64) (float64, error) {
	result, err := http.Get(fmt.Sprintf("https://api.freecurrencyapi.com/v1/latest?apikey=%v&currencies=%v&base_currency=%v", APi_KEY, to, from))

	if err != nil {
		return 0, err
	}

	response, err := ioutil.ReadAll(result.Body)

	if err != nil {
		return 0, err
	}

	var value Response

	json.Unmarshal(response, &value)

	if from == "EUR" {
		return value.Data.Euro * amount, nil
	}

	if from == "USD" {
		return value.Data.Dollars * amount, nil
	}

	if from == "SEK" {
		return value.Data.Kronar * amount, nil
	}

	return 0, nil
}
