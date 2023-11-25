CREATE TABLE users (
    id serial PRIMARY KEY,
    username varchar(128) NOT NULL,
    lastname varchar(128) NOT NULL,
    email varchar(128) NOT NULL UNIQUE,
    password varchar(256) NOT NULL,
    birthday TIMESTAMP NOT NULL
);

CREATE TABLE posts (
    id serial PRIMARY KEY,
    title varchar(128) NOT NULL,
    description text NOT NULL,
    user_id integer REFERENCES users(id) on delete cascade NOT NULL
);

CREATE TABLE comments (
    id serial PRIMARY KEY,
    text varchar(1024) NOT NULL,
    user_id integer REFERENCES users(id) on delete cascade NOT NULL,
    user_email varchar(128) REFERENCES users(email) on delete cascade NOT NULL,
    post_id integer REFERENCES posts(id) on delete cascade NOT NULL
);

CREATE TABLE likes (
    id serial PRIMARY KEY,
    user_id integer REFERENCES users(id) on delete cascade NOT NULL,
    comment_id integer REFERENCES comments(id) on delete cascade NOT NULL
);
