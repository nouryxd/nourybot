build:
	cd cmd/bot && go build -o Nourybot

run:
	cd cmd/bot && ./Nourybot

dev:
	cd cmd/bot && go build -o Nourybot && ./Nourybot -env development

prod:
	cd cmd/bot && go build -o Nourybot && ./Nourybot -env production

xd:
	cd cmd/bot && go build -o Nourybot && ./Nourybot
