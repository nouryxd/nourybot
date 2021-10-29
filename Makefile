build:
	cd cmd && go build -o Nourybot

run:
	cd cmd && ./Nourybot

dev:
	cd cmd && go build -o Nourybot && ./Nourybot -mode dev

prod:
	cd cmd && go build -o Nourybot && ./Nourybot -mode production

