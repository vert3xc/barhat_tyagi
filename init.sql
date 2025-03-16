CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    user_role TEXT NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE votings(
    id SERIAL PRIMARY KEY,
    thread_id INTEGER NOT NULL REFERENCES threads(id),
    title TEXT NOT NULL,
    descr TEXT
);

CREATE TABLE options(
    id SERIAL PRIMARY KEY,
    voting_id INTEGER REFERENCES votings(id),
    option_text TEXT NOT NULL,
    vote_count INTEGER DEFAULT 0
    FOREIGN KEY(voting_id) REFERENCES votings(id) ON DELETE CASCADE
);

CREATE TABLE comments(
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    voting_id INTEGER REFERENCES votings(id),
    comment_text TEXT NOT NULL,
    FOREIGN KEY(voting_id) REFERENCES votings(id) ON DELETE CASCADE
);

CREATE TABLE votes(
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    voting_id INTEGER REFERENCES votings(id),
    vote INTEGER NOT NULL
);

CREATE TABLE threads(
    id SERIAL PRIMARY KEY,
    thread_name TEXT NOT NULL
);
