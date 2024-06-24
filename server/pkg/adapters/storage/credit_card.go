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
	if err = tx.Where("number = ?", newCreditCard.Number).First(&newCreditCard).Error; err != nil {
		if err = tx.Create(&newCreditCard).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
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

func (r *creditCardRepo) GetUserWalletCards(ctx context.Context) ([]*creditCard.CreditCard, error) {
	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	var creditCardEntities []*entities.CreditCard

	err = r.db.Joins("JOIN wallet_credit_cards ON wallet_credit_cards.credit_card_id = credit_cards.id").
		Joins("JOIN wallets ON wallets.id = wallet_credit_cards.wallet_id").
		Where("wallets.user_id = ?", userID).
		Find(&creditCardEntities).Error

	if err != nil {
		return nil, err
	}
	allDomainCards := mappers.BatchCreditCardEntityToDomain(creditCardEntities)
	return allDomainCards, nil
}
