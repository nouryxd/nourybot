CREATE TABLE IF NOT EXISTS timers (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    identifier text NOT NULL,
    text text NOT NULL,
    channel text NOT NULL,
    repeat text NOT NULL
);

INSERT INTO timers (name,identifier,"text",channel,repeat) VALUES
	 ('nourybot-5m','2ad01f96-05d3-444e-9dd6-524d397caa96','5m timer','nourybot','5m'),
	 ('nourybot-10m','2ad01f96-05d3-444e-9dd6-524d397caa96','10m timer','nourybot','10m');

