package aggregate

import (
	"errors"
	"simple-go-ddd/entity"
	"simple-go-ddd/valueobject"

	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("invalid person")
)

type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []*valueobject.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		ID:   uuid.New(),
		Name: name,
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]*valueobject.Transaction, 0),
	}, nil
}
