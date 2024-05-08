package crud_functions

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db_connection *gorm.DB

func init() {
	godotenv.Load()
	dsn_reco_client_reader := "host=" + os.Getenv("DB_HOST_READER") + " port=" + os.Getenv("DB_PORT") + " user=" + os.Getenv("DB_USER") + " dbname=" + os.Getenv("DB_NAME") + " password=" + os.Getenv("DB_PASSWORD") + " sslmode=disable"

	//dsn := "host=localhost port=5432 user=postgres dbname=notification password=54321 sslmode=disable"
	//fmt.Printf("dsn: %v\n", dsn)
	Reco_client_db_reader, err := gorm.Open(postgres.Open(dsn_reco_client_reader), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Db_connection = Reco_client_db_reader
}
