package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"goodfood-app/internal/model"
)

// RecipeRepository хранит ссылку на наше подключение к БД.
type RecipeRepository struct {
	db *DB
}

// NewRecipeRepository — конструктор репозитория.
func NewRecipeRepository(db *DB) *RecipeRepository {
	return &RecipeRepository{db: db}
}

// GetRandomByCategory — выбирает случайный рецепт из базы.
func (r *RecipeRepository) GetRandomByCategory(ctx context.Context, category string) (*model.Recipe, error) {
	// SQL-запрос: выбираем всё, фильтруем по категории, сортируем случайно, берём 1 запись
	query := `
		SELECT id, title, category, description, ingredients, instructions, 
		       kcal, protein, fat, carbs, created_at
		FROM recipes
		WHERE category = $1
		ORDER BY RANDOM()
		LIMIT 1
	`

	var recipe model.Recipe
	var ingredientsJSON []byte // PostgreSQL вернёт JSONB как байты

	// Выполняем запрос и сразу распаковываем результаты в переменные
	err := r.db.Pool.QueryRow(ctx, query, category).Scan(
		&recipe.ID, &recipe.Title, &recipe.Category, &recipe.Description,
		&ingredientsJSON, &recipe.Instructions,
		&recipe.Kcal, &recipe.Protein, &recipe.Fat, &recipe.Carbs,
		&recipe.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении рецепта: %w", err)
	}

	// Парсим JSON-ингредиенты в слайс структур
	if err := json.Unmarshal(ingredientsJSON, &recipe.Ingredients); err != nil {
		return nil, fmt.Errorf("ошибка при разборе ингредиентов: %w", err)
	}

	return &recipe, nil
}
