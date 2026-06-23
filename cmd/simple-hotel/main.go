package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/Lalka12235/simple-hotel.git/internal/handler"
	"github.com/Lalka12235/simple-hotel.git/internal/repository"
)

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" { dbHost = "localhost" }
	if dbPort == "" { dbPort = "5432" }
	if dbUser == "" { dbUser = "postgres" }
	if dbPass == "" { dbPass = "postgres" }
	if dbName == "" { dbName = "hotel_db" }

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", 
		dbUser, dbPass, dbHost, dbPort, dbName)
	
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Ошибка конфигурации БД: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Не удалось подключиться к PostgreSQL: %v", err)
	}
	fmt.Println("Успешно подключено к PostgreSQL!")

	clientRepo := repository.NewClientRepository(db)
	categoryRepo := repository.NewRoomCategoryRepository(db)
	roomRepo := repository.NewRoomRepository(db)
	bookingRepo := repository.NewBookingRepository(db)

	clientHandler := handler.NewClientHandler(clientRepo)
	categoryHandler := handler.NewRoomCategoryHandler(categoryRepo)
	roomHandler := handler.NewRoomHandler(roomRepo)
	bookingHandler := handler.NewBookingHandler(bookingRepo)

	mux := http.NewServeMux()
	
	mux.Handle("/api/clients", clientHandler)
	mux.Handle("/api/clients/", clientHandler)

	mux.Handle("/api/categories", categoryHandler)
	mux.Handle("/api/categories/", categoryHandler)

	mux.Handle("/api/rooms", roomHandler)
	mux.Handle("/api/rooms/", roomHandler)

	mux.Handle("/api/bookings", bookingHandler)
	mux.Handle("/api/bookings/", bookingHandler)

	port := ":8080"
	fmt.Printf("Сервер отеля запущен на http://localhost%s\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}