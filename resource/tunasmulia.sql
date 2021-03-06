CREATE TABLE IF NOT EXISTS user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    createdDate DATE,
    markForDelete BOOLEAN default false,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(100),
    password VARCHAR(75),
    guardian VARCHAR(255),
    userPhone VARCHAR(16),
    guardianPhone VARCHAR(16),
    birthDate DATE,
    city VARCHAR(25),
    role VARCHAR(15) NOT NULL,
    lastEducation VARCHAR(255)
);

INSERT INTO user (id, createdDate, name, username, password, role)
VALUES (null, now(), 'admin', 'admin', '$2a$10$LWcLabDoYw9vt3NPnS2AD.liTuU5AGhh.Xu/qdRDf7lm1MJgHZRpa', 'ADMIN');

CREATE TABLE IF NOT EXISTS subject (
    id INT AUTO_INCREMENT PRIMARY KEY,
    createdDate DATE,
    markForDelete BOOLEAN default false,
    name VARCHAR(255) NOT NULL,
    book VARCHAR(255),
    author VARCHAR(255),
    type VARCHAR(15),
    duration int
);