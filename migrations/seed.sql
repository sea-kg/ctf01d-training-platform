INSERT INTO users (user_name, password_hash, role, avatar_url, status) VALUES
('user1', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/user1', 'teammate'),
('user2', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/user2', 'teammate'),
('user3', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/user3', 'teammate'),
('user4', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/user4', 'teammate'),
('user5', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/user5', 'teammate'),
('user6', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/user6', 'teammate'),
('user7', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/user7', 'teammate'),
('user8', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/user8', 'teammate'),
('user9', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/user9', 'teammate'),
('user10', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/user10', 'teammate'),
('admin', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'admin', 'https://robohash.org/admin', '');

INSERT INTO teams (name, description, university_id, social_links, avatar_url) VALUES
('Team A', 'Description A', 401, '', 'https://robohash.org/A'),
('Team B', 'Description B', 402, '', 'https://robohash.org/B'),
('Team C', 'Description C', 403, '', 'https://robohash.org/C');

INSERT INTO team_members (user_id, team_id) VALUES
(1, 4), (2, 4), (3, 4),
(4, 5), (5, 5), (6, 5),
(7, 6), (8, 6), (9, 6);

INSERT INTO games (start_time, end_time, description) VALUES
('2023-10-01 12:00:00', '2023-10-01 15:00:00', 'Game A'),
('2023-10-02 12:00:00', '2023-10-02 15:00:00', 'Game B');

INSERT INTO team_games (team_id, game_id) VALUES
(4, 1),
(5, 1),
(6, 2),
(4, 2);

INSERT INTO services (name, author, logo_url, description, is_public) VALUES
('Service A', 'Author Biba', '', 'Service Description A', TRUE),
('Service B', 'Author Boba', '', 'Service Description B', TRUE);

INSERT INTO game_services (game_id, service_id) VALUES
(1, 1),
(2, 2);
