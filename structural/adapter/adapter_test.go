package adapter

import (
	"strings"
	"testing"
)

func TestPrinterAdapter_PrintStored(t *testing.T) {
	adapter := PrinterAdapter{OldPrinter: &MyLagencyPrinter{}, Msg: "Hello World"}
	old := adapter.OldPrinter.Print("Hello World")
	if !strings.Contains(old, "Lagency") {
		t.Errorf("OldPrinter Error : %v", old)
	}
	new := adapter.PrintStored()
	if !strings.Contains(new, "Modern") {
		t.Errorf("Modern Error : %v", new)
	}
}
