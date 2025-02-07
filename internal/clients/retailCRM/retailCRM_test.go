package retailCRM

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest" // Для создания тестового HTTP-сервера
	"ooo_goods/config"
	"testing"

	"github.com/stretchr/testify/assert" // Для удобных проверок
)

func TestGetOrders(t *testing.T) {
	// Настройка тестового логгера
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	// Тестовые случаи
	testCases := []struct {
		name           string
		serverResponse string // Ответ сервера
		expectedOrders []Order
		expectedPage   int
		expectError    bool
	}{
		
		{
			name:           "Server error", // Ошибка сервера
			serverResponse: `{"success": false}`,
			expectError:    true, // Ожидаем ошибку парсинга JSON
		},
		{
			name:           "Invalid JSON", // Некорректный JSON
			serverResponse: `{неправильный json}`,
			expectError:    true, // Ожидаем ошибку парсинга JSON
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Создаем тестовый HTTP-сервер
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, tc.serverResponse) // Отправляем ответ сервера
			}))
			defer server.Close() // Закрываем сервер после теста

			// Настраиваем тестовую конфигурацию
			cfg := config.CFG{
				Domain: server.URL[7:], // Извлекаем домен из URL тестового сервера (без http://)
				Apikey: "test_api_key", // Тестовый API-ключ
				Log:    logger,         // Наш тестовый логгер
			}

			// Вызываем функцию GetOrders
			orders, page := GetOrders(cfg, 1)

			// Проверки
			if tc.expectError {
				// Ожидаем ошибку (проверяем, что заказы и страница не заполнены)
				assert.Empty(t, orders, "Ожидаем пустой слайс заказов при ошибке")
				assert.Equal(t, 0, page, "Ожидаем страницу 0 при ошибке")
			} else {
				assert.Equal(t, tc.expectedOrders, orders, "Заказы не совпадают")
				assert.Equal(t, tc.expectedPage, page, "Номер страницы не совпадает")
			}
		})
	}
}
