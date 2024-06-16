package services

import (
	"gorm.io/gorm"
	"log"
	"rabbit-food/config"
	"rabbit-food/internal/models/user"
	"rabbit-food/pkg/adapters/storage"
)

type AppContainer struct {
	cfg         config.Config
	dbConn      *gorm.DB
	UserService *UserService
}

func NewAppContainer(cfg config.Config) (*AppContainer, error) {
	app := &AppContainer{
		cfg: cfg,
	}

	app.mustInitDB()
	storage.Migrate(app.dbConn)

	app.setUserService()

	return app, nil
}

func (a *AppContainer) mustInitDB() {
	if a.dbConn != nil {
		return
	}

	db, err := storage.NewPostgresGormConnection(a.cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	a.dbConn = db
}

func (a *AppContainer) setUserService() {
	if a.UserService != nil {
		return
	}
	a.UserService = NewUserService(user.NewUserOps(a.dbConn, storage.NewUserRepo(a.dbConn)))
}
