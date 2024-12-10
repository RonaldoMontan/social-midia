CREATE DATABASE IF NOT EXISTS devbook;

Use devbook;

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS publication;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(100) not null,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE followers (
    user_id int NOT NULL,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,
    follower_id int NOT NULL,
    FOREIGN KEY (follower_id) 
    REFERENCES users(id) 
    ON DELETE CASCADE,
    PRIMARY KEY (user_id, follower_id)
) ENGINE=INNODB;

CREATE TABLE publication (
    publication_id int auto_increment primary key,
    title varchar(50) NOT NULL,
    content varchar(320) NOT NULL,
    author_id int NOT NULL,
    FOREIGN KEY (author_id)
    REFERENCES users(id)
    ON DELETE CASCADE,    
    likes int DEFAULT 0,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;