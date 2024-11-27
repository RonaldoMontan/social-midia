CREATE DATABASE IF NOT EXISTS devbook;
Use devbook;
DROP TABLE IF EXISTS users;
CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(100) not null,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE followers (
    id INT NOT NULL,
    follower_id INT NOT NULL,
    PRIMARY KEY (id, follower_id),
    FOREIGN KEY (id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=INNODB;