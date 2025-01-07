# Nexus

salimon nexus service is the gate that connects users to entities and manages profiles, authorization, credits and balancing

## Environment variables

a sample of all required environment variables are in `.env.sample` make sure you have them set as envrionment variable while service is running.
if you provide a `.env` same as `.env.example`, while using docker compose to bootstrap the service, it will take values from there.

| name           | type   | description                               |
| -------------- | ------ | ----------------------------------------- |
| ENV            | string | envrionment of service running            |
| PGSQL_HOST     | string | pgsql database host ip/domain             |
| PGSQL_PORT     | string | pgsql database port (no default provided) |
| PGSQL_USERNAME | string | pgsql username                            |
| PGSQL_PASSWORD | string | pgsql password                            |
| PGSQL_DBNAME   | string | pgsql database name for service           |

## Building project locally

to build and run the service locally you need latest golang installed on your local machine. then run:

```shell
go build -o bootstrap .
```

and then:

```shell
./bootstrap
```

## Running with docker

make sure you have docker installed and running on host machine and you have all required envrionment variables `.env` then run:

```shell
docker compose up
```
