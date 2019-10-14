package prototype

import (
	"fmt"
	"github.com/pkg/errors"
)

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

func GetShirtCloner () ShirtCloner {
	return &ShirtCache{}
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

type ItemInfoGetter interface {
	GetInfo() string
}
type ShirtColor byte
type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Info: SKU - %v, Color - %v, Price - %v", s.SKU, s.Color, s.Price)
}
func (s *Shirt) GetPrice() float32 {
	return s.Price
}
func (s *Shirt) GetSKU() string {
	return s.SKU
}
func (s *Shirt) GetColor() ShirtColor {
	return s.Color
}

var white = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}
var black = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}
var blue = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

type ShirtCache struct{}

func (c *ShirtCache) GetClone(s int) (ItemInfoGetter, error) {
	switch s {
	case White:
		n := *white
		return &n, nil
	case Black:
		n := *black
		return &n, nil
	case Blue:
		n := *blue
		return &n, nil
	default:
		return nil, errors.New("Unsupported")
	}
}