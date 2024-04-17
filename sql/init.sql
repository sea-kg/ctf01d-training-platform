-- Пользователи
CREATE TABLE `users` (
    `id` int NOT NULL AUTO_INCREMENT,
    `user_name` varchar(255) NOT NULL,
    `password_hash` varchar(255) NOT NULL,
    `avatar_url` varchar(255),
    `role` enum('admin', 'player') NOT NULL,
    PRIMARY KEY (`id`)
);

-- Команды
CREATE TABLE `teams` (
    `id` int NOT NULL AUTO_INCREMENT,
    `team_name` varchar(255) NOT NULL,
    `description` varchar(255) NOT NULL,
    `university` varchar(255),
    `social_links` text,
    `avatar_url` varchar(255),
    PRIMARY KEY (`id`)
);

-- Сервисы
CREATE TABLE `services` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `author` varchar(255) NOT NULL,
    `logo_url` varchar(255),
    `description` text,
    `is_public` boolean NOT NULL,
    PRIMARY KEY (`id`)
);

-- Игры
CREATE TABLE `games` (
    `id` int NOT NULL AUTO_INCREMENT,
    `start_time` DATETIME NOT NULL,
    `end_time` DATETIME NOT NULL,
    `description` text,
    PRIMARY KEY (`id`)
);

-- Флаги
CREATE TABLE `flags` (
    `id` int NOT NULL AUTO_INCREMENT,
    `external_id` varchar(255) NOT NULL,
    `flag` varchar(255) NOT NULL,
    `service_id` int,
    `game_id` int,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`service_id`) REFERENCES `services` (`id`),
    FOREIGN KEY (`game_id`) REFERENCES `games` (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;


-- Результаты игр
CREATE TABLE `results` (
    `id` int NOT NULL AUTO_INCREMENT,
    `team_id` int,
    `game_id` int,
    `score` int NOT NULL,
    `rank` int NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`team_id`) REFERENCES `teams` (`id`),
    FOREIGN KEY (`game_id`) REFERENCES `games` (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;
