# yalo-api

Codechallenge for Yalo

The project is a REST API that allows to create and process user interactions.

## Installation
In order to install the project I added a make file to make it easier.

```bash
make install-all
```

This command will install all the dependencies and run the tests.

### Dockerfile
The project has a Dockerfile that allows to run the project in a container.

```bash
make docker-build
make docker-run
```

This will build the image and run the container.

### Dockercompose

The project has a docker-compose file that allows to run the project with a database.

```bash
make docker-compose-up
```

This will run the project.

You can also turn it off with:

```bash
make docker-compose-down
```

## Usage

The project has a few endpoints that allow to create and process user interactions.

### Create Interaction

```bash
curl --location 'http://localhost:8080/user_interaction' \
--header 'Content-Type: application/json' \
--data '{"user_interaction": [{
    "user_id": "12345",
    "product_sku": "x30",
    "action": "add_to_cart",
    "interaction_timestamp": "2024-06-17T12:31:56Z",
    "interaction_duration": 300
    }]
}'
```

### Get Interactions

```bash
curl --location 'http://localhost:8080/user_interaction/12345'
```

### Collection Interactions

You get the collection with the request on the file `YALO.postman_collection.json`

## Populate the database

The project has a command that allows to populate the database with some interactions, for this is required:
*  to have python 3 installed.
* project already running

```bash
make populate-db
```