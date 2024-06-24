package services

import (
	"log"
	"server/config"
	"server/internal/models/user"
	"server/pkg/adapters/storage"

	"gorm.io/gorm"
)

type AppContainer struct {
	cfg         config.Config
	dbConn      *gorm.DB
	AuthService *AuthService
	UserService *UserService
}

func NewAppContainer(cfg config.Config) (*AppContainer, error) {
	app := &AppContainer{
		cfg: cfg,
	}

	app.mustInitDB()
	err := storage.Migrate(app.dbConn)
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}

	app.setAuthService([]byte(cfg.Server.TokenSecret), uint(cfg.Server.TokenExpMinutes), uint(cfg.Server.RefreshTokenExpMinutes))
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

func (a *AppContainer) setAuthService(secret []byte,
	tokenExpiration uint, refreshTokenExpiration uint) {
	if a.AuthService != nil {
		return
	}
	a.AuthService = NewAuthService(user.NewUserOps(a.dbConn, storage.NewUserRepo(a.dbConn)), secret, tokenExpiration, refreshTokenExpiration)
}

func (a *AppContainer) setUserService() {
	if a.UserService != nil {
		return
	}
	a.UserService = NewUserService(user.NewUserOps(a.dbConn, storage.NewUserRepo(a.dbConn)))
}
