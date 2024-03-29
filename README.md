# Simple RestAPI using GO Echo Framework

simple RestAPI with Go, Echo, Gorm, and MySQL

## Requirements

Simple RestAPI is currently extended with the following requirements.  
Instructions on how to use them in your application are linked below.

| Requirement | Version |
| ----------- | ------- |
| Go          | 1.21.5  |
| Mysql       | 8.0.35  |

## Installation

Make sure the requirements above are already installed on your system.  
Clone the project to your directory and install the dependencies.

```bash
$ git clone https://github.com/sun01822/Store-RestAPI
$ cd Store-RestAPI
$ go mod tidy
```

## Configuration
Copy the .env.example file and rename it to .env  
Change the config for your local server

```bash
DB_PORT=      3306
DB_USER=      root
DB_PASSWORD=
DB_NAME=      store
SERVER_PORT=  8080
```

## Running Server

```bash
$ go run main.go
```

## Endpoints

These are the endpoints we will use to create, read, update and delete the course data.

```bash
POST course → Add new course data
GET course → Retrieves all the course data
PUT course/{id} → Update the course data
DELETE course/{id} → Delete the course data
```
