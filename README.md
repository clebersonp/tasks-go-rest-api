# tasks-go-rest-api
Simple rest api on `golang` to manager tasks in `postgresql` database.

## Requirements:
- [Go](https://go.dev/) 1.20 or newer
- Third party [chi package](https://pkg.go.dev/github.com/go-chi/chi) router for rest api
- Third party [viper package](https://pkg.go.dev/github.com/spf13/viper) for configuration
- [Postman](https://www.postman.com/) or `curl` to test the Rest API
- [Docker](https://www.docker.com/) to run [postgresql](https://hub.docker.com/_/postgres) container for persisting data

`NOTE`: Database name, table name, username, password, database host, database port and application port used in the following configuration need to match the [config.toml](config.toml) configuration file.

## Baby step setup:
`NOTE`: Run the following commands in terminal to run `postgresql` as a `docker container`, create database and table for the application:
- Build the `docker` image and run `postgresql`:
```docker
$ docker run --name api-tasks -p 5432:5432 -e POSTGRES_PASSWORD=secret -d postgres:15.3-alpine
```
- List all running docker processes:
```docker
$ docker container ps
```
- List all docker processes:
```docker
$ docker ps -a
```
- Logging into `psql` process as `interactive mode`:
```docker
$ docker exec -it api-tasks psql postgres -U postgres
```
- Create the user:
```postgresql
$ create user user_tasks;
```
- Give user password:
```postgresql
$ alter user user_tasks with encrypted password '5995';
```
- Create the database:
```postgresql
$ create database api_tasks;
```
- List all databases:
```postgresql
$ \l
```
- Connect to database:
```postgresql
$ \c api_tasks postgres
```
- Grant user to all on schema public:
```postgresql
$ GRANT ALL ON SCHEMA public TO user_tasks;
```
- Connect the user to the new database:
```postgresql
$ \c api_tasks user_tasks
```
- Create the `tasks` table:
```postgresql
$ CREATE TABLE tasks (id serial primary key, title varchar, description text, done bool default false);
```
- List all tables:
```postgresql
$ \d
```
- Exit the `psql` process:
```postgresql
$ \q
```
- Stopping `postgresql` docker container:
```docker
$ docker stop api-tasks
```
- Starting `postgresql` docker container:
```docker
$ docker start api-tasks
```
- Logging into `psql` process as `interactive mode` directly into the new database with new user:
```docker
$ docker exec -it api-tasks psql api_tasks -U user_tasks
```

## Starting:
- Int the root directory, run the following command to run the application:
```go
$ go run main.go
```

## Testing the Rest API with `Postman` or `curl`:
- `POST`: Create a new task:
```bash
$ curl --location --request POST 'localhost:8080/tasks' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Home work",
    "description": "I need to do all school tasks",
    "done": false
}'
```

- `GET`: List all tasks:
```bash
$ curl --location --request GET 'localhost:8080/tasks'
```

- `GET`: Retrieve a task by `id`:
```bash
$ curl --location --request GET 'localhost:8080/tasks/1'
```

- `PUT`: Update a task:
```bash
$ curl --location --request PUT 'localhost:8080/tasks/1' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Home work",
    "description": "I need to do all school tasks",
    "done": true
}'
```

- `DELETE`: Delete a task by `id`:
```bash
$ curl --location --request DELETE 'localhost:8080/tasks/1'
```