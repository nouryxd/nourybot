build:
	cd cmd/bot && go build -o Nourybot

run:
	cd cmd/bot && ./Nourybot

xd:
	cd cmd/bot && go build -o Nourybot && ./Nourybot

jq:
	cd cmd/bot && go build -o Nourybot && ./Nourybot | jq

jqapi:
	cd cmd/api && go build -o Nourybot-Api && ./Nourybot-Api | jq
