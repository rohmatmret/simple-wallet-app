# Simple app wallet
feature
 - Debit
 - Credit


# How to run
- Clone 
- setup connection db 
 ```go
 func Connect() *gorm.DB  {
 	dsn := "root:root@tcp(127.0.0.1:3306)/scoopwallet?charset=utf8mb4&parseTime=True&loc=Local"
 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Panic("error connection db", err)
    }
        db.Debug()
        return db
    }
```
- run main.go