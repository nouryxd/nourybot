CREATE TABLE IF NOT EXISTS timers (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    identifier text NOT NULL,
    text text NOT NULL,
    channel text NOT NULL,
    repeat text NOT NULL
);

INSERT INTO timers (name,identifier,"text",channel,repeat) VALUES
	 ('nouryxd-60m','678efbe2-fa2f-4849-8dbc-9ec32e6ffd3b','gopherDance','nouryxd','60m'),
	 ('nourybot-60m','2ad01f96-05d3-444e-9dd6-524d397caa96','gopherDance','nourybot','60m'),
	 ('xnoury-1h','2353fd22-fef9-4cbd-b01e-bc8804992f4c', 'AlienPls','xnoury','1h'),
	 ('uudelleenkytkeytynyt-1h','6e178e14-36c2-45e1-af59-b5dea4903fee','pajaDink','uudelleenkytkeytynyt','1h');

