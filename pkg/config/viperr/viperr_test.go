package viperr

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
			configName:  "config",
			expectedErr: nil,
		}, {
			test:        "Успех",
			configName:  "",
			expectedErr: nil,
		}, {
			test:        "Провал",
			configName:  "config_not_exist",
			expectedErr: ErrNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			config := New(tc.configName)
			config.AddConfigPath("../../../internal/config")
			config.AddConfigType("yml")

			err := config.Load()
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
