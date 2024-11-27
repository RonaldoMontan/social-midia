INSERT INTO users (name, nick, email, password)
VALUES
("user1", "@1", "user1@gmail.com", "$2a$10$YnFqJEwImS7DHb9KZuaw/OMz/gaDFr1LUL3ZZ/Rtx6qaHPGXERn9O"),
("user2", "@2", "user2@gmail.com", "$2a$10$YnFqJEwImS7DHb9KZuaw/OMz/gaDFr1LUL3ZZ/Rtx6qaHPGXERn9O"),
("user3", "@3", "user3@gmail.com", "$2a$10$YnFqJEwImS7DHb9KZuaw/OMz/gaDFr1LUL3ZZ/Rtx6qaHPGXERn9O");

INSERT INTO followers (user_id, follower_id)
VALUES
(1,2),
(1,3),
(2,1),
(3,1);
