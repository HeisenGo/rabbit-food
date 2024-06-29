package storage

import (
	"context"
	"errors"
	"gorm.io/gorm"
	creditCard "server/internal/models/wallet/credit_card"
	"server/internal/models/wallet/wallet"
	"server/pkg/adapters/storage/entities"
	"server/pkg/adapters/storage/mappers"
	"server/pkg/utils"
)

type walletRepo struct {
	db *gorm.DB
}

func NewWalletRepo(db *gorm.DB) wallet.Repo {
	return &walletRepo{
		db: db,
	}
}
func (r *walletRepo) Create(ctx context.Context, wallet *wallet.Wallet) (*wallet.Wallet, error) {
	newWallet := mappers.WalletDomainToEntity(wallet)
	err := r.db.Create(&newWallet).Error
	if err != nil {
		return nil, err
	}
	createdWallet := mappers.WalletEntityToDomain(newWallet)
	return createdWallet, nil
}

func (r *walletRepo) Deposit(ctx context.Context, card *creditCard.CreditCard, amount uint) (*wallet.Wallet, error) {
	var userWalletEntity *entities.Wallet
	var cardEntity *entities.CreditCard

	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if rv := recover(); rv != nil {
			tx.Rollback()
		}
	}()

	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if err = tx.Where("user_id = ?", userID).First(&userWalletEntity).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Check if the credit card exists and belongs to the user's wallet
	if err = tx.Joins("JOIN wallet_credit_cards ON wallet_credit_cards.credit_card_id = credit_cards.id").
		Joins("JOIN wallets ON wallets.id = wallet_credit_cards.wallet_id").
		Where("credit_cards.number = ? AND wallets.user_id = ?", card.Number, userID).
		First(&cardEntity).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Increase the wallet balance
	userWalletEntity.Balance += amount
	if err = tx.Save(&userWalletEntity).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		return nil, err
	}
	createdWallet := mappers.WalletEntityToDomain(userWalletEntity)
	return createdWallet, nil
}

func (r *walletRepo) Withdraw(ctx context.Context, card *creditCard.CreditCard, amount uint) (*wallet.Wallet, error) {
	var userWalletEntity *entities.Wallet
	var cardEntity *entities.CreditCard

	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if rv := recover(); rv != nil {
			tx.Rollback()
		}
	}()

	userID, err := utils.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if err = tx.Where("user_id = ?", userID).First(&userWalletEntity).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Check if the credit card exists and belongs to the user's wallet
	if err = tx.Joins("JOIN wallet_credit_cards ON wallet_credit_cards.credit_card_id = credit_cards.id").
		Joins("JOIN wallets ON wallets.id = wallet_credit_cards.wallet_id").
		Where("credit_cards.number = ? AND wallets.user_id = ?", card.Number, userID).
		First(&cardEntity).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if userWalletEntity.Balance < amount {
		return nil, errors.New("not enough balance to withdraw")
	}
	userWalletEntity.Balance -= amount
	if err = tx.Save(&userWalletEntity).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		return nil, err
	}
	createdWallet := mappers.WalletEntityToDomain(userWalletEntity)
	return createdWallet, nil
}

func (r *walletRepo) GetWallet(ctx context.Context) (*wallet.Wallet, error) {
	userID, err := utils.GetUserIDFromContext(ctx)

	var userWalletEntity *entities.Wallet
	err = r.db.Where("user_id = ?", userID).First(&userWalletEntity).Error
	if err != nil {
		return nil, err
	}
	var fetchedWalletEntity *entities.Wallet
	err = r.db.WithContext(ctx).Model(&entities.Wallet{}).Where("id = ?", userWalletEntity.ID).First(&fetchedWalletEntity).Error
	if err != nil {
		return nil, err
	}
	fetchedWalletDomain := mappers.WalletEntityToDomain(fetchedWalletEntity)
	return fetchedWalletDomain, nil
}
