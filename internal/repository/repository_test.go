package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "gorm.io/driver/mysql"
	"ooo_goods/config"
)

// Цель: Проверить, что функция NewDB успешно устанавливает соединение с базой данных при корректных параметрах конфигурации.
func TestNewDB(t *testing.T) {

	testCases := []struct {
		name        string
		cfg         config.CFG
		mockSetup   func(sqlmock.Sqlmock)
		expectError bool
	}{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Ошибка при создании mock DB: %v", err)
			}
			defer db.Close()

			tc.mockSetup(mock)

			// Вызываем тестируемую функцию
			result := NewDB(tc.cfg)

			if tc.expectError {
				if result.DB != nil {
					t.Error("Ожидалась ошибка, получено соединение")
				}
			} else {
				if result.DB == nil {
					t.Error("Ожидалось соединение, получена ошибка")
				}
				if err := mock.ExpectationsWereMet(); err != nil {
					t.Errorf("Не все ожидания выполнены: %s", err)
				}
			}
		})
	}
}
