<div align="center">
  <h1>Graphql bookstore in go</h1>
  <img alt="Last commit" src="https://img.shields.io/github/last-commit/janapc/graphql-books-go"/>
  <img alt="Language top" src="https://img.shields.io/github/languages/top/janapc/graphql-books-go"/>
  <img alt="Repo size" src="https://img.shields.io/github/repo-size/janapc/graphql-books-go"/>

<a href="#project">Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#requirement">Requirement</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#run-project">Run Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#request-api">Request API</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#technologies">Technologies</a>

</div>

## Project

Api to manager bookstore using graphql with authentication.

## Requirement

To this project your need:

- golang v1.21 [Golang](https://go.dev/)
- docker [Docker](https://www.docker.com/)

Inside **.docker** create a first file mysql_password.txt and put your password mysql and create a second file mysql_root_password and put your password root mysql.This configuration wil be used by docker-compose.yaml.

In root folder create a file **.env** with:

```env
PORT_USER=4000
PORT_BOOKSTORE=8080
MYSQL_URL= # mysql url
JWT_SECRET= # secret jwt
```

## Run Project

Start Docker in your machine and run this commands in your terminal:

```sh
## up mysql
❯ docker compose up -d

## run this command to install dependencies:
❯ go mod tidy

## run this command to start user api(localhost:4000/user):
❯ go run user/cmd/server/server.go

## run this command to start bookstore api(localhost:8080/bookstore):
❯ go run bookstore/cmd/server/server.go

```

## Request API

Examples stay inside api folder

## Technologies

- golang
- mysql
- docker
- graphql

<div align="center">

Made by Janapc 🤘 [Get in touch!](https://www.linkedin.com/in/janaina-pedrina/)

</div>
