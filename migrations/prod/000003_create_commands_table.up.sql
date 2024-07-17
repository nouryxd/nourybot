CREATE TABLE IF NOT EXISTS commands (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    channel text NOT NULL,
    text text NOT NULL,
    level integer NOT NULL,
    description text NOT NULL
);


INSERT INTO commands (name,"channel","text","level","description") VALUES
	 ('kek','nouryxd','lmao',0,'kek');
