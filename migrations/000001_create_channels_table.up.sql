CREATE TABLE IF NOT EXISTS channels (
   id bigserial PRIMARY KEY,
   added_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
   login text UNIQUE NOT NULL,
   twitchid text NOT NULL
);