package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectToDatabase() {
	fmt.Println("Connect To Database")
	conf, err := LoadEnvVariables()
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Db_Host, conf.Db_Port, conf.Db_Username, conf.Db_Password, conf.Db_Name)
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		fmt.Println("MASUK ERROR PAS CONNECT DATABASE NIHH")
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	DB = db
}

func SyncDB() {
	// err := DB.AutoMigrate(&models.User{})
	// if err != nil {
	// 	fmt.Printf("%#v\n", &models.User{})
	// 	fmt.Println("MASUK DI ERROR AUTO MIGRATE")
	// 	log.Fatal(err)
	// }
	// SeedDataUsers()
	// SeedDataCategory()
	// DeleteUsers()
	// DeleteCategories()
}
