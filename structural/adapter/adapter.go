package adapter

import "fmt"

type LagencyPrinter interface {
	Print(s string) string
}

type MyLagencyPrinter struct{}

func (l *MyLagencyPrinter) Print(s string) string {
	msg := fmt.Sprintf("Lagency Printer: %s\n", s)
	println(msg)
	return msg
}

type ModernPrinter interface {
	PrintStored() string
}

type PrinterAdapter struct {
	OldPrinter LagencyPrinter
	Msg        string
}

func (p *PrinterAdapter) PrintStored() string {
	msg := fmt.Sprintf("Modern Printer: %s\n", p.Msg)
	println(msg)
	return msg
}
