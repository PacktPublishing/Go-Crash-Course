package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type coinData struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
}

const TICKER_URL = "https://api.binance.com/api/v3/ticker/price"

func getCoinPriceInUSDT(symbol string) (float64, error) {
	req, _ := http.NewRequest("GET", TICKER_URL, nil)
	q := req.URL.Query()
	q.Add("symbol", symbol)
	req.URL.RawQuery = q.Encode()

	resp, err := http.Get(req.URL.String())
	if err != nil {
		return 0, fmt.Errorf("error in fetching coins data: %s", err.Error())
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("error in reading Binance response: %s", err.Error())
	}
	//fmt.Println(string(respBody))

	var data coinData
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return 0, fmt.Errorf("error in parsing Binance response: %s", err.Error())
	}

	return data.Price, nil
}

func main() {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		price, err := getCoinPriceInUSDT("ADAUSDT")
		if err != nil {
			fmt.Printf("Error fetching ADA price data: %s\n", err.Error())
		}
		fmt.Printf("ADA - $%f\n", price)
	}
}
