CREATE TABLE IF NOT EXISTS users (
    id serial,
    username text NOT NULL,
    password text NOT NULL
);

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

ALTER TABLE ONLY users
    ADD CONSTRAINT users_username_key UNIQUE (username);