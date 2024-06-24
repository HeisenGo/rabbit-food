package storage

import (
	"context"
	"gorm.io/gorm"
	creditCard "server/internal/models/wallet/credit_card"
	wallet "server/internal/models/wallet/wallet_credit_card"
	"server/pkg/adapters/storage/entities"
	"server/pkg/adapters/storage/mappers"
	"server/pkg/utils"
)

type creditCardRepo struct {
	db *gorm.DB
}

func NewCreditCardRepo(db *gorm.DB) creditCard.Repo {
	return &creditCardRepo{
		db: db,
	}
}
func (r *creditCardRepo) CreateCardAndAddToWallet(ctx context.Context, creditCard *creditCard.CreditCard) (*creditCard.CreditCard, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if rv := recover(); rv != nil {
			tx.Rollback()
		}
	}()

	var userWalletEntity *entities.Wallet
	userID, err := utils.GetUserIDFromContext(ctx)
	if err = tx.Where("user_id = ?", userID).First(&userWalletEntity).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	newCreditCard := mappers.CreditCardDomainToEntity(creditCard)
	if err = tx.FirstOrCreate(&newCreditCard).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	walletCreditCardEntity := wallet.NewWalletCreditCard(userWalletEntity.ID, newCreditCard.ID)
	if err = tx.Create(&walletCreditCardEntity).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err = tx.Commit().Error; err != nil {
		return nil, err
	}

	createdCreditCard := mappers.CreditCardEntityToDomain(newCreditCard)
	return createdCreditCard, nil
}
