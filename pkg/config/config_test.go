package config

// TODO: разобраться с моками
//func TestInit_GoDotEnv(t *testing.T) {
//	type testCase struct {
//		test        string
//		config      godotenv.GoDotEnv
//		expectedErr error
//	}
//
//	testCases := []testCase{
//		{
//			test:        "Успех godotenv с путем к существующему файлу",
//			config:      godotenv.New("../../.env"),
//			expectedErr: nil,
//		}, {
//			test:        "Провал godotenv без указания пути к файлу",
//			config:      godotenv.New(),
//			expectedErr: godotenv.ErrNotFound,
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.test, func(t *testing.T) {
//			_, err := Load(tc.config)
//            assert.ErrorIs(t, tc.expectedErr, err)
//		})
//	}
//}
//
//func TestInit_Viperr(t *testing.T) {
//	type testCase struct {
//		test        string
//		config      viperr.Viperr
//		expectedErr error
//	}
//
//    testCases := []testCase{
//        {
//            test:        "Успех viperr с путем к существующему файлу",
//            config:      viperr.New("config"),
//            expectedErr: nil,
//        }, {
//            test:        "Успех viperr без указания пути к файлу",
//            config:      viperr.New(),
//            expectedErr: nil,
//        }, {
//            test:        "Провал viperr с несуществующим файлом",
//            config:      viperr.New("config_not_exist"),
//            expectedErr: viperr.ErrNotFound,
//        },
//    }
//
//	for _, tc := range testCases {
//		t.Run(tc.test, func(t *testing.T) {
//            tc.config.AddConfigPath("../../internal/config")
//            tc.config.AddConfigType("yml")
//
//			_, err := Load(tc.config)
//            assert.ErrorIs(t, tc.expectedErr, err)
//		})
//	}
//}
