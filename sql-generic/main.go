package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rezaahmadk/dts-Introduction-Database-1/sql-generic/config"
	"github.com/rezaahmadk/dts-Introduction-Database-1/sql-generic/database"
	"github.com/spf13/viper"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Println(err)
		return
	}

	db, err := connect(cfg.Database)
	if err != nil {
		log.Println(err)
		return
	}

	// database.InsertCustomer(database.Customer{
	// 	FirstName:    "Reza",
	// 	LastName:     "Ahmad Kurniawan",
	// 	NpwpId:       "npwp321",
	// 	Age:          24,
	// 	CustomerType: "Free",
	// 	Street:       "Kebon Agung",
	// 	City:         "Sleman",
	// 	State:        "Indonesia",
	// 	ZipCode:      "55562",
	// 	PhoneNumber:  "085713997774",
	// }, db)

	database.GetCustomers(db)

	//database.UpdateCustomer(25, 1, db)

	//database.DeleteCustomer(2, db)
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

func connect(cfg config.Database) (*sql.DB, error) {
	db, err := sql.Open(cfg.Driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName, cfg.Config))
	if err != nil {
		return nil, err
	}

	log.Println("db successfully connected")
	return db, nil
}
