CREATE TABLE movie (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name varchar(50) NOT NULL,
    minutes int NOT NULL,
    genre varchar(50) NOT NULL
);

INSERT INTO movie (name, minutes, genre)
VALUES('Dark Knight', 150, 'ACTION'),
      ('Memento', 150, 'THRILLER'),
      ('There Something About Mary', 119, 'COMEDY'),
      ('Super 8', 112, 'THRILLER'),
      ('Scream', 150, 'HORROR'),
      ('Home Alone', 150, 'COMEDY'),
      ('Matrix', 136, 'ACTION');