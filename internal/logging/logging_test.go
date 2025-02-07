package logging

import (
	"os"
	"testing"
)

// Цель: Проверить, что функция CreateLogger успешно создает экземпляр логгера zap.Logger.
func TestCreateLogger(t *testing.T) {
	logger := CreateLogger()
	if logger == nil {
		t.Error("Ожидался не нулевой логгер")
	}
}

// Цель: Убедиться, что функция CreateLogger создает папку "logs", если она не существует.
func TestCreateLogFolder(t *testing.T) {
	logFolder := "logs"
	os.RemoveAll(logFolder) // Удаляем папку, если она есть
	CreateLogger()
	if _, err := os.Stat(logFolder); os.IsNotExist(err) {
		t.Error("Папка logs не была создана")
	}
}
