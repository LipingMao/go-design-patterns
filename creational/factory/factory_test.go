package factory

import (
	"strings"
	"testing"
)

func TestCash(t *testing.T) {
	cash, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Errorf("unexpected error : %v", err)
	}

	msg := cash.Pay(1.0)
	if !strings.Contains(msg, "Cash") {
		t.Errorf("Cash pay error")
	}
}

func TestCredit(t *testing.T) {
	credit, err := GetPaymentMethod(Credit)
	if err != nil {
		t.Errorf("unexpected error : %v", err)
	}

	msg := credit.Pay(1.0)
	if !strings.Contains(msg, "Credit") {
		t.Errorf("Credit pay error")
	}
}

func TestUnsupported(t *testing.T) {
	_, err := GetPaymentMethod(3)
	if err == nil {
		t.Error("Unsupported should return error")
	}
}
