# rest-oracle-go-seed
A seed repository that contains a Go project that accepts input via a REST API and saves data to an Oracle database.

## Why Oracle?
The cost of an Oracle license is less than the cost of man hours that it would take to convert to another DBMS for most organizations. 

That being said, this seed moduralizes the data access module to make moving to Postgres or SQL server simple. 

## Why Go?
Golang uses signficantly less resources than Java, and is optimized to be a web development language.

The cloud savings could be significant.

## How to run

Build a docker image:

```bash
docker build -t go-seed .
```

Run the docker image:
```bash
# Modify DATA_SOURCE_NAME to contain your Oracle DSN
docker run --env DATA_SOURCE_NAME=user/pass@oracle_host:1521/orclcdb -t go-seed
```

Use CURL to hit an endpoint:
```bash
container_id=`docker ps | grep go-seed | cut -d " " -f 1`
ip=`docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $container_id`
curl -X GET $ip:8080/person
```

## Development and debugging using VSCode

Open this directory in VScode and install the following plugin: https://code.visualstudio.com/docs/languages/go

The developer machine must have Oracle Instant Client installed for this to work. Look at the Dockerfile for hints on what to install.

The directory `./vscode/` contains a launch file that runs this project in debug mode. This is useful for debugging and development. 

[launch.json](https://github.com/spardue/rest-oracle-go-seed/blob/main/.vscode/launch.json) must be modified to point at the correct Oracle DB instance.

To start debug mode, Press `F5` with main.go open.




