CREATE TABLE IF NOT EXISTS timers (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    text text NOT NULL,
    channel text NOT NULL,
    repeat text NOT NULL
);

INSERT INTO timers (name,"text",channel,repeat) VALUES
	 ('nourylul-60m','timer every 60 minutes :)','nourylul','60m'),
	 ('nourybot-60m','timer every 60 minutes :)','nourybot','60m'),
	 ('nourybot-1h','timer every 1 hour :)','nourybot','1h'),
	 ('xnoury-60m','timer every 420 minutes :)','xnoury','420m'),
	 ('xnoury-1h','timer every 1 hour :)','xnoury','1h'),
	 ('xnoury-15m','180 minutes timer :)','xnoury','180m');

