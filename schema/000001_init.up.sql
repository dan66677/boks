-- Создание таблицы для хранения информации о беременных женщинах
CREATE TABLE pregnant_women (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    date_of_birth DATE NOT NULL,
    due_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы для хранения информации о схватках
CREATE TABLE contractions (
    id SERIAL PRIMARY KEY,
    woman_id INT REFERENCES pregnant_women(id) ON DELETE CASCADE,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    intensity INT CHECK (intensity BETWEEN 1 AND 10),
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- Создание индексов для ускорения поиска
CREATE INDEX idx_contractions_woman_id ON contractions(woman_id);