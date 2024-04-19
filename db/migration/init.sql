-- Пользователи
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    avatar_url VARCHAR(255),
    role VARCHAR(255) NOT NULL CHECK (role IN ('admin', 'player'))
);

-- Команды
CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    team_name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    university VARCHAR(255),
    social_links TEXT,
    avatar_url VARCHAR(255)
);

-- Сервисы
CREATE TABLE services (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    logo_url VARCHAR(255),
    description TEXT,
    is_public BOOLEAN NOT NULL
);

-- Игры
CREATE TABLE games (
    id SERIAL PRIMARY KEY,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    description TEXT
);

-- Результаты игр
CREATE TABLE results (
    id SERIAL PRIMARY KEY,
    team_id INTEGER REFERENCES teams(id),
    game_id INTEGER REFERENCES games(id),
    score INTEGER NOT NULL,
    rank INTEGER NOT NULL
);
