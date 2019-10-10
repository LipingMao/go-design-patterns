package builder

type Vehicle struct {
	Wheels int
	Seats  int
	Name   string
}

type CarBuilder struct {
	v Vehicle
}

func (b *CarBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 4
	return b
}

func (b *CarBuilder) SetSeats() BuildProcess {
	b.v.Seats = 4
	return b
}

func (b *CarBuilder) SetName() BuildProcess {
	b.v.Name = "Car"
	return b
}

func (b *CarBuilder) GetVehicle() Vehicle {
	return b.v
}

type MotoBuilder struct {
	v Vehicle
}

func (b *MotoBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *MotoBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

func (b *MotoBuilder) SetName() BuildProcess {
	b.v.Name = "Moto"
	return b
}

func (b *MotoBuilder) GetVehicle() Vehicle {
	return b.v
}

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetName() BuildProcess
	GetVehicle() Vehicle
}

type Director struct {
	builder BuildProcess
}

func (d *Director) SetBuilder(b BuildProcess) *Director {
	d.builder = b
	return d
}

func (d *Director) Construct() {
	d.builder.SetSeats().SetWheels().SetName()
}
