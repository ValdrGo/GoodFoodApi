package main

import (
	"context"
	"goodfood-app/internal/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "goodfood-app/docs" // Добавлен импорт документации
	"goodfood-app/internal/handler"
	"goodfood-app/internal/repository"
	"goodfood-app/internal/service"
)

// @title GoodFood API
// @version 1.0
// @description API для получения случайных рецептов блюд
// @description
// @description Категории: breakfast, lunch, dinner, snack
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env не найден, используем переменные окружения")
	}

	ctx := context.Background()

	// Подключение к БД
	db, err := repository.NewDB(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Создание таблицы
	if err := db.InitDB(ctx); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Внедрение зависимостей
	recipeRepo := repository.NewRecipeRepository(db)
	recipeService := service.NewRecipeService(recipeRepo)
	recipeHandler := handler.NewRecipeHandler(recipeService)

	// ✅ Роутинг (регистрируем ВСЕ маршруты ДО запуска сервера)

	// Твой API
	http.HandleFunc("/api/recipe", middleware.CORS(recipeHandler.GetRandomRecipe))

	// Swagger UI
	http.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DeepLinking(true),
	))

	// Запуск сервера
	server := &http.Server{Addr: ":8080"}

	go func() {
		log.Println("🚀 Server starting on http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Ожидание сигнала остановки
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
}
