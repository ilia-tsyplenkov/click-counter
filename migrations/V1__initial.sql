CREATE SCHEMA core;

CREATE TABLE core.banner_clicks (
    id integer NOT NULL,
    clicks integer NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX id_idx ON core.banner_clicks (id) WITH (deduplicate_items = off);
