package repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type DB struct {
	DB        *sqlx.DB
	Log       *zap.Logger
	tablename string
}
