package aggregate_test

import (
	"simple-go-ddd/aggregate"
	"testing"
)

type testCase struct {
	test        string
	name        string
	expectedErr error
}

func TestCustomer_NewCustomer(t *testing.T) {
	testCases := []testCase{
		{
			test:        "empty name",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		}, 
		{
			test:        "valid name",
			name:        "John Doe",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)
			if err != tc.expectedErr {
				t.Errorf("expected %v, got %v", tc.expectedErr, err)
			}
		})
	}	
}
