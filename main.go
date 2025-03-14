package main

import (
	"fmt"

	"jooby/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Подключаем базу данных
	database.ConnectDB()

	// Создаем новый роутер
	router := gin.Default()

	// Добавляем тестовый маршрут
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Запускаем сервер на порту 8080
	fmt.Println("Сервер запущен на порту 8080")
	router.Run(":8080")
}
