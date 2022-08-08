CREATE TABLE IF NOT EXISTS channels (
   id bigserial PRIMARY KEY,
   joined_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
   login text NOT NULL,
   twitchid text NOT NULL,
   announce BOOLEAN NOT NULL
);