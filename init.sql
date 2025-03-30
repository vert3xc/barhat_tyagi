CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    user_role TEXT NOT NULL DEFAULT 'User'
);

CREATE TABLE threads(
    id SERIAL PRIMARY KEY,
    thread_name TEXT NOT NULL
);

CREATE TABLE votings(
    id SERIAL PRIMARY KEY,
    thread_id INTEGER NOT NULL REFERENCES threads(id),
    title TEXT NOT NULL,
    descr TEXT
);

CREATE TABLE options(
    voting_id INTEGER REFERENCES votings(id),
    option_text TEXT NOT NULL,
    vote_count INTEGER DEFAULT 0,
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
    vote TEXT NOT NULL
);

INSERT INTO users (username, password_hash, user_role) VALUES ('admin', '713bfda78870bf9d1b261f565286f85e97ee614efe5f0faf7c34e7ca4f65baca', 'Admin')
