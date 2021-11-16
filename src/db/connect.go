package db

import (
	"fmt"
	"github.com/scoop-wallet/src/users"
	"github.com/scoop-wallet/src/wallet"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Connect() *gorm.DB  {
	dsn := "root:root@tcp(127.0.0.1:3306)/scoopwallet?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("error connection db", err)
	}
	db.Debug()
	return db
}

func Migrate(db *gorm.DB)  {
	var Wallet wallet.Wallet
	var Users users.Users
	var Wallet_transactions wallet.Wallet_transactions
	err := db.AutoMigrate(&Wallet, &Users,&Wallet_transactions)
	if err != nil {
		log.Fatal("failed migrate", err)
	}
	fmt.Println("success migrate")
}