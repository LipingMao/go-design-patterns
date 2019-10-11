package absFactory

import "testing"

func TestCarFactory(t *testing.T) {
	Car, err := BuildFactory(CarFactoryType)
	if err != nil{
		t.Errorf("Fail to get car.")
	}
	luxury, err := Car.Build(LuxuryCarType)
	if err != nil {
		t.Errorf("Fail to get luxury")
	}
	seats := luxury.NumSeats()
	if seats != 7{
		t.Errorf("Luxury seat should be 7, but it is %v", seats)
	}
	wheels := luxury.NumWheels()
	if wheels != 6{
		t.Errorf("Luxury wheels should be 6, but it is %v", wheels)
	}
}