package config

import (
	"flag"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	// Создаем тестовый логгер
	//testLogger := zaptest.NewLogger(t)

	// Сохраняем оригинальные значения переменных окружения, чтобы восстановить их после теста
	originalApikey := os.Getenv("Apikey")
	originalDomain := os.Getenv("Domain")
	originalDbhost := os.Getenv("Dbhost")
	originalDbname := os.Getenv("Dbname")
	originalDblogin := os.Getenv("DbLogin")
	originalDbpass := os.Getenv("Dbpass")

	defer func() {
		// Восстанавливаем оригинальные значения переменных окружения
		os.Setenv("Apikey", originalApikey)
		os.Setenv("Domain", originalDomain)
		os.Setenv("Dbhost", originalDbhost)
		os.Setenv("Dbname", originalDbname)
		os.Setenv("DbLogin", originalDblogin)
		os.Setenv("Dbpass", originalDbpass)
	}()

	// Функция для очистки флагов после теста
	resetFlags := func() {
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.PanicOnError)
	}

	t.Run("Missing required fields", func(t *testing.T) {
		resetFlags()
		createConfigFile(t, "") // Пустой конфиг

		// Перехватываем panic и проверяем сообщение
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					t.Errorf("Паника должна быть типа error, получено %T", r)
				}
				expectedErrors := []string{
					"Не указан флаг конфигурации. FIELD: Apikey",
					"Не указан флаг конфигурации. FIELD: Domain",
					"Не указан флаг конфигурации. FIELD: Dbhost",
					"Не указан флаг конфигурации. FIELD: Dbname",
					"Не указан флаг конфигурации. FIELD: ", // Для Dblogin
				}

				for _, expectedError := range expectedErrors {
					if expectedError != expectedError {
						t.Errorf("Ожидалась ошибка: %s, получено: %s", expectedError, err.Error())
					}
				}
			} else {
				t.Error("Ожидалась паника, но она не произошла")
				t.Error("Ожидалась паника, но она не произошла")
			}
		}()

		Init() // Не должно быть паники
	})

}

func createConfigFile(t *testing.T, content string) {
	t.Helper()
	err := os.WriteFile("config-test.yaml", []byte(content), 0644)
	if err != nil {
		t.Fatalf("Ошибка создания конфиг файла: %v", err)
	}
	t.Cleanup(func() {
		os.Remove("config-test.yaml")
	})
}
