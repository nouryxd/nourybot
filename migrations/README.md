# Migrations

Tool: [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

```
$ sudo -u postgres psql
psql (14.3)
Type "help" for help.

postgres=# CREATE DATABASE nourybot;
CREATE DATABASE
postgres=# \c nourybot;
You are now connected to database "nourybot" as user "postgres".
nourybot=# CREATE ROLE username WITH LOGIN PASSWORD 'password';
CREATE ROLE
nourybot=# CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION
nourybot=# 
```

```
$ psql --host=localhost --dbname=nourybot --username=username
psql (14.3)
Type "help" for help.

nourybot=> 
```