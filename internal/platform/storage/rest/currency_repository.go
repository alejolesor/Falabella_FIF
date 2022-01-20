package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

type CurrencyRepository struct{}

func NewCurrencyRepository() *CurrencyRepository {
	return &CurrencyRepository{}
}

func (m *CurrencyRepository) ConvertToUSD(currencies []string) (map[string]float32, error) {
	accessKey := "fb7b98c699b393b1dcb6013f82fab396"
	url := fmt.Sprintf("http://api.currencylayer.com/live?access_key=%s&currencies=%s", accessKey, strings.Join(currencies, ","))
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	bytes := gjson.ParseBytes(body)

	quotes := make(map[string]float32)
	for _, currency := range currencies {
		quote := bytes.Get(fmt.Sprintf("quotes.USD%s", currency)).Float()
		quotes[currency] = float32(quote)
	}
	return quotes, nil
}
