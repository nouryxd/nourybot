CREATE TABLE IF NOT EXISTS timers (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    text text NOT NULL,
    channel text NOT NULL,
    repeat text NOT NULL
);

INSERT INTO timers (name,"text",channel,repeat) VALUES
	 ('nourybot-4m',' 4 minute timer xD','nourybot','4m'),
	 ('nourybot-10m','10 minute timer xd','nourybot','10m'),
	 ('nourybot-20m',' 20 minutes XD','nourybot','20m'),
	 ('uude-5m',' 5 minutes timer :)','uudelleenkytkeytynyt','5m'),
	 ('uude-10m',' 10 minutes timer :)','uudelleenkytkeytynyt','10m'),
	 ('xnoury-3m',' 3m timer','xnoury','3m'),
	 ('xnoury-15m',' 15 minutes timer :)','xnoury','15m');

