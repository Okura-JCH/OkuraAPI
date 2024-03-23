CREATE TABLE IF NOT EXISTS articles (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    image TEXT,
    category_id INT NOT NULL,
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    nickname TEXT,
    age INT,
    gender_id INT,
    image TEXT,
    in_kitakyushu BOOLEAN,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
);

CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS genders (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);


