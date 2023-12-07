CREATE TABLE IF NOT EXISTS commands (
    id bigserial PRIMARY KEY,
    name text UNIQUE NOT NULL,
    text text NOT NULL,
    category text NOT NULL,
    level integer NOT NULL,
    help text NOT NULL
);

INSERT INTO commands (name,"text","category","level","help") VALUES
	 ('repeat','xset r rate 175 75','default',0,'Command to set my keyboard repeat rate'),
	 ('xset','xset r rate 175 75','default',0,'Command to set my keyboard repeat rate'),
	 ('kek','lmao','default',0,'kek'),
	 ('lmao','kek','default',0,'lmao'),
	 ('dockerclean','docker system prune -a --volumes','default',0,'clean docker');

