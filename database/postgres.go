package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectDB() {
	var err error
	url := "postgresql://jooby_user:jooby_pass@localhost:5432/jooby_db?sslmode=disable"

	DB, err = pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	fmt.Println("✅ Успешное подключение к PostgreSQL!")
}
