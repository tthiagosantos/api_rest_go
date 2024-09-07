package entity

import (
	"errors"
	"github.com/api_golang/pkg/entity"
	"time"
)

var (
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrPriceIvalid     = errors.New("price is invalid")
)

type Product struct {
	ID        entity.ID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Price     float64   `json:"price,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {

	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Time{},
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p Product) Validate() error {
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price < 0 {
		return ErrPriceIvalid
	}
	if p.Price == 0 {
		return ErrPriceIsRequired
	}
	return nil
}
