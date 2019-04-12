CREATE TABLE users (
    id serial PRIMARY KEY,
    username text UNIQUE NOT NULL,
    email text UNIQUE NOT NULL,
    password text UNIQUE NOT NULL
);