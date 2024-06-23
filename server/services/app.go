package services

import (
	"gorm.io/gorm"
	"server/config"
	"server/internal/models/user"
	"server/internal/models/wallet/wallet"
	"server/pkg/adapters/storage"
	"server/pkg/logger"
)

type AppContainer struct {
	cfg         config.Config
	dbConn      *gorm.DB
	logger      *logger.CustomLogger
	UserService *UserService
}

func NewAppContainer(cfg config.Config, logger *logger.CustomLogger) (*AppContainer, error) {
	app := &AppContainer{
		cfg:    cfg,
		logger: logger,
	}

	app.mustInitDB()
	err := storage.Migrate(app.dbConn)
	if err != nil {
		logger.Fatal("Migration failed: ", err)
	}

	app.setAuthService([]byte(cfg.Server.TokenSecret), uint(cfg.Server.TokenExpMinutes), uint(cfg.Server.RefreshTokenExpMinutes))
	app.setWalletService()

	return app, nil
}

func (a *AppContainer) mustInitDB() {
	if a.dbConn != nil {
		return
	}

	db, err := storage.NewPostgresGormConnection(a.cfg.DB)
	if err != nil {
		a.logger.Fatal(err)
	}

	a.dbConn = db
}

func (a *AppContainer) setUserService() {
	if a.UserService != nil {
		a.logger.Error("Error In Running Service")
		return
	}
	a.WalletService = NewWalletService(wallet.NewWalletOps(a.dbConn, storage.NewWalletRepo(a.dbConn)))
}
