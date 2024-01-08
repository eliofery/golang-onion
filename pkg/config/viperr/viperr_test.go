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
			test:        "Успех с указанием имени файла",
			configName:  "config",
			expectedErr: nil,
		}, {
			test:        "Успех без указания имени файла",
			configName:  "",
			expectedErr: nil,
		}, {
			test:        "Провал с указанием не существующего файла",
			configName:  "config_not_exist",
			expectedErr: ErrNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			config := New(tc.configName)
			config.AddConfigPath("../../../internal/config")
			config.AddConfigType("yml")

			err := config.Init()
			assert.ErrorIs(t, tc.expectedErr, err)
		})
	}
}
