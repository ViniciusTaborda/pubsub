package uuid

import (
	"testing"
)

func TestMustNewUUID(t *testing.T) {
	testCases := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "MustNewUUID returns a non-empty string",
			test: func(t *testing.T) {
				uuid := MustNewUUID()
				if uuid == "" {
					t.Errorf("Expected non-empty string, got empty string")
				}
			},
		},
		{
			name: "MustNewUUID returns a valid UUID",
			test: func(t *testing.T) {
				uuid := MustNewUUID()
				if len(uuid) != 36 {
					t.Errorf("Expected UUID of length 36, got UUID of length %d", len(uuid))
				}
			},
		},
		{
			name: "MustNewUUID returns unique UUIDs",
			test: func(t *testing.T) {
				uuid1 := MustNewUUID()
				uuid2 := MustNewUUID()
				if uuid1 == uuid2 {
					t.Errorf("Expected unique UUIDs, got identical UUIDs")
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.test)
	}
}
