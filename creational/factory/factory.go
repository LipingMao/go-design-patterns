package factory

import (
	"fmt"
	"github.com/pkg/errors"
)

const (
	Cash = iota
	Credit
)

type PaymentMethod interface {
	Pay(amount float32) string
}

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case Credit:
		return new(CreditPM), nil
	default:
		return nil, errors.New("Unsupported")
	}
}

type CashPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("Cash : %v", amount)
}

type CreditPM struct{}

func (c *CreditPM) Pay(amount float32) string {
	return fmt.Sprintf("Credit : %v", amount)
}
