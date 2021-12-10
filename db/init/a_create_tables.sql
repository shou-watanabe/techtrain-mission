CREATE DATABASE IF NOT EXISTS go_database;

USE go_database;

CREATE TABLE IF NOT EXISTS users (
    id int UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(120) NOT NULL,
    token VARCHAR(36) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS rarities (
    id int UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    rarity VARCHAR(3) NOT NULL,
    probability int UNSIGNED NOT NULL
);

CREATE TABLE IF NOT EXISTS characters (
    id int UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(120) NOT NULL,
    rarity_id int UNSIGNED NOT NULL,
    FOREIGN KEY (rarity_id) REFERENCES rarities(id) ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS user_character_possessions (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id int UNSIGNED NOT NULL,
    character_id int UNSIGNED NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (character_id) REFERENCES characters(id) ON DELETE RESTRICT ON UPDATE CASCADE
);
