CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    login TEXT UNIQUE NOT NULL, 
    password_hash TEXT NOT NULL, 
    email TEXT UNIQUE NOT NULL, 
    name TEXT, 
    info TEXT
);

CREATE TABLE note (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES "user",
    name TEXT NOT NULL, 
    data TEXT NOT NULL, 
    public BOOLEAN,
    created TIMESTAMP NOT NULL,
    updated TIMESTAMP NOT NULL
);
