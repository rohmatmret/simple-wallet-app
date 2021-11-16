package wallet

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type RepositoryWallet interface {
	creditBalance() error
	debitBalance() error
}
type repository struct {
	db gorm.DB
}

func NewRepository(db gorm.DB) *repository {
	return &repository{db: db}
}
type Wallet struct {
	UserId string `json:"user_id"`
	Balance int64 `json:"balance"`
}
type Wallet_transactions struct {
	UserId string `json:"user_id"`
	Amount int64 `json:"amount"`
	Status string `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewwalletTransactions(userId string, amount int64, status string) *Wallet_transactions {
	return &Wallet_transactions{UserId: userId, Amount: amount, Status: status}
}


func (d *repository) CreditBalance(amount int64,UserId string) error {
	w := Wallet{Balance: amount,UserId: UserId}
	err := d.db.Create(&w).Error
	return err
}

func (d *repository) ValidateBalance(amount int64, userID string) bool {
	var b Wallet
	findBalance := d.db.Debug().Where(&Wallet{UserId: userID}).First(&b).Error
	if findBalance != nil {
		log.Println("balance not found ",userID)
		return false
	}

	if b.Balance >= amount {
		return true
	} else {
		return false
	}

}

func (d repository) CreateTransactionWallet(amount int64,userId string) error {
	tx := d.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	wt := NewwalletTransactions(userId,amount, "credit")

	if err := tx.Create(&wt).Error; err != nil {
		tx.Rollback()
		return err
	}

	w := Wallet{Balance: amount,UserId: userId}

	if err := tx.Create(&w).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (d repository) Debit(userId string, amount int64) error  {
	var w  Wallet
	tx := d.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	wt := NewwalletTransactions(userId,amount, "debit")

	if err := tx.Create(&wt).Error; err != nil {
		tx.Rollback()
		return err
	}

	err := tx.Where(&Wallet{UserId: userId}).Find(&w).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	w.Balance -= amount
	tx.Save(&w)
	if err := tx.Create(&w).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}