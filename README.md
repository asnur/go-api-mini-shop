# API Mini Shop

## Description

This is a mini shop API, where you can create, update, delete and get products and orders. This API is built using the [GoFiber](https://gofiber.io/) framework and uses [PostgreSQL](https://www.postgresql.org/) as the database and [Redis](https://redis.io/) as the cache. This API also uses [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/) to run the project. This API is also equipped with [Postman](https://www.postman.com/) to make it easier for developers to see the documentation. This ERD diagram for this API is as follows:

<img src="mini_shop - ERD.png">

## Prerequisites

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Redis](https://redis.io/docs/install/install-redis/)
- [PostgreSQL](https://www.postgresql.org/download/)

## Installation

1. Clone the repository

```bash
git clone https://github.com/asnur/go-api-mini-shop.git
```

2. Create a `.env` file in the root directory of the project and copy the contents of the `.env.example` file into it.
3. Install the dependencies before running the project

```bash
go mod download
```

4. Run the following command to start the project without docker-compose

```bash
go run main.go server -i 0.0.0.0 -p 3000
```

5. Run the following command to start the project using docker-compose

```bash
docker-compose up . -d
```
