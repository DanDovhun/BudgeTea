package forex

import (
	"encoding/json"
	"errors"
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

type Rates struct {
	First  float64
	Second float64
}

const API_KEY string = "g6cMYaCD3g8QIpqp9z8pQK5piOd82ROCY16uSU7P"

func GetRates(to string) (Rates, error) {
	switch to {
	case "SEK":
		euro, err := Convert("EUR", to, 1)

		if err != nil {
			return Rates{}, err
		}

		dollars, err := Convert("USD", to, 1)

		return Rates{
			First:  euro,
			Second: dollars,
		}, nil

	case "EUR":
		sek, err := Convert("SEK", to, 1)

		if err != nil {
			return Rates{}, err
		}

		dollars, err := Convert("USD", to, 1)

		return Rates{
			First:  sek,
			Second: dollars,
		}, nil

	case "USD":
		euro, err := Convert("EUR", to, 1)

		if err != nil {
			return Rates{}, err
		}

		sek, err := Convert("SEK", to, 1)

		return Rates{
			First:  euro,
			Second: sek,
		}, nil

	default:
		return Rates{}, errors.New("Invalid currency")
	}
}

func Convert(from, to string, amount float64) (float64, error) {
	result, err := http.Get(fmt.Sprintf("https://api.freecurrencyapi.com/v1/latest?apikey=%v&currencies=%v&base_currency=%v", API_KEY, to, from))

	if err != nil {
		return 0, err
	}

	response, err := ioutil.ReadAll(result.Body)

	if err != nil {
		return 0, err
	}

	var value Response
	var conversion float64

	json.Unmarshal(response, &value)

	fmt.Println(err)

	if to == "EUR" {
		conversion = value.Data.Euro * amount
	}

	if to == "USD" {
		conversion = value.Data.Dollars * amount
	}

	if to == "SEK" {
		conversion = value.Data.Kronar * amount
	}

	return conversion, nil
}
