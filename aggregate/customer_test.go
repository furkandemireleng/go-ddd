package aggregate

import (
	"errors"
	"testing"
)

func TestNewCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		age         int
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			age:         12,
			expectedErr: ErrInvalidName,
		},
		{
			test:        "Valid name validation",
			name:        "Name",
			age:         12,
			expectedErr: nil,
		},
		{
			test:        "Empty age validation",
			name:        "Furkan",
			age:         0,
			expectedErr: ErrInvalidAge,
		},
		{
			test:        "Valid age validation",
			name:        "Name",
			age:         12,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := NewCustomer(tc.name, tc.age)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected: %v, got: %v", tc.expectedErr, err)
			}
		})
	}

}
