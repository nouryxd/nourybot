build:
	cd cmd/bot && go build -o Nourybot

run:
	cd cmd/bot && ./Nourybot

dev:
	cd cmd/bot && go build -o Nourybot && ./Nourybot -mode dev

prod:
	cd cmd/bot && go build -o Nourybot && ./Nourybot -mode production

xd:
	cd cmd/bot && go build -o Nourybot && ./Nourybot