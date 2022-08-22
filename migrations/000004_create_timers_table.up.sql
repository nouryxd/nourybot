CREATE TABLE IF NOT EXISTS timers (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    text text NOT NULL,
    channel text NOT NULL,
    repeat text NOT NULL
);

INSERT INTO timers (name,"text",channel,repeat) VALUES
	 ('nourylul-2m','2m timer','nourylul','2m'),
	 ('nourylul-3m','3 minute timer xD','nourylul','3m'),
	 ('nourylul-5m','5minute timer lulw','nourylul','5m'),
	 ('nourylul-10m','10 minute timer xD','nourylul','10m'),
	 ('nourylul-15m',' every 15 minutes :)','nourylul','15m'),
	 ('nourybot-4m',' 4 minute timer xD','nourybot','4m'),
	 ('nourybot-10m','10 minute timer xd','nourybot','10m'),
	 ('nourybot-20m',' 20 minutes XD','nourybot','20m'),
	 ('uude-5m',' 5 minutes timer :)','uudelleenkytkeytynyt','5m'),
	 ('uude-10m',' 10 minutes timer :)','uudelleenkytkeytynyt','10m'),
	 ('xnoury-3m',' 3m timer','xnoury','3m'),
	 ('xnoury-15m',' 15 minutes timer :)','xnoury','15m');

