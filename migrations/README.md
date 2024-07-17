# Migrations

Dump postgresql database from docker container:
```sh
docker exec -t nourybot-postgres-1 pg_dumpall -c -U nourybot > dump_$(date +%Y-%m-%d_%H_%M_%S).sql
```

Restore postgres database to docker container:
```sh
cat dump_2024-07-17_22_23_32.sql | docker exec -i nourybot-postgres-1 psql -U nourybot
```
