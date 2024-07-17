CREATE TABLE IF NOT EXISTS timers (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    identifier text NOT NULL,
    text text NOT NULL,
    channel text NOT NULL,
    repeat text NOT NULL
);

INSERT INTO timers (name,identifier,"text",channel,repeat) VALUES
	 ('nouryxd-60m','678efbe2-fa2f-4849-8dbc-9ec32e6ffd3b','gopherDance','nouryxd','60m');

