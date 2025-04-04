// cmd/server/main.go
package main

import (
	"database/sql"
	"log"
	"net/http"
	"ticket-service/internal/delivery/http_handler"
	"ticket-service/internal/infrastructure/config"
	"ticket-service/internal/infrastructure/db"
	"ticket-service/internal/usecase"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	connStr := "host=" + cfg.DBHost +
		" port=" + cfg.DBPort +
		" user=" + cfg.DBUser +
		" password=" + cfg.DBPassword +
		" dbname=" + cfg.DBName +
		" sslmode=disable"

	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	dbConn.SetMaxOpenConns(25)
	dbConn.SetMaxIdleConns(25)
	dbConn.SetConnMaxLifetime(5 * time.Minute)

	err = dbConn.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}

	defer dbConn.Close()

	repo := db.NewPostgresRepo(dbConn)
	postUC := usecase.NewPostUseCase(repo)
	prePostUC := usecase.NewPrePostUseCase(repo)
	handler := http_handler.NewHandler(postUC, prePostUC, cfg.AdminLogin, cfg.AdminPass)

	http.HandleFunc("/posts", http_handler.CORSMiddleware(handler.AuthMiddleware(handler.GetPosts)))
	http.HandleFunc("/pre_posts", http_handler.CORSMiddleware(handler.AuthMiddleware(handler.GetPrePosts)))

	log.Printf("Server started on :%s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, nil))
}
