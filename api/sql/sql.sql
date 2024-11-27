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

CREATE TABLE followers(
    id int not null,
    FOREIGN KEY (id),
    REFERENCES users(id),
    ON DELETE CASCADE,
    follower_id int not null,
    FOREIGN KEY (id),
    REFERENCES users(id),
    ON DELETE CASCADE,
    PRIMARY KEY(id, follower_id)
) ENGINE=INNODB;