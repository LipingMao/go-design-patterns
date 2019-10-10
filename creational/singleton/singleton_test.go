package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	c1 := GetInstance()
	if c1 == nil {
		t.Error("New singleton should be generated")
	}

	currectCount := c1.AddOne()
	if currectCount != 1 {
		t.Errorf("Must be 1 after 1 call, but it is %v", currectCount)
	}

	c2 := GetInstance()
	if c2 != c1 {
		t.Error("singleton must be single")
	}

	currectCount = c2.AddOne()
	if currectCount != 2 {
		t.Errorf("Must be 2 after 2 call, but it is %v", currectCount)
	}
}
