package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"ooo_goods/config"
)

func NewDB(config config.CFG) DB {
	// Инициализируем подключение к базе данных
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", config.DB.Login, config.DB.Pass, config.DB.Host, config.DB.Name))
	if err != nil {
		config.Log.Fatal("Некорректные данные, для подключения к БД", zap.String("host", config.DB.Host), zap.String("name", config.DB.Name), zap.String("login", config.DB.Login), zap.String("err", err.Error()))
		return DB{}
	}
	// Проверяем, что подключение к БД еще рабочее
	err = db.Ping()
	if err != nil {
		config.Log.Fatal("Сервер сразу сбрасывает подключение  подключение к БД", zap.String("host", config.DB.Host), zap.String("name", config.DB.Name), zap.String("login", config.DB.Login), zap.String("err", err.Error()))
		return DB{}
	}
	config.Log.Info("Инициировано подключение к Базе")
	return DB{DB: db, Log: config.Log}
}
