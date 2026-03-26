package middleware

import "net/http"

// CORS добавляет заголовки для разрешения кросс-доменных запросов
func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Разрешаем запросы с любых доменов (для портфолио)
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Разрешённые методы
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Разрешённые заголовки
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Обрабатываем preflight-запрос (браузер спрашивает разрешение)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Вызываем следующий хендлер
		next(w, r)
	}
}
