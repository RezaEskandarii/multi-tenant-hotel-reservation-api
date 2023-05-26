package domain_services

import (
	"context"
	"github.com/shopspring/decimal"
	. "reservation-api/internal/models"
	. "reservation-api/internal/repositories"
	"sync"
)

type WalletService struct {
	repo *WalletRepository
	mu   sync.Mutex
}

func NewWalletService(repo *WalletRepository) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) Create(ctx context.Context, wallet *Wallet) (*Wallet, error) {
	return s.repo.CreateWallet(ctx, wallet)
}

func (s *WalletService) Deposit(ctx context.Context, walletID uint64, amount decimal.Decimal) (*Wallet, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.repo.Deposit(ctx, walletID, amount)
}

func (s *WalletService) Withdraw(ctx context.Context, walletID uint64, amount decimal.Decimal) (*Wallet, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.repo.Withdraw(ctx, walletID, amount)
}

func (s *WalletService) GetWalletByID(ctx context.Context, walletID uint64) (*Wallet, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.repo.GetWalletByID(ctx, walletID)
}
