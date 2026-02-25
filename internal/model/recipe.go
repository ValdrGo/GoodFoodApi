package model

import "time"

// Recipe — это структура, которая описывает один рецепт.
// Теги `json:"..."` нужны, чтобы при отправке клиенту поля назывались правильно.
type Recipe struct {
	ID           int          `json:"id"`
	Title        string       `json:"title"`
	Category     string       `json:"category"`
	Description  string       `json:"description"`
	Ingredients  []Ingredient `json:"ingredients"` // Слайс (массив) ингредиентов
	Instructions string       `json:"instructions"`
	Kcal         int          `json:"kcal"`
	Protein      int          `json:"protein"`
	Fat          int          `json:"fat"`
	Carbs        int          `json:"carbs"`
	CreatedAt    time.Time    `json:"created_at"`
}

// Ingredient — отдельная структура для ингредиента.
type Ingredient struct {
	Name   string `json:"name"`   // Название: "Молоко"
	Amount string `json:"amount"` // Количество: "200мл"
}
