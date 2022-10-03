CREATE TABLE movie (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name varchar(50) NOT NULL,
    minutes int NOT NULL,
    genre varchar(50) NOT NULL,
    director varchar(50) NOT NULL
);

INSERT INTO movie (name, minutes, genre, director)
VALUES('Dark Knight', 150, 'ACTION', 'Steven'),
      ('Memento', 150, 'THRILLER', 'Spielberg'),
      ('There Something About Mary', 119, 'COMEDY', 'Cano'),
      ('Super 8', 112, 'THRILLER', 'Cesar'),
      ('Superman', 150, 'HORROR', 'Juanes'),
      ('Scream', 150, 'HORROR', 'Caicedo'),
      ('Home Alone', 150, 'COMEDY', 'Dominic'),
      ('Matrix', 136, 'ACTION', 'Duber');