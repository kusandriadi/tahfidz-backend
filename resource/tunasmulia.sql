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

CREATE TABLE IF NOT EXISTS quranprogress (
    id INT AUTO_INCREMENT PRIMARY KEY,
    createdDate DATE,
    markForDelete BOOLEAN default false,
    surat VARCHAR(5),
    ayat VARCHAR(5),
    juz VARCHAR(10),
    method VARCHAR(25),
    userid int NOT NULL,
    FOREIGN KEY (userid) REFERENCES user(id)
);

CREATE TABLE IF NOT EXISTS subjectprogress (
    id INT AUTO_INCREMENT PRIMARY KEY,
    createdDate DATE,
    markForDelete BOOLEAN default false,
    userid int NOT NULL,
    subjectid int NOT NULL,
    FOREIGN KEY (userid) REFERENCES user(id),
    FOREIGN KEY (subjectid) REFERENCES subject(id)
);