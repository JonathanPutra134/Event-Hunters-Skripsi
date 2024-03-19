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
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}
func ResetSerial(tableName string) error {
	_, err := DB.Exec(fmt.Sprintf("SELECT setval(pg_get_serial_sequence('%s', 'id'), 1, false)", tableName))
	return err
}

func SyncDB() {
	//UNCOMMENT FOR PRODUCTION
	//RegenerateAllScenarioForRecommendation()
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
