package builder

import "testing"

func TestDirector(t *testing.T) {
	d := Director{}

	car := &CarBuilder{}
	d.SetBuilder(car).Construct()
	v := car.GetVehicle()

	if v.Wheels != 4 {
		t.Errorf("Car Wheels should be 4, but it is %v", v.Wheels)
	}
	if v.Name != "Car" {
		t.Errorf("Name should be Car, but it is %v", v.Name)
	}
	if v.Seats != 4 {
		t.Errorf("Seats should be 4, but it is %v", v.Seats)
	}

	moto := &MotoBuilder{}
	d.SetBuilder(moto)
	d.Construct()
	v = moto.GetVehicle()

	if v.Wheels != 2 {
		t.Errorf("Moto Wheels should be 2, but it is %v", v.Wheels)
	}
	if v.Name != "Moto" {
		t.Errorf("Name should be Moto, but it is %v", v.Name)
	}
	if v.Seats != 2 {
		t.Errorf("Seats should be 2, but it is %v", v.Seats)
	}
}
