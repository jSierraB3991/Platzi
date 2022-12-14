DROP TABLE IF EXISTS student;
CREATE TABLE student(
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    age INTEGER NOT NULL
);

DROP TABLE IF EXISTS test;
CREATE TABLE test (
    id VARCHAR(32) PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

DROP TABLE IF EXISTS questions;
CREATE TABLE questions (
    id VARCHAR(32) PRIMARY KEY,
    test_id VARCHAR(32) NOT NULL,
    question VARCHAR(255) NOT NULL,
    answer VARCHAR(255) NOT NULL,
    FOREIGN KEY (test_id) REFERENCES test(id)
);

DROP TABLE IF EXISTS enrollment;
CREATE TABLE enrollment (
    student_id VARCHAR(32),
    test_id VARCHAR(32),
    FOREIGN KEY (test_id) REFERENCES test(id),
    FOREIGN KEY (student_id) REFERENCES student(id)
);
