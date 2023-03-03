BINARY_NAME=Nourybot.out
BINARY_NAME_API=NourybotApi.out

jqapi:
	cd cmd/api && go build -o Nourybot-Api && ./Nourybot-Api | jq
  go build -o ${BINARY_NAME_API} cmd/api
	./${BINARY_NAME} | jq

xd:
  go build -o ${BINARY_NAME} cmd/bot/main.go
	./${BINARY_NAME} -env="dev"

xdprod:
  go build -o ${BINARY_NAME} cmd/bot/main.go
	./${BINARY_NAME} -env="prod"

jq:
  go build -o ${BINARY_NAME} cmd/bot/main.go
  ./${BINARY_NAME} -env="dev" | jq

jqprod:
  go build -o ${BINARY_NAME} cmd/bot/main.go
  ./${BINARY_NAME} -env="prod" | jq

