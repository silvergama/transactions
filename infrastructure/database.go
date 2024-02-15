package infrastructure

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/silvergama/transations/pkg/logger"
	"go.uber.org/zap"
)

type Database struct {
	Connection *sql.DB
}

func NewDatabase(connectionString string) (*Database, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		logger.Error(
			"failed to connect database",
			zap.Error(err),
		)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Error("failed to ping database",
			zap.Error(err),
		)
		return nil, err
	}

	return &Database{Connection: db}, nil
}
