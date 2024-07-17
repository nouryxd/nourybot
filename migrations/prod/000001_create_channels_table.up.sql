CREATE TABLE IF NOT EXISTS channels (
   id bigserial PRIMARY KEY,
   added_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
   login text UNIQUE NOT NULL,
   twitchid text NOT NULL
);

INSERT INTO channels (added_at,login,twitchid) VALUES
	 (NOW(),'nouryxd','31437432');

