package service

import (
	"context"
	"fmt"

	"goodfood-app/internal/model"
	"goodfood-app/internal/repository"
)

// RecipeService хранит ссылку на репозиторий (базу данных).
type RecipeService struct {
	repo *repository.RecipeRepository
}

// NewRecipeService — конструктор. Создаёт новый сервис.
func NewRecipeService(repo *repository.RecipeRepository) *RecipeService {
	return &RecipeService{repo: repo}
}

// GetRandomRecipe — главная функция: получить случайный рецепт по категории.
func (s *RecipeService) GetRandomRecipe(ctx context.Context, category string) (*model.Recipe, error) {
	// 1. Валидация: проверяем, что категория допустима
	validCategories := map[string]bool{
		"breakfast": true,
		"lunch":     true,
		"dinner":    true,
		"snack":     true,
	}

	if !validCategories[category] {
		return nil, fmt.Errorf("неверная категория: %s", category)
	}

	// 2. Делегируем запрос к базе данных репозиторию
	recipe, err := s.repo.GetRandomByCategory(ctx, category)
	if err != nil {
		return nil, fmt.Errorf("ошибка в сервисе: %w", err)
	}

	return recipe, nil
}
