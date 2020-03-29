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

type Money struct {
	Monetary int64
	Cents    int64
}

func (c Money) String() string {
	return fmt.Sprintf("%d.%02d", c.Monetary, c.Cents%100)
}
