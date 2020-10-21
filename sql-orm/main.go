package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rezaahmadk/dts-Introduction-Database-1/sql-generic/config"
	"github.com/rezaahmadk/dts-Introduction-Database-1/sql-orm/database"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Println(err)
		return
	}

	db, err := initDB(cfg.Database)
	if err != nil {
		log.Println(err)
		return
	}

	// database.InsertCustomer(database.CustomerORM{
	// 	FirstName:    "Reza",
	// 	LastName:     "Ahmad Kurniawan",
	// 	NpwpId:       "npwp321",
	// 	Age:          24,
	// 	CustomerType: "Premium",
	// 	Street:       "Kebon Agung",
	// 	City:         "Sleman",
	// 	State:        "Indonesia",
	// 	ZipCode:      "55562",
	// 	PhoneNumber:  "085713997774",
	// 	AccountORM: []database.AccountORM{
	// 		{
	// 			Balance:     1000,
	// 			AccountType: "Premium",
	// 		},
	// 		{
	// 			Balance:     3000,
	// 			AccountType: "Deposito",
	// 		},
	// 	},
	// }, db)

	database.GetCustomer(db)
	//database.DeleteCustomer(1,db)
	// database.UpdateCustomer(database.CustomerORM{
	// 	FirstName: "Reza A K",
	// 	Age:       27,
	// 	City:      "Jakarta",
	// }, 1, db)
}

func getConfig() (config.Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName("config.yml")

	if err := viper.ReadInConfig(); err != nil {
		return config.Config{}, err
	}

	var cfg config.Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}

func initDB(dbConfig config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.Config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&database.CustomerORM{}, &database.AccountORM{})

	log.Println("db successfully connected")

	return db, nil
}
