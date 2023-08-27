package memory

import (
	"errors"
	"simple-go-ddd/aggregate"
	"simple-go-ddd/domain/customer"
	"testing"

	"github.com/google/uuid"
)

type testCase struct {
	id          uuid.UUID
	name        string
	expectedErr error
}

func TestMemory_GetCustom(t *testing.T) {
	cust, err := aggregate.NewCustomer("John Doe")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()

	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}

	testCase := []testCase{
		{
			id:          uuid.New(),
			name:        "no customer found",
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			id:          id,
			name:        "customer found",
			expectedErr: nil,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
