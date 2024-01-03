package config

import (
	"github.com/eliofery/golang-angular/pkg/config/godotenv"
	"github.com/eliofery/golang-angular/pkg/config/viperr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	type testCase struct {
		test        string
		config      Config
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Успех godotenv",
			config:      godotenv.New("../../.env"),
			expectedErr: nil,
		}, {
			test:        "Успех viperr",
			config:      viperr.New("config"),
			expectedErr: nil,
		}, {
			test:        "Успех viperr",
			config:      viperr.New(),
			expectedErr: nil,
		}, {
			test:        "Провал godotenv",
			config:      godotenv.New(),
			expectedErr: godotenv.ErrNotFound,
		}, {
			test:        "Провал viperr",
			config:      viperr.New("config_not_exist"),
			expectedErr: viperr.ErrNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			v, ok := tc.config.(*viperr.Viper)
			if ok {
				v.AddConfigPath("../../internal/config")
				v.AddConfigType("yml")
			}

			_, err := Init(tc.config)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
