package enums

import (
	"fmt"
)

type Currency string

const (
	EUR Currency = "EUR"
	USD Currency = "USD"
	RUB Currency = "RUR"
)

// Using float for payments can be dangerous, so there are 2 choices in the structure
//
// Float64ToMoney - if you are 100% sure that your math is correct
//
// Monetary, Cents - created to be able to manage an basic monetary unit and 1/100 of it
// Like, for example, ruble and kopeck or dollar and cent
type Money struct {
	Float64ToMoney float64
	Monetary       int64
	Cents          int64
}

func (c Money) String() string {
	if c.Float64ToMoney > 0 {
		return fmt.Sprintf("%.2f", c.Float64ToMoney)
	}
	return fmt.Sprintf("%d.%02d", c.Monetary, c.Cents%100)
}
