package forex

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Value struct {
	From   string  `json:"base"`
	To     string  `json:"to"`
	OrgAmt float64 `json:"amount"`
	CvtAmt float64 `json:"converted"`
	Rate   float64 `json:"rate"`
}

type Rates struct {
	First  float64
	Second float64
}

const API_KEY string = "k90svvlpccoa0bdn1e01ebh6d6n1nfgorbb5fb8oma8pk3gh5prjoo"

// Gets conversion rate between two currencies
func getRate(from, to string) (float64, error) {
	result, err := http.Get(fmt.Sprintf("https://anyapi.io/api/v1/exchange/convert?apiKey=%v&base=%v&to=%v&amount=%v", API_KEY, from, to, 1))

	if err != nil {
		return 0, err
	}

	response, err := ioutil.ReadAll(result.Body)

	if err != nil {
		return 0, err
	}

	var value Value
	var conversionRate float64

	json.Unmarshal(response, &value)

	conversionRate = value.Rate

	return conversionRate, nil
}

// Get rates between preferred currrency and other supported currencies
func GetRates(to string) (Rates, error) {
	switch to {
	case "SEK":
		euro, err := getRate("EUR", to)

		if err != nil {
			return Rates{}, err
		}

		dollars, err := getRate("USD", to)

		return Rates{
			First:  euro,
			Second: dollars,
		}, nil

	case "EUR":
		sek, err := getRate("SEK", to)

		if err != nil {
			return Rates{}, err
		}

		dollars, err := getRate("USD", to)

		return Rates{
			First:  sek,
			Second: dollars,
		}, nil

	case "USD":
		euro, err := getRate("EUR", to)

		if err != nil {
			return Rates{}, err
		}

		sek, err := getRate("SEK", to)

		return Rates{
			First:  euro,
			Second: sek,
		}, nil

	default:
		return Rates{}, errors.New("Invalid currency")
	}
}

// Converts between currencies
func Convert(from, to string, amount float64) (float64, error) {
	// Connect to the API
	result, err := http.Get(fmt.Sprintf("https://anyapi.io/api/v1/exchange/convert?apiKey=%v&base=%v&to=%v&amount=%v", API_KEY, from, to, amount))

	if err != nil { // If API cannot be reached
		return 0, err
	}

	// Reads the response as bytes
	response, err := ioutil.ReadAll(result.Body)

	if err != nil {
		return 0, err
	}

	var value Value        // Stores the response as a struct
	var conversion float64 // Holds the converted value

	// Read the response into struct
	json.Unmarshal(response, &value)

	conversion = value.CvtAmt // Store converted value as a variable

	return conversion, nil // return the converted value
}
