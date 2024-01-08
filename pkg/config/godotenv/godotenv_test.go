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
			test:        "Успех с указанием имени файла",
			configName:  "../../../.env",
			expectedErr: nil,
		}, {
			test:        "Провал без указания имени файла",
			configName:  "",
			expectedErr: ErrNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			err := New(tc.configName).Init()
			assert.ErrorIs(t, tc.expectedErr, err)
		})
	}
}
