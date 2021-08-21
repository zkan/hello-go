package client

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zkan/hello-go/making_http_requests/model"
)

//Fetch is exported ...
func FetchCrypto(fiat string, crypto string) (string, error) {
	//Build The URL string
	URL := "https://api.nomics.com/v1/currencies/ticker?key=3990ec554a414b59dd85d29b2286dd85&interval=1d&ids=" + crypto + "&convert=" + fiat
	//We make HTTP request using the Get function
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal("ooopsss an error occurred, please try again")
	}
	defer resp.Body.Close()
	//Create a variable of the same type as our model
	var cResp model.Cryptoresponse
	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("ooopsss! an error occurred, please try again")
	}
	//Invoke the text output function & return it with nil as the error value
	return cResp.TextOutput(), nil
}
