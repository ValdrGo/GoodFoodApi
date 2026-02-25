package handler

import (
	"encoding/json"
	"net/http"

	"goodfood-app/internal/service"
)

// RecipeHandler хранит ссылку на сервис.
type RecipeHandler struct {
	service *service.RecipeService
}

// NewRecipeHandler — конструктор хендлера.
func NewRecipeHandler(service *service.RecipeService) *RecipeHandler {
	return &RecipeHandler{service: service}
}

// GetRandomRecipe — обработчик HTTP-запроса.
// Эта функция будет вызвана, когда пользователь откроет ссылку в браузере.
// GetRandomRecipe godoc
// @Summary Получить случайный рецепт
// @Description Возвращает случайный рецепт из выбранной категории (breakfast, lunch, dinner, snack)
// @Tags recipes
// @Accept json
// @Produce json
// @Param category query string true "Категория блюда" Enums(breakfast, lunch, dinner, snack)
// @Success 200 {object} model.Recipe
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/recipe [get]
func (h *RecipeHandler) GetRandomRecipe(w http.ResponseWriter, r *http.Request) {
	// 1. Получаем параметр category из URL (?category=breakfast)
	category := r.URL.Query().Get("category")
	if category == "" {
		// Если параметра нет — возвращаем ошибку 400
		http.Error(w, `{"error": "параметр category обязателен"}`, http.StatusBadRequest)
		return
	}

	// 2. Вызываем бизнес-логику (сервис)
	recipe, err := h.service.GetRandomRecipe(r.Context(), category)
	if err != nil {
		// Если ошибка — возвращаем 500
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	// 3. Отдаём результат в формате JSON
	w.Header().Set("Content-Type", "application/json") // Говорим браузеру, что это JSON
	json.NewEncoder(w).Encode(recipe)                  // Превращаем структуру в JSON и отправляем
}
