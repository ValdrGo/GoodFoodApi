-- Удаляем старую таблицу (в пет-проекте можно)
DROP TABLE IF EXISTS recipes;

-- Создаём новую с расширенной структурой
CREATE TABLE recipes (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    category VARCHAR(50) NOT NULL,           -- breakfast, lunch, dinner, snack
    cuisine_type VARCHAR(50) DEFAULT 'ПП',   -- ПП, восточная, русская и т.д. (на будущее)
    description TEXT,
    ingredients JSONB NOT NULL,
    instructions TEXT NOT NULL,
    cooking_time INT DEFAULT 0,              -- время приготовления в минутах
    difficulty VARCHAR(20) DEFAULT 'easy',   -- easy, medium, hard (на будущее)
    kcal INT NOT NULL,
    protein INT NOT NULL,
    fat INT NOT NULL,
    carbs INT NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,          -- можно скрывать рецепты
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Индексы для быстрого поиска
CREATE INDEX idx_recipes_category ON recipes(category);
CREATE INDEX idx_recipes_cuisine ON recipes(cuisine_type);
CREATE INDEX idx_recipes_active ON recipes(is_active);

-- Комментарий к таблице
COMMENT ON TABLE recipes IS 'Рецепты правильного питания для похудения';