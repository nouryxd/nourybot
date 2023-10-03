CREATE TABLE IF NOT EXISTS timers (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    identifier text NOT NULL,
    text text NOT NULL,
    channel text NOT NULL,
    repeat text NOT NULL
);

INSERT INTO timers (name,identifier,"text",channel,repeat) VALUES
	 ('nourylul-60m','678efbe2-fa2f-4849-8dbc-9ec32e6ffd3b','timer every 60 minutes :)','nourylul','60m'),
	 ('nourylul-2m','63142f10-1672-4353-8b03-e72f5a4dd566','timer every 2 minutes :)','nourylul','2m'),
	 ('nourybot-60m','2ad01f96-05d3-444e-9dd6-524d397caa96','timer every 60 minutes :)','nourybot','60m'),
	 ('nourybot-1h','2353fd22-fef9-4cbd-b01e-bc8804992f4c', 'timer every 1 hour :)','nourybot','1h'),
	 ('xnoury-15m','6e178e14-36c2-45e1-af59-b5dea4903fee','180 minutes timer :)','xnoury','180m');

