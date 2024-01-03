package godotenv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoad(t *testing.T) {
	type testCase struct {
		test        string
		configName  string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Успех",
			configName:  "../../../.env",
			expectedErr: nil,
		}, {
			test:        "Провал",
			configName:  "",
			expectedErr: ErrNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			err := New(tc.configName).Load()
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
