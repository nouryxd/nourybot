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
	 ('xnoury-60m','timer every 60 minutes :)','xnoury','60m'),
	 ('xnoury-1h','timer every 1 hour :)','xnoury','1h'),
	 ('nourybot-2m',' 2 minute timer xD','nourybot','2m'),
	 ('nourybot-4m',' 4 minute timer xD','nourybot','4m'),
	 ('nourybot-7m',' 7 minute timer xD','nourybot','7m'),
	 ('nourybot-10m','10 minute timer xd','nourybot','10m'),
	 ('nourybot-20m','20 minutes XD','nourybot','20m'),
	 ('uude-5m','5 minutes timer :)','uudelleenkytkeytynyt','5m'),
	 ('uude-10m','10 minutes timer :)','uudelleenkytkeytynyt','10m'),
	 ('uude-15m','10 minutes timer :)','uudelleenkytkeytynyt','15m'),
	 ('uude-30m','10 minutes timer :)','uudelleenkytkeytynyt','30m'),
	 ('xnoury-3m','3m timer','xnoury','3m'),
	 ('xnoury-15m','15 minutes timer :)','xnoury','15m');

