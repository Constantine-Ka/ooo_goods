package repository

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"strings"
)

func (db *DB) MigrationUp() {
	db.Log.Named("Миграции")
	entries, err := os.ReadDir("./migration/")
	if err != nil {
		db.Log.Fatal("Не найдена папка", zap.String("folder", "./migration/"))
	}

	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".up.sql") {
			table := strings.TrimSuffix(e.Name(), ".up.sql")
			filename := fmt.Sprintf("./migration/%s.up.sql", table)
			content, err := os.ReadFile(filename)
			if err != nil {
				db.Log.Fatal("Не получается открыть файл", zap.String("folder", "./migration/"), zap.String("filename", filename))
				return
			}
			query := string(content)
			query = strings.Replace(query, "\t", " ", -1)
			query = strings.Replace(query, "\n", " ", -1)
			_, err = db.DB.Exec(query)
			if err != nil {
				db.Log.Error("Ошибка выполнения запроса", zap.String("folder", "./migration/"), zap.String("filename", filename), zap.Error(err))
				//os.Exit(200)
				return
			}
			db.tablename = table
		}

	}

}
