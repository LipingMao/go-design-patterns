package prototype

import (
	"fmt"
	"testing"
)

func TestClone(t *testing.T) {
	shirtCloner := GetShirtCloner()
	white, err:= shirtCloner.GetClone(White)
	if err != nil {
		t.Errorf("Error to get White shirt, %v", err)
	}
	fmt.Print(white.GetInfo())
}