BINARY_NAME=Nourybot.out
BINARY_NAME_API=NourybotApi.out

xd:
	cd cmd/nourybot && go build -o ${BINARY_NAME}  
	mv cmd/nourybot/${BINARY_NAME} ./bin/${BINARY_NAME}
	./bin/${BINARY_NAME} -env="dev"

xdprod:
	cd cmd/nourybot && go build -o ${BINARY_NAME}  
	mv cmd/nourybot/${BINARY_NAME} ./bin/${BINARY_NAME}
	./bin/${BINARY_NAME} -env="prod"

jq:
	cd cmd/nourybot && go build -o ${BINARY_NAME}  
	mv cmd/nourybot/${BINARY_NAME} ./bin/${BINARY_NAME}
	./bin/${BINARY_NAME} -env="dev" | jq

jqprod:
	cd cmd/nourybot && go build -o ${BINARY_NAME}  
	mv cmd/nourybot/${BINARY_NAME} ./bin/${BINARY_NAME}
	./bin/${BINARY_NAME} -env="prod" | jq

build:
	cd cmd/nourybot && go build -o ${BINARY_NAME}  
	mv cmd/nourybot/${BINARY_NAME} ./bin/${BINARY_NAME}

run:
	./bin/${BINARY_NAME} -env="prod"

up:
	docker compose up

down:
	docker compose down

prod:
	cd cmd/nourybot && go build -o ${BINARY_NAME}  
	mv cmd/nourybot/${BINARY_NAME} ./bin/${BINARY_NAME}
	./bin/${BINARY_NAME} -env="prod"
