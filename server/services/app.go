package services

import (
	"log"
	"server/config"
	"server/internal/models/restaurant/menu"
	"server/internal/models/restaurant/restaurant"
	"server/internal/models/user"
	creditCard "server/internal/models/wallet/credit_card"
	"server/internal/models/wallet/wallet"
	"server/pkg/adapters/storage"
	"server/internal/models/address"
	"gorm.io/gorm"
)

type AppContainer struct {
	cfg               config.Config
	dbConn            *gorm.DB
	AuthService       *AuthService
	WalletService     *WalletService
	RestaurantService *RestaurantService
  AddressService    *AddressService
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
	app.setWalletService()
	app.setAddressService()
	app.setRestaurantService()
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

func (a *AppContainer) setWalletService() {
	if a.WalletService != nil {
		return
	}
	a.WalletService = NewWalletService(wallet.NewWalletOps(a.dbConn, storage.NewWalletRepo(a.dbConn)), creditCard.NewCreditCardOps(a.dbConn, storage.NewCreditCardRepo(a.dbConn)))
}

func (a *AppContainer) setAddressService() {
	if a.AddressService != nil {
		return
	}
	a.AddressService = NewAddressService(address.NewAddressOps(a.dbConn, storage.NewAddressRepo(a.dbConn)))
}

func (a *AppContainer) setRestaurantService() {
	if a.RestaurantService != nil {
		return
	}
	a.RestaurantService = NewRestaurantService(restaurant.NewRestaurantOps(a.dbConn, storage.NewRestaurantRepo(a.dbConn)), menu.NewMenuOps(a.dbConn, storage.NewMenuRepo(a.dbConn)))
}

