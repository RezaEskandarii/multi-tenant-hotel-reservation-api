package repositories

import (
	"context"
	"errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	. "reservation-api/internal/models"
	"reservation-api/pkg/multi_tenancy_database/tenant_database_resolver"
)

type WalletRepository struct {
	DbResolver *tenant_database_resolver.TenantDatabaseResolver
}

func NewWalletRepository(resolver *tenant_database_resolver.TenantDatabaseResolver) *WalletRepository {
	return &WalletRepository{resolver}
}

func (r *WalletRepository) CreateWallet(ctx context.Context, wallet *Wallet) (*Wallet, error) {
	db := r.DbResolver.GetTenantDB(ctx)
	if err := db.Create(wallet).Error; err != nil {
		return nil, err
	}
	return wallet, nil
}

func (r *WalletRepository) UpdateWallet(ctx context.Context, wallet *Wallet) (*Wallet, error) {
	db := r.DbResolver.GetTenantDB(ctx)
	if err := db.Model(wallet).Updates(wallet).Error; err != nil {
		return nil, err
	}
	return wallet, nil
}

func (r *WalletRepository) GetWalletByID(ctx context.Context, walletID uint) (*Wallet, error) {
	db := r.DbResolver.GetTenantDB(ctx)
	var wallet Wallet
	if err := db.First(&wallet, walletID).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *WalletRepository) Deposit(ctx context.Context, walletID uint, amount decimal.Decimal) (*Wallet, error) {
	db := r.DbResolver.GetTenantDB(ctx)
	if err := db.Transaction(func(tx *gorm.DB) error {
		var wallet Wallet
		if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&wallet, walletID).Error; err != nil {
			return err
		}
		wallet.Balance = wallet.Balance.Add(amount)
		if err := tx.Model(&wallet).Updates(wallet).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return r.GetWalletByID(ctx, walletID)
}

func (r *WalletRepository) Withdraw(ctx context.Context, walletID uint, amount decimal.Decimal) (*Wallet, error) {
	db := r.DbResolver.GetTenantDB(ctx)
	if err := db.Transaction(func(tx *gorm.DB) error {
		var wallet Wallet
		if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&wallet, walletID).Error; err != nil {
			return err
		}
		if wallet.Balance.LessThan(amount) {
			return errors.New("insufficient balance")
		}
		wallet.Balance = wallet.Balance.Sub(amount)
		if err := tx.Model(&wallet).Updates(wallet).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return r.GetWalletByID(ctx, walletID)
}
