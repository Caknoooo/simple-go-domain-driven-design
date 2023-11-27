package memory

import (
	"fmt"
	"simple-go-ddd/aggregate"
	"simple-go-ddd/domain/customer"
	"sync"

	"github.com/google/uuid"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (r *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := r.customers[id]; ok {
		return customer, nil
	}

	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (r *MemoryRepository) Add(c aggregate.Customer) error {
	if r.customers == nil {
		r.Lock()
		r.customers = make(map[uuid.UUID]aggregate.Customer)
		r.Unlock()
	}

	if _, ok := r.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToAddCustomer)
	}

	r.Lock()
	r.customers[c.GetID()] = c
	r.Unlock()

	return nil
}

func (r *MemoryRepository) Update(c aggregate.Customer) error {
	if _, ok := r.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}

	r.Lock()
	r.customers[c.GetID()] = c
	r.Unlock()

	return nil
}