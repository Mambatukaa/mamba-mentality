package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"mamba-mentality.com/datatypes"
)

const apiUrl string = "https://cex.io/api/ticker/%s/USD";

func GetRate(currency string) (*datatypes.Rate, error) {
	upCurrency := strings.ToUpper(currency);

	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency));

	if(err != nil) {
		return nil, err;
	};

	var response CEXResponse;

	// we have response
	if(res.StatusCode == http.StatusOK) {
		// await
		bodyBytes, err := io.ReadAll(res.Body);

		if(err != nil) {
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &response)

		if(err != nil) {
			return nil, err
		}

		
	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode); 
	};

	rate := datatypes.Rate{Currency: currency, Price: response.Bid}

	return &rate, nil;
};
