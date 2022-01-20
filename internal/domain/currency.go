package domain

type CurrencyRepository interface {
	ConvertToUSD(currencies []string) (map[string]float32, error)
}
