package repository

import (
	"blog/infrastructure"
	"blog/models"
)

type WalletRepository struct {
	db infrastructure.Database
}

func NewWalletRepository(db infrastructure.Database) WalletRepository {
	return WalletRepository{
		db: db,
	}
}

func (w WalletRepository) Create(wallet models.Wallet) error {
	return w.db.DB.Create(&wallet).Error
}

func (w WalletRepository) GetWalletBalance(phoneNumber int) float64 {
	var wallet models.Wallet = models.Wallet{PhoneNumber: phoneNumber}
	result := &wallet
	return result.Balance
}

func (w WalletRepository) TopUpWallet(wallet models.Wallet) error {
	return w.db.DB.Save(&wallet).Error
}

func (w WalletRepository) Find(phoneNumber int) (models.Wallet, error) {
	var wallet models.Wallet
	err := w.db.DB.Model(models.Wallet{PhoneNumber: phoneNumber}).First(&wallet).Error
	return wallet, err
}

func (w WalletRepository) Delete(wallet models.Wallet) error {
	return w.db.DB.Delete(&wallet).Error
}
