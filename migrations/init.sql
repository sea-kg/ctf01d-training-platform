-- Пользователи
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    avatar_url VARCHAR(255),
    status VARCHAR(255),
    role VARCHAR(255) NOT NULL CHECK (role IN ('admin', 'player', 'guest'))
);

-- Университеты
CREATE TABLE universities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL
);

-- Команды
CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description VARCHAR(255) NOT NULL,
    university_id INTEGER REFERENCES universities(id),
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

-- Пользователи и команды
CREATE TABLE team_members (
    user_id INTEGER REFERENCES users(id),
    team_id INTEGER REFERENCES teams(id),
    PRIMARY KEY (user_id, team_id)
);

-- Команды и игры
CREATE TABLE team_games (
    team_id INTEGER REFERENCES teams(id),
    game_id INTEGER REFERENCES games(id),
    PRIMARY KEY (team_id, game_id)
);

-- Игры и сервисы
CREATE TABLE game_services (
    game_id INTEGER REFERENCES games(id),
    service_id INTEGER REFERENCES services(id),
    PRIMARY KEY (game_id, service_id)
);
