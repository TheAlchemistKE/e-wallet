package service

import (
	"blog/api/repository"
	"blog/models"
)

type WalletService struct {
	repository repository.WalletRepository
}

func NewWalletService(r repository.WalletRepository) WalletService {
	return WalletService{
		repository: r,
	}
}

func (w WalletService) Create(wallet models.Wallet) error {
	return w.repository.Create(wallet)
}

func (w WalletService) TopUpWallet(wallet models.Wallet) error {
	return w.repository.TopUpWallet(wallet)
}

func (w WalletService) GetWalletBalance(phoneNumber int) float64 {
	return w.repository.GetWalletBalance(phoneNumber)
}

func (w WalletService) Find(phoneNumber int) (models.Wallet, error) {
	return w.repository.Find(phoneNumber)
}

func (w WalletService) Delete(phoneNumber int) error {
	var wallet models.Wallet
	wallet.PhoneNumber = phoneNumber
	return w.repository.Delete(wallet)

}
