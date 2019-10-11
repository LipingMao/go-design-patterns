package absFactory

import "github.com/pkg/errors"

type Moto interface {
	GetMotoType() int
}

const (
	SportMotoType = iota
	CruiseMotoType
)



type SportMoto struct {}

func (* SportMoto) NumSeats() int {
	return 2
}

func (* SportMoto) NumWheels() int {
	return 2
}

func (* SportMoto) GetMotoType() int {
	return SportMotoType
}

type CruiseMoto struct {}

func (* CruiseMoto) NumSeats() int {
	return 1
}

func (* CruiseMoto) NumWheels() int {
	return 2
}

func (* CruiseMoto) GetMotoType() int {
	return CruiseMotoType
}

type MotoFactory struct {}

func (* MotoFactory) Build(v int) (Vehicle, error) {
	switch  v {
	case SportMotoType:
		return new(SportMoto), nil
	case CruiseMotoType:
		return new(CruiseMoto), nil
	default:
		return nil, errors.New("Unsupported Moto")
	}
}