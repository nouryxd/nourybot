CREATE TABLE IF NOT EXISTS timers (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    text text NOT NULL,
    channel text NOT NULL,
    repeat text NOT NULL
);
