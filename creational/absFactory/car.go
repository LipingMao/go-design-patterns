package absFactory

import "github.com/pkg/errors"

type Car interface {
	NumDoors() int
}

const (
	LuxuryCarType = iota
	FamilyCarType
)

type LuxuryCar struct{}

func (* LuxuryCar) NumDoors() int {
	return 5
}

func (* LuxuryCar) NumWheels() int {
	return 6
}

func (* LuxuryCar) NumSeats() int {
	return 7
}

type FamilyCar struct {}

func (* FamilyCar) NumDoors() int {
	return 4
}

func (* FamilyCar) NumWheels() int {
	return 4
}

func (* FamilyCar) NumSeats() int {
	return 4
}

type CarFactory struct{}

func (c *CarFactory) Build(v int) (Vehicle, error){
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New("Unsupported Cars")
	}
}
