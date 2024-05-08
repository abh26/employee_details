package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		h.ServeHTTP(w, r)
	})
}
func main() {
	godotenv.Load()
	dsn_reco_client_reader := "host=" + os.Getenv("DB_HOST_READER") + " port=" + os.Getenv("DB_PORT") + " user=" + os.Getenv("DB_USER") + " dbname=" + os.Getenv("DB_NAME") + " password=" + os.Getenv("DB_PASSWORD") + " sslmode=disable"
	dsn_reco_client_writer := "host=" + os.Getenv("DB_HOST_WRITER") + " port=" + os.Getenv("DB_PORT") + " user=" + os.Getenv("DB_USER") + " dbname=" + os.Getenv("DB_NAME") + " password=" + os.Getenv("DB_PASSWORD") + " sslmode=disable"

	//dsn := "host=localhost port=5432 user=postgres dbname=notification password=54321 sslmode=disable"
	//fmt.Printf("dsn: %v\n", dsn)
	client_db_reader, err := gorm.Open(postgres.Open(dsn_reco_client_reader), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	client_db_writer, errwriter := gorm.Open(postgres.Open(dsn_reco_client_writer), &gorm.Config{})
	if errwriter != nil {
		log.Fatal(errwriter)
	}
	fmt.Println(client_db_reader, client_db_writer)
	mux := http.NewServeMux()
	//mux.Handle("/", router)
	mux.HandleFunc("/api/rs/v1/health", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Service is running"})
	})

	s := grpc.NewServer(grpc.MaxSendMsgSize(int(50e9)))

	listener, _ := net.Listen("tcp", "0.0.0.0:50051")
	listener2, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatal(err)
	}
	go s.Serve(listener)
	go http.Serve(listener2, cors(mux))
	log.Print("server started")
	for {
		time.Sleep(10000 * time.Second)
	}
	//var crud_functions.Server()
}
