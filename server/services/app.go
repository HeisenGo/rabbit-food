package services

import (
	"gorm.io/gorm"
	"log"
	"rabbit-food/server/config"
	"rabbit-food/server/internal/models/user"
	storage2 "rabbit-food/server/pkg/adapters/storage"
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
	err := storage2.Migrate(app.dbConn)
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}

	app.setUserService()

	return app, nil
}

func (a *AppContainer) mustInitDB() {
	if a.dbConn != nil {
		return
	}

	db, err := storage2.NewPostgresGormConnection(a.cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	a.dbConn = db
}

func (a *AppContainer) setUserService() {
	if a.UserService != nil {
		return
	}
	a.UserService = NewUserService(user.NewUserOps(a.dbConn, storage2.NewUserRepo(a.dbConn)))
}
