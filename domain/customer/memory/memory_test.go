package memory

import (
	"errors"
	"github.com/furkandemireleng/go-ddd/aggregate"
	customer2 "github.com/furkandemireleng/go-ddd/domain/customer"
	"github.com/google/uuid"
	"testing"
)

func TestMemoryRepository_Get(t *testing.T) {
	type testCase struct {
		name        string
		age         int
		id          uuid.UUID
		expectedErr error
	}

	customer, err := aggregate.NewCustomer("Furkan", 12)
	if err != nil {
		t.Fatal(err)
	}
	id := customer.GetId()

	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: customer,
		},
	}

	testCases := []testCase{
		{
			name:        "no customer by id",
			age:         12,
			id:          uuid.MustParse("7f397108-ae35-4918-9245-08a12141fe30"),
			expectedErr: customer2.ErrCustomerNotFound,
		},
		{
			name:        "customer by id",
			age:         12,
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)

			if !errors.Is(err, tc.expectedErr) {
				t.Fatalf("expected: %v, got: %v", tc.expectedErr, err)
			}

		})
	}
}
