CREATE TABLE IF NOT EXISTS commands (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    channel text NOT NULL,
    text text NOT NULL,
    category text NOT NULL,
    level integer NOT NULL,
    help text NOT NULL
);

INSERT INTO commands (name,"channel","text","category","level","help") VALUES
	 ('repeat','nouryxd','xset r rate 175 75','default',0,'Command to set my keyboard repeat rate'),
	 ('xset','nouryxd','xset r rate 175 75','default',0,'Command to set my keyboard repeat rate'),
	 ('kek','nouryxd','lmao','default',0,'kek'),
	 ('lmao','nourybot','kek','default',0,'lmao'),
	 ('dockerclean','nouryxd','docker system prune -a --volumes','default',0,'clean docker');

