package absFactory

import "github.com/pkg/errors"

type Vehicle interface {
	NumWheels() int
	NumSeats() int
}

type VehicleFactory interface {
	Build(v int) (Vehicle, error)
}

const(
	CarFactoryType = iota
	MotoFactoryType
)

func BuildFactory(v int) (VehicleFactory, error) {
	switch v{
	case CarFactoryType:
		return new(CarFactory), nil
	case MotoFactoryType:
		return new(MotoFactory), nil
	default:
		return nil, errors.New("Unsupported Factory")
	}
}