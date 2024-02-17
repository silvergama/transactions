package app

import (
	"fmt"

	"github.com/silvergama/transations/config"
	"github.com/silvergama/transations/infrastructure"
	accountRepo "github.com/silvergama/transations/internal/account/postgres"
	"github.com/silvergama/transations/internal/http"
	transactionRepo "github.com/silvergama/transations/internal/transaction/postgres"
	"github.com/silvergama/transations/pkg/logger"
	"go.uber.org/zap"
)

func Run() {
	cfg := config.ReadProperties()

	stringConnection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Database.User,
		cfg.Database.Pwd,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Base,
	)

	db, err := infrastructure.NewDBConnection(stringConnection)
	if err != nil {
		logger.Error(
			"failed to create an instance of the database",
			zap.Error(err),
		)
		return
	}

	account := accountRepo.NewAccount(db.Connection)
	transaction := transactionRepo.NewTransaction(db.Connection)

	http.StartServer(cfg, http.Services{Account: account, Transaction: transaction})
}
