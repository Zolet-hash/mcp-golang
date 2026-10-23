
package currency

import (
	"fmt"
	"strings"
)

type Converter struct {
	rates map[string]float64
}

func NewConverter() *Converter {
	return &Converter{
		rates: map[string]float64{
			"USD:KES": 130.0,
			"EUR:KES": 140.0,
			"GBP:KES": 165.0,
			"USD:EUR": 0.92,
			"EUR:USD": 1.09,
		},
	}
}

func (c *Converter) Convert(
	amount float64,
	from string,
	to string,
) (float64, error) {
	if amount <= 0 {
		return 0, fmt.Errorf("amount must be greater than zero")
	}

	from = strings.ToUpper(strings.TrimSpace(from))
	to = strings.ToUpper(strings.TrimSpace(to))

	if from == to {
		return amount, nil
	}

	key := from + ":" + to
	rate, ok := c.rates[key]
	if !ok {
		// Use the inverse rate when available.
		inverseKey := to + ":" + from
		inverseRate, inverseOK := c.rates[inverseKey]
		if !inverseOK {
			return 0, fmt.Errorf(
				"unsupported currency pair: %s to %s",
				from, to,
			)
		}
		rate = 1 / inverseRate
	}

	return amount * rate, nil
}
