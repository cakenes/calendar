CREATE TABLE IF NOT EXISTS items (
    id serial,
    type text NOT NULL,
    text text NOT NULL,
    duration integer NOT NULL
);

ALTER TABLE ONLY items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);