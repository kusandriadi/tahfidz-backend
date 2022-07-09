CREATE TABLE IF NOT EXISTS user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(100),
    password VARCHAR(75),
    guardian VARCHAR(255),
    userPhone VARCHAR(16),
    guardianPhone VARCHAR(16),
    birthDate DATE,
    city VARCHAR(25),
    role VARCHAR(15) NOT NULL,
    lastEducation VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS subject (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    book VARCHAR(255),
    author VARCHAR(255),
    type VARCHAR(15),
    duration int
);