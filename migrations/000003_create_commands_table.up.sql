CREATE TABLE IF NOT EXISTS commands (
   id bigserial PRIMARY KEY,
   name text UNIQUE NOT NULL,
   text text NOT NULL,
   permission integer NOT NULL
);
