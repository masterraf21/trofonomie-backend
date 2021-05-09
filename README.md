# Trofonomie backend
Backend for Trofonomie app

# Development Setup

## Tools

- Go version 1.3
- Go Modules for dependency management.
- MongoDB for Database

## Run Program

This program needs environment variables. For example there has been env file for dev named env-dev. Copy it to file name `.env` run with:

```shell
make dev
```

# Test Program
To ensure code quality, one must test the code by its units. One can use docker to initiate the test dependency such as MySQL. To do test, one also need env configuration. See example below that is suited for the docker-compose.test configuration. Name the below file to `env-test`

```env
MONGO_HOSTS=localhost:27117
MONGO_DATABASE=test

TIMEOUT_ON_SECONDS=120
OPERATION_ON_EACH_CONTEXT=500
```

Run the test using
```shell
make test
```

# Deployment
One can deploy this application using Docker. Make sure docker is installed on your system. There is already an example of env file for production called env-prod.
Create .env file from env-prod:
```shell
cp env-prod .env
```
Then deploy the database depency docker container using:

``` shell
./cmds/docker-up-prod.sh
```

Finally deploy the Go application using:
``` shell
./cmds/deploy.sh
```