# 🍽️ GoodFood API

> **Ешь больше — худей легче.** Случайные рецепты правильного питания для каждого приёма пищи.

[![Go](https://img.shields.io/badge/Go-1.24-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-17-336791?logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![Neon](https://img.shields.io/badge/Database-Neon-FF7A00?logo=amazonaws&logoColor=white)](https://neon.tech/)
[![Docker](https://img.shields.io/badge/Docker-✅-2496ED?logo=docker&logoColor=white)](https://www.docker.com/)
[![Render](https://img.shields.io/badge/Deploy-Render-46E3B7?logo=render&logoColor=white)](https://render.com/)
[![GitHub Pages](https://img.shields.io/badge/Frontend-GitHub%20Pages-222222?logo=github&logoColor=white)](https://pages.github.com/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

---

## 🌐 Live Demo

| Что | Ссылка | Статус |
|-----|--------|--------|
| 🎨 **Фронтенд (демо)** | [▶️ Открыть приложение](https://valdrgo.github.io/GoodFoodApi/) | ✅ Работает |
| 🔧 **Backend API** | [📡 GET /api/recipe](https://goodfoodapi.onrender.com/api/recipe?category=breakfast) | ✅ Работает |
| 📚 **Swagger Docs** | [📖 Документация API](https://goodfoodapi.onrender.com/swagger/index.html) | ✅ Работает |
| 💻 **Исходный код** | [🔍 Посмотреть код](https://github.com/ValdrGo/GoodFoodApi) | 🟢 Public |

> ⚠️ **Примечание:** Бэкенд на бесплатном тарифе Render может «засыпать» после 15 мин без запросов. Первый запрос может занять ~30-60 секунд («холодный старт»).

---

## 📖 О проекте

**GoodFood** — это пет-проект для изучения бэкенд-разработки на Go, который решает простую задачу:

> *«Хочу быстро получить рецепт сбалансированного блюда без долгих поисков»*.

### 💡 Концепция
- 🎲 **Случайный рецепт** — не нужно выбирать из сотен вариантов
- 🥗 **Только ПП** — все рецепты с расчётом КБЖУ, без «мусорных» калорий
- ⚡ **Быстро** — один клик → готовый рецепт с ингредиентами и инструкцией
- 📱 **Доступно** — работает в браузере, не требует установки

### 🎯 Для кого
- Для тех, кто следит за питанием, но не хочет тратить время на поиск рецептов
- Для изучения работы с **Go + PostgreSQL + Docker + деплоем в облако**
- Для портфолио начинающего **Golang Backend Developer**

---

## ✨ Возможности

### 🔹 Для пользователя
- [ ] Выбор категории: завтрак 🍳 / обед 🍲 / ужин  / перекус 
- [ ] Мгновенный показ рецепта с:
  - 📝 Названием и описанием
  - 🥘 Списком ингредиентов с количеством
  - 👨‍ Пошаговой инструкцией
  - 📊 КБЖУ на порцию (калории, белки, жиры, углеводы)
- [ ] Адаптивный дизайн (мобильные + десктоп)
- [ ] Красивый интерфейс с анимациями

### 🔹 Для разработчика (архитектура)
- [ ] **REST API** с чистыми хендлерами
- [ ] **Слоистая архитектура**: Handler → Service → Repository → Model
- [ ] **Пул соединений** к БД через `pgxpool`
- [ ] **Миграции** базы данных (SQL-файлы)
- [ ] **CORS middleware** для кросс-доменных запросов
- [ ] **Swagger-документация** (автогенерация через `swaggo`)
- [ ] **Переменные окружения** для конфигурации (.env + prod)
- [ ] **Docker Compose** для локальной разработки

---

## 🛠 Технологический стек
┌─────────────────────────────────────────┐
│ 🖥️ Frontend │
│ • HTML5 / CSS3 / Vanilla JS │
│ • Flexbox / Grid / Animations │
│ • Deploy: GitHub Pages + Actions │
├─────────────────────────────────────────┤
│ ⚙️ Backend (Go) │
│ • Go 1.24 + stdlib (net/http) │
│ • pgx/v5 — драйвер PostgreSQL │
│ • godotenv — переменные окружения │
│ • swaggo/http-swagger — документация │
├─────────────────────────────────────────┤
│ 🗄️ Database │
│ • PostgreSQL 17 (Neon Cloud) │
│ • JSONB для хранения ингредиентов │
│ • Индексы для оптимизации запросов │
│ • SSL-соединение (sslmode=require) │
│ • Connection pooling через pgxpool │
├─────────────────────────────────────────┤
│ 🐳 DevOps │
│ • Docker + Docker Compose (локально) │
│ • Render — деплой бэкенда │
│ • GitHub Actions — деплой фронтенда │
│ • Neon — облачная БД с SSL │
└─────────────────────────────────────────┘

---

## 🏗 Архитектура проекта
goodfood-app/
├── cmd/api/
│ └── main.go # Точка входа, инициализация
├── internal/
│ ├── handler/ # HTTP-хендлеры (валидация, ответ)
│ │ └── recipe_handler.go
│ ├── service/ # Бизнес-логика
│ │ └── recipe_service.go
│ ├── repository/ # Работа с БД
│ │ ├── db.go # Подключение, миграции
│ │ └── recipe_repository.go
│ ├── model/ # Структуры данных
│ │ └── recipe.go
│ └── middleware/ # CORS, логирование
│ └── cors.go
├── frontend/
│ └── index.html # SPA: HTML+CSS+JS в одном файле
├── migrations/
│ ├── 001_update_schema.sql # Создание таблиц
│ └── 002_seed_recipes.sql # 80 рецептов (ПП)
├── docs/ # Swagger-документация (автоген)
├── docker-compose.yml # Локальный запуск БД
├── .env.example # Шаблон переменных окружения
├── go.mod / go.sum # Зависимости Go
└── README.md # Этот файл

### 🔄 Поток запроса
[Пользователь]
↓
[Frontend: fetch /api/recipe?category=breakfast]
↓
[Handler: валидация параметра, вызов сервиса]
↓
[Service: бизнес-правила, подготовка запроса]
↓
[Repository: SQL-запрос к PostgreSQL]
↓
[PostgreSQL: SELECT с ORDER BY RANDOM() LIMIT 1]
↓
[Ответ: JSON с рецептом → Фронтенд → Пользователь]

---

## 📡 API Reference

### Получить случайный рецепт
GET /api/recipe?category={category}

#### Параметры

| Параметр | Тип | Обязательный | Описание | Пример |
|----------|-----|-------------|----------|--------|
| `category` | string | ✅ Да | Категория блюда | `breakfast`, `lunch`, `dinner`, `snack` |

#### Ответ (200 OK)

```json
{
  "id": 8,
  "title": "Смузи-боул с гранолой",
  "category": "breakfast",
  "description": "Освежающий завтрак с хрустящей текстурой",
  "ingredients": [
    {"name": "Банан замороженный", "amount": "1шт"},
    {"name": "Ягоды", "amount": "100г"},
    {"name": "Йогурт греческий", "amount": "100г"},
    {"name": "Гранола без сахара", "amount": "30г"},
    {"name": "Семена чиа", "amount": "1ч.л."}
  ],
  "instructions": "1. Банан, ягоды и йогурт взбить в блендере до густоты.\n2. Перелить в миску.\n3. Сверху посыпать гранолой и семенами чиа.\n4. Украсить свежими ягодами.",
  "kcal": 310,
  "protein": 15,
  "fat": 8,
  "carbs": 48,
  "created_at": "2026-03-26T18:02:17.995259Z"
}
Коды ответов
200 ✅ Успех, рецепт возвращён
400 ❌ Неверный параметр category
404 ❌ Рецепт не найден (маловероятно)
500 ❌ Ошибка сервера / БД

🚀 Быстрый старт (локальная разработка)
Требования
Go 1.21+
Docker + Docker Compose
1. Клонировать репозиторий
bash
git clone https://github.com/ValdrGo/GoodFoodApi.git
cd GoodFoodApi
2. Настроить переменные окружения
# Скопировать шаблон
cp .env.example .env
# Отредактировать .env при необходимости:
# POSTGRES_USER=postgres
# POSTGRES_PASSWORD=secret_password
# POSTGRES_DB=goodfood_db
# DB_PORT=5433
# DB_HOST=localhost
# DB_SSLMODE=disable
3. Запустить базу данных
docker-compose up -d
# Проверить: docker-compose ps
4. Запустить сервер
go run cmd/api/main.go
# Сервер запустится на http://localhost:8080
5. Открыть документацию
👉 http://localhost:8080/swagger/index.html
6. Протестировать API
# Через curl
curl "http://localhost:8080/api/recipe?category=breakfast"
# Или через браузер
# Откройте фронтенд: ./frontend/index.html

☁️ Деплой в облако
Проект полностью развёрнут на бесплатных облачных сервисах:
🗄️ База данных Neon ✅ 80 рецептов в облаке
⚙️ Бэкенд Render ✅ goodfoodapi.onrender.com
🎨 Фронтенд GitHub Pages ✅ valdrgo.github.io/GoodFoodApi
🔄 CI/CD GitHub Actions ✅ Авто-деплой при пуше

🗄️ Настройки базы данных (Neon)
Провайдер Neon (Serverless PostgreSQL)
Регион AWS EU (Frankfurt)
Версия PostgreSQL 17
SSL sslmode=require (обязательно)
Пул соединений pgxpool с авто-реконнектом
База данных goodfood_db

🔐 Строка подключения (пример)
postgresql://user:pass@ep-xxx.eu-central-1.aws.neon.tech/goodfood_db?sslmode=require&channel_binding=require
🔐 Безопасность
✅ Строка подключения хранится в DATABASE_URL (Render Environment Variables)
✅ .env добавлен в .gitignore — креды не в репозитории
✅ Локальная разработка: отдельный .env с sslmode=disable
📦 Миграции и сиды
Данные загружаются через SQL-миграции:
bash
# Применить структуру БД
psql $DATABASE_URL -f migrations/001_update_schema.sql
# Загрузить 80 рецептов ПП
psql $DATABASE_URL -f migrations/002_seed_recipes.sql

⚙️ Настройки Render (бэкенд)
DATABASE_URL postgresql: ... Подключение к Neon с SSL
PORT 10000 Порт для Render (free tier)

🎨 Настройки GitHub Pages (фронтенд)
Workflow: .github/workflows/static.yml
Source: GitHub Actions
Path: ./frontend (только фронтенд-файлы)
Авто-деплой: при пуше в ветку main

🗓️ Roadmap
✅ Готово
Базовый API с 4 категориями
80 рецептов с КБЖУ (ПП)
Фронтенд с адаптивным дизайном
Деплой бэкенда на Render
Деплой фронтенда на GitHub Pages
Swagger-документация
Облачная БД на Neon с миграциями
🔄 В работе
Авторизация пользователей (JWT)
Личный кабинет: избранное, история
Поиск рецептов по ингредиентам
Фильтры: время приготовления, сложность
💡 Идеи на будущее
Мобильное приложение (React Native / Flutter)
Интеграция с сервисами доставки продуктов
Генерация списка покупок на неделю
Социальные функции: рейтинг, комментарии
🤝 Вклад в проект
Проект учебный, но пулл-реквесты приветствуются!
Форкните репозиторий
Создайте ветку (git checkout -b feature/AmazingFeature)
Закоммитьте изменения (git commit -m 'Add some AmazingFeature')
Запушьте (git push origin feature/AmazingFeature)
Откройте Pull Request
📄 Лицензия
Распространяется под лицензией MIT. См. файл LICENSE для деталей.
👤 Автор
Vladimir Pluzhnikov — Golang Backend Developer

Сделано с ❤️ и упорством в 2026 году