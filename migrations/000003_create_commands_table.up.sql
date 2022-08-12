CREATE TABLE IF NOT EXISTS commands (
   id bigserial PRIMARY KEY,
   name text UNIQUE NOT NULL,
   text text NOT NULL,
   level integer NOT NULL
);

INSERT INTO commands (name,"text","level") VALUES
	('repeat','xset r rate 175 50',0),
	('eurkey','setxkbmap -layout eu',0),
	('clueless','ch02 ch21 ch31',0),
	('justinfan','64537',0),
	('streamlink','https://haste.zneix.eu/udajirixep put this in ~/.config/streamlink/config on Linux (or %appdata%\streamlink\streamlinkrc on Windows)',0),
	('gyazo','Gyazo is the worst screenshot uploader in human history. At best, it’s inconvenient, slow, and missing features: at worst, it’s a bandwidth-draining malware risk for everyone who views your images. There is absolutely no reason to use it unless you’re too lazy to spend 5 minutes installing another program.',250),
	('arch','Your friend isnt wrong. Being on the actual latest up to date software, having a single unified community repository for out of repo software (AUR) instead of a bunch of scattered broken PPAs for extra software, not having so many hard dependencies that removing GNOME removes basic system utilities, broader customization support and other things is indeed, pretty nice.',250),
	('arch2','One time I was ordering coffee and suddenly realised the barista didnt know I use Arch. Needless to say, I stopped mid-order to inform her that I do indeed use Arch. I must have spoken louder than I intended because the whole café instantly erupted into a prolonged applause. I walked outside with my head held high. I never did finish my order that day, but just knowing that everyone around me was aware that I use Arch was more energising than a simple cup of coffee could ever be.',250),
	('feelsdankman','⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⢀⣾⣿⣿⣿⣿⣷⣄⠹⣿⣿⣿⣿⣿⣿⣿⣿⣿ ⣿⣿⣿⣿⣿⣿⣿⣿⠟⢁⣴⣿⣿⣿⣿⣿⣿⣿⣿⣦⡈⢻⣿⣿⣿⣿⣿⣿⣿ ⣿⣿⣿⣿⣿⣿⡿⠁⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣄⠙⢿⣿⣿⣿⣿⣿ ⣿⣿⣿⣿⣿⡿⠃⣰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⡈⢻⣿⣿⣿⣿ ⣿⣿⣿⣿⡿⢁⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣄⠹⣿⣿⣿ ⣿⣿⣿⣿⠁⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠿⠿⠿⠿⠿⠆⠘⢿⣿ ⣿⣿⠟⠉⠄⠄⠄⠄⢤⣀⣦⣤⣤⣤⣤⣀⣀⡀⠄⠄⡀⠄⠄⠄⠄⠄⠄⠄⠙ ⣿⠃⠄⠄⠄⠄⠄⠄⠙⠿⣿⣿⠋⠩⠉⠉⢹⣿⣧⣤⣴⣶⣷⣿⠟⠛⠛⣿⣷ ⠇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠁⠒⠄⠄⠄⠄⠈⠉⠛⢻⣿⣿⢿⠁⠄⠄⠁⠘⢁ ⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣂⣀⣐⣂⣐⣒⣃⠠⠥⠤⠴⠶⠖⠦⠤⠖⢂⣽ ⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠛⠂⠐⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣠⣴⣶⣿⣿ ⠃⣠⣄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣠⣤⣄⠚⢿⣿⣿⣿⣿ ⣾⣿⣿⣿⣶⣦⣤⣤⣄⣀⣀⣀⣀⣀⣀⣠⣤⣤⣶⣿⣿⣿⣿⣷⡄⢻⣿⣿⣿ ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠈⣿⣿⣿',500),
	('dankhug','⣼⣿⣿⣿⣷⡄⠀⠀⠀⠀⠀⠀⠀⠀⢀⣤⣾⣿⣿⣿⣿⣿⣄⠀⠀⠀⠀⠀⠀ ⣿⣿⣿⣿⣿⣿⣆⠀⠀⠀⠀⠀⠀⢠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣧⠀⠀⠀⠀⠀ ⣿⣿⣿⣿⣿⣿⣿⣧⡀⠀⠀⠀⣠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⡀⠀⠀⠀ ⣿⣿⣿⣿⣿⣿⣿⣿⣷⡄⠀⢠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡄⠀⠀ ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣄⠺⠿⠿⠿⠿⠿⠿⠿⠿⠿⠿⠿⠿⠟⠛⠛⠀⠀ ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠆⠒⠀⠀⠶⢶⣶⣶⣭⣤⠹⠟⣛⢉⣉⣉⣀⣀ ⣿⣿⣾⣿⣶⣶⣶⣶⣶⣶⣿⣿⣶⠀⢬⣒⣂⡀⠀⠀⠀⠀⣈⣉⣉⣉⣉⡉⠅ ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡄⢭⣭⠭⠭⠉⠠⣷⣆⣂⣐⣐⣒⣒⡈ ⢿⣿⣿⣿⠋⢁⣄⡈⠉⠛⠛⠻⡿⠟⢠⡻⣿⣿⣛⣛⡋⠉⣀⠤⣚⠙⠛⠉⠁ ⠀⠙⠛⠛⠀⠘⠛⠛⠛⠛⠋⠀⠨⠀⠀⠀⠒⠒⠒⠒⠒⠒⠒⠊⡀⠀⠀⠀⠀ ⠀⠀⠀⠀⠀⠀⠀⠰⣾⣿⣿⣷⣦⠀⠀⠀⠀⠀⠀⠀⢀⣠⠴⢊⣭⣦⡀⠀⠀ ⠀⠀⠀⠀⠀⠀⠀⣠⣌⠻⣿⣿⣷⣞⣠⣖⣠⣶⣴⣶⣶⣶⣾⣿⣿⣿⣿⡀⠀ ⠀⢀⣀⣀⣠⣴⣾⣿⣿⣷⣌⠛⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⢻⣿⣷⡁ ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢻⣶⣬⣭⣉⣛⠛⠛⢛⣛⣉⣭⣴⣾⣿⣿⣿⡇',500),
	('shotgun','⡏⠛⢿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿ ⣧⣀⡀⠄⠹⠟⠻⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿ ⣿⣿⣧⠄⢈⡄⣄⠄⠙⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿ ⣿⣿⣿⠄⢸⣧⠘⢹⣦⣄⠈⠛⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿ ⣿⣿⣿⠄⢸⣿⡇⢸⣿⣿⣿⣶⣄⠉⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿ ⣿⣿⣿⠄⢸⣿⣷⠄⣿⣿⣿⣿⣿⣷⣦⡈⣙⠟⠉⠉⠙⠋⠉⠹⣿⣿ ⣿⣿⣿⠄⢸⣿⣿⡄⠸⣿⣿⣿⣿⣿⣿⡿⠃⠄⠄⣀⠄⢠⣀⠄⡨⣹ ⣿⣿⣿⠄⢸⣿⣿⣇⠄⠹⣿⣿⣿⣿⣿⠁⠄⠄⠄⠈⠄⠄⠄⠄⠠⣾ ⣿⣿⣿⠄⠈⣿⣿⣿⣆⠄⠈⠛⠿⣿⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⣿ ⣿⣿⣿⠄⠄⣿⣿⣿⣿⣦⣀⠄⠄⠈⠉⠄⠄⠄⠄⠄⠄⠄⠤⣶⣿⣿ ⣿⣿⣿⠄⠄⢻⣿⣿⣿⣿⣿⠷⠂⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠘⣿⣿ ⣿⣿⣿⣇⠄⠈⠻⣿⣿⠟⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣿ ⣿⣿⣿⣿⣦⠄⠄⠈⠋⠄⠄⣠⣄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣼⣿',500),
	('rope','⣿⣿⣿⡇⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿ ⣿⣿⣿⡇⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿ ⣿⣿⣿⡇⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿ ⣿⣿⣿⡇⠄⣿⣿⣿⡿⠟⠋⣉⣉⣉⡙⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿ ⣿⣿⣿⠃⠄⠹⠟⣡⣶⡿⢟⣛⣛⡻⢿⣦⣩⣤⣤⣤⣬⡉⢻⣿⣿⣿⣿⣿⣿⣿ ⣿⣿⣿⠄⢀⢤⣾⣿⣿⣿⣿⡿⠿⠿⠿⢮⡃⣛⣛⡻⠿⢿⠈⣿⣿⣿⣿⣿⣿⣿ ⣿⡟⢡⣴⣯⣿⣿⣿⣉⠤⣤⣭⣶⣶⣶⣮⣔⡈⠛⠛⠛⢓⠦⠈⢻⣿⣿⣿⣿⣿ ⠏⣠⣿⣿⣿⣿⣿⣿⣿⣯⡪⢛⠿⢿⣿⣿⣿⡿⣼⣿⣿⣿⣶⣮⣄⠙⣿⣿⣿⣿ ⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣾⡭⠴⣶⣶⣽⣽⣛⡿⠿⠿⠿⠿⠇⣿⣿⣿⣿ ⣿⣿⣿⣿⣿⣿⣿⠿⠿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣝⣛⢛⡛⢋⣥⣴⣿⣿⣿⣿⣿ ⣿⣿⣿⣿⣿⢿⠱⣿⣿⣛⠾⣭⣛⡿⢿⣿⣿⣿⣿⣿⣿⣿⡀⣿⣿⣿⣿⣿⣿⣿ ⠑⠽⡻⢿⣿⣮⣽⣷⣶⣯⣽⣳⠮⣽⣟⣲⠯⢭⣿⣛⣛⣿⡇⢸⣿⣿⣿⣿⣿⣿ ⠄⠄⠈⠑⠊⠉⠟⣻⠿⣿⣿⣿⣿⣷⣾⣭⣿⣛⠷⠶⠶⠂⣴⣿⣿⣿⣿⣿⣿⣿ ⠄⠄⠄⠄⠄⠄⠄⠄⠁⠙⠒⠙⠯⠍⠙⢉⣉⣡⣶⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿ ⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠙⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿',500),
	('porosad','⠀⠀⠀⠀⠀⠀⠀⠀⡀⠀⠀⢀⣠⣤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀ ⠀⠀⠀⠒⣛⠐⣠⣾⣿⣤⣾⣿⣿⣿⣿⣷⣠⣤⣤⢀⣀⣖⠒⠒⠀⠀⠀⠀⠀ ⠀⠀⠀⢀⣘⣘⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣯⣬⡕⠀⠀⠀⠀⣀⠠⠀⠀ ⠀⣠⣾⣿⣿⣿⣿⠿⠿⣿⣿⣿⠿⢿⣿⣿⣿⣿⣿⣿⣿⣷⣄⠀⠀⠀⠀⠀⠀ ⠐⢛⣿⡟⠋⢉⠉⠀⠔⣿⣿⣿⠆⠀⠩⡉⠙⢛⣿⠿⠿⠿⢛⣃⣀⡀⠀⠀⠀ ⠀⣾⣟⠁⢤⣀⣔⣤⣼⣿⣿⣿⣆⣀⠂⠴⣶⡁⠸⣿⣶⣾⣿⣿⣿⣿⠷⠂⠀ ⠀⣿⡿⡿⣧⣵⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣝⡛⣸⣿⣿⣿⣿⣿⣿⣷⡀⠀ ⠀⢿⢃⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⠸⣿⣿⣿⣿⣿⣿⣿⡇⠀ ⠀⠻⢸⣿⣿⣿⣿⣿⣿⠋⠻⣿⣿⡻⢿⣿⣿⣿⣿⡆⣿⣿⣿⠿⣿⠟⢿⡇⠀ ⠀⠠⣸⣿⣿⠟⣩⣈⣩⣴⣶⣌⣁⣄⡉⠻⣿⣿⣿⣧⢸⣿⣧⣬⣤⣤⣦⣤⠀ ⠀⠀⠸⢫⣶⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣄⢻⡇⣼⣿⣿⣿⣿⣿⣿⡟⠀ ⠀⠀⠀⠈⢹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⣱⣿⣿⣿⣿⣿⣿⡟⠀⠀ ⠀⠀⠀⠀⠀⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠁⠀⠀ ⠀⠀⠀⠀⠀⠈⠉⠉⠛⠛⠛⠛⠛⠛⠉⠉⠉⠛⠛⠉⠁⠀⠀⠈⠉⠁⠀⠀⠀ ',500),
	('toucan','░░░░░░░░▄▄▄▀▀▀▄▄███▄░░░░░░░░░░ ░░░░░▄▀▀░░░░░░░▐░▀██▌░░░░░░░░░ ░░░▄▀░░░░▄▄███░▌▀▀░▀█░░░░░░░░░ ░░▄█░░▄▀▀▒▒▒▒▒▄▐░░░░█▌░░░░░░░░ ░▐█▀▄▀▄▄▄▄▀▀▀▀▌░░░░░▐█▄░░░░░░░ ░▌▄▄▀▀░░░░░░░░▌░░░░▄███████▄░░ ░░░░░░░░░░░░░▐░░░░▐███████████ ░░░░░le░░░░░░░▐░░░░▐███████████ ░░░░toucan░░░░░░▀▄░░░▐██████████ ░░░░has arrived░░░░░░▀▄▄███████████',500);

