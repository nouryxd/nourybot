BINARY_NAME=Nourybot.out
BINARY_NAME_API=NourybotApi.out

cup:
	sudo docker compose up

xd:
	cd cmd/bot && go build -o ${BINARY_NAME} && ./${BINARY_NAME} -env="dev"

xdprod:
	cd cmd/bot && go build -o ${BINARY_NAME} && ./${BINARY_NAME} -env="prod"

jq:
	cd cmd/bot && go build -o ${BINARY_NAME} && ./${BINARY_NAME} -env="dev" | jq

jqprod:
	cd cmd/bot && go build -o ${BINARY_NAME} && ./${BINARY_NAME} -env="prod" | jq

jqapi:
	go build -o ${BINARY_NAME_API} cmd/api && ./${BINARY_NAME} | jq

