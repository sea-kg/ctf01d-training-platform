INSERT INTO users (user_name, password_hash, role, avatar_url, status) VALUES
('Neo', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/neo', 'active'),
('Morpheus', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/morpheus', 'active'),
('Trinity', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/trinity', 'active'),
('Cipher', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/cipher', 'active'),
('Seraph', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/seraph', 'active'),
('Smith', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/smith', 'inactive'),
('Oracle', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/oracle', 'active'),
('Sati', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/sati', 'active'),
('Apoc', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'player', 'https://robohash.org/apoc', 'active'),
('Dozer', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'admin', 'https://robohash.org/dozer', '');

INSERT INTO teams (name, description, university_id, social_links, avatar_url) VALUES
('HackersX', 'Specialized in network attacks and defense', 401, '', 'https://robohash.org/hackersx'),
('CodeRed', 'Expert in cryptography and steganography', 402, '', 'https://robohash.org/codered'),
('NullByte', 'Skilled in web security and binary exploitation', 403, '', 'https://robohash.org/nullbyte');

INSERT INTO team_members (user_id, team_id) VALUES
(1, 1), (2, 1), (3, 1),
(4, 2), (5, 2), (6, 2),
(7, 3), (8, 3), (9, 3);

INSERT INTO games (start_time, end_time, description) VALUES
('2023-10-01 12:00:00', '2023-10-01 15:00:00', 'Capture the Network Flags'),
('2023-10-02 12:00:00', '2023-10-02 15:00:00', 'Decrypt the Hidden Messages');

INSERT INTO team_games (team_id, game_id) VALUES
(1, 1), (2, 1), (3, 2), (1, 2);

INSERT INTO services (name, author, logo_url, description, is_public) VALUES
('NetAttack', 'Phantom', '', 'Simulated network attack platform', TRUE),
('CryptoBox', 'Enigma', '', 'Cryptography challenge service', TRUE);

INSERT INTO game_services (game_id, service_id) VALUES
(1, 1), (2, 2);
