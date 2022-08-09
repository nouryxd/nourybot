CREATE TABLE IF NOT EXISTS users (
   id bigserial PRIMARY KEY,
   added_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
   login text UNIQUE NOT NULL,
   twitchid text NOT NULL,
   level integer NOT NULL
);