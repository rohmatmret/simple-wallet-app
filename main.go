package main

import (
	"github.com/google/uuid"
	db2 "github.com/scoop-wallet/src/db"
	"github.com/scoop-wallet/src/users"
	"github.com/scoop-wallet/src/wallet"
	"log"
)

func main()  {
	db := db2.Connect()
	db2.Migrate(db) //create table
	id := uuid.New().String() // create uuid for primary userid
	user := users.NewUsers(id,"haidir saha ") // setter user model

	s := users.NewRepository(*db).Save(*user) //save users
	if s != nil {
		log.Panic("errors create users",s)
	} else {
		log.Println("create transactions wallet")
		newWallet := wallet.NewRepository(*db).CreateTransactionWallet(2200,id)

		if newWallet != nil {
			log.Panic("error create transactions",newWallet)
		} else {
			// call debit
			log.Println("process debit")
			allowed := wallet.NewRepository(*db).ValidateBalance(2000,id)
			if allowed {

				debit := wallet.NewRepository(*db).Debit(id,2000)
				if debit != nil {
					log.Println(debit)
				} else {
					log.Println("success debit")
				}
			} else {
				log.Println("yo not allow to debit ",id)
			}

		}
	}
}
