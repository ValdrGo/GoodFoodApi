package repository

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DB — наша обёртка над пулом соединений.
type DB struct {
	Pool *pgxpool.Pool
}

// NewDB создаёт новое подключение к базе данных.
func NewDB(ctx context.Context) (*DB, error) {
	// Читаем настройки из переменных окружения (.env)
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := "localhost"
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("POSTGRES_DB")

	// Формируем строку подключения (DSN)
	connStr := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user, password, host, port, dbName,
	)

	// Создаём пул соединений
	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания пула: %w", err)
	}

	// Проверяем, что база действительно отвечает
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ошибка подключения к базе: %w", err)
	}

	log.Println("✅ Подключение к базе данных успешно!")
	return &DB{Pool: pool}, nil
}

// InitDB создаёт таблицы, если их ещё нет.
// Обрати внимание: (db *DB) — это привязка метода к структуре.
// А InitDB с большой буквы — делает метод видимым снаружи.
func (db *DB) InitDB(ctx context.Context) error {
	query := `
	CREATE TABLE IF NOT EXISTS recipes (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		category VARCHAR(50) NOT NULL,
		description TEXT,
		ingredients JSONB NOT NULL,
		instructions TEXT NOT NULL,
		kcal INTEGER DEFAULT 0,
		protein INTEGER DEFAULT 0,
		fat INTEGER DEFAULT 0,
		carbs INTEGER DEFAULT 0,
		created_at TIMESTAMP DEFAULT NOW()
	);

	CREATE INDEX IF NOT EXISTS idx_recipes_category ON recipes(category);
	`

	_, err := db.Pool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("ошибка создания таблиц: %w", err)
	}

	log.Println("✅ Таблицы базы данных инициализированы!")
	return nil
}

// Close закрывает все соединения в пуле.
func (db *DB) Close() {
	if db.Pool != nil {
		db.Pool.Close()
		log.Println("🔌 Соединение с базой закрыто")
	}
}
