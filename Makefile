build:
	cd cmd/bot && go build -o Nourybot

run:
	cd cmd/bot && ./Nourybot

xd:
	cd cmd/bot && go build -o Nourybot && ./Nourybot

jq:
	cd cmd/bot && go build -o Nourybot && ./Nourybot | jq