CREATE TABLE IF NOT EXISTS commands (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    channel text NOT NULL,
    text text NOT NULL,
    level integer NOT NULL,
    description text NOT NULL
);

INSERT INTO commands (name,"channel","text","level","description") VALUES
	 ('repeat','nouryxd','xset r rate 175 75',0,'Command to set my keyboard repeat rate'),
	 ('xset','nouryxd','xset r rate 175 75',0,'Command to set my keyboard repeat rate'),
	 ('kek','nouryxd','lmao',0,'kek'),
	 ('lmao','nourybot','kek',0,'lmao'),
	 ('dockerclean','nouryxd','docker system prune -a --volumes',0,'clean docker');

