BINARY_NAME=Nourybot.out
BINARY_NAME_API=NourybotApi.out


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

xd:
	docker compose down
	docker compose build
	docker compose up

prod:
	cd cmd/nourybot && go build -o ${BINARY_NAME}  
	mv cmd/nourybot/${BINARY_NAME} ./bin/${BINARY_NAME}
	./bin/${BINARY_NAME} -env="prod"
