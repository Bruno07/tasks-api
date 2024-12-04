# Tasks Manager API
The Task Management API is an application developed in Go that offers features to facilitate task monitoring and management. It has two access levels: Technician and Manager, each with specific specifications.

Technicians create, update and view only their own tasks, allowing them to personally organize and monitor their activities. Managers, on the other hand, have global access to view all tasks for any user and delete them if necessary, ensuring complete oversight and control.

To improve communication, the API uses an asynchronous messaging system with RabbitMQ, automatically notifying managers whenever a new task is created. This ensures that those responsible are informed in real time about new demands.

The API follows RESTful principles, with JWT-based authentication for access control, and persists your data with the MYSQL database.


# Table of content
* [Features](#features)
* [Getting started](#getting-started)
    * [Prerequisites](#prerequisites)
    * [Installation](#installation)
* [Usage](#usage)
    * [Docker](#docker)
* [API Endpoints](#api-endpoints)

# Features
* Authentication with JWT
* Create tasks
* Update tasks
* Delete tasks
* Retrieve tasks by ID
* List tasks

# Getting started

## Prerequisites
* Go (1.23 or later)
* MYSQL (5.7 or later)
* RabbitMQ (latest For messaging)

## Installation
To install the Tasks API, follow these steps:

1. Clone this repository to your local machine:
```bash
git clone git@github.com:Bruno07/tasks-api.git
```
2. Install dependencies:
```bash
go mod tidy
```
# Usage

## Docker compose
To execute the Tasks API using docker-compose, follow these steps:

1. Configure environment variables:

```bash
cp .env.axample .env
```

For testing purposes, some variables are already filled in, feel free to change them.

2. Start container 
```bash
docker compose up -d --build
```
The API can be accessed at http://localhost:3000. You can use tools like Postman or cURL to interact with the API.

This will start the following containers:
* RabbitMQ
* MYSQL

3. Access container
```bash
docker compose exec tasks-app bash
```

4. Run seeder:
```bash
go run cmd/seeder/main.go
```

Users created after seeder execution:

| Email                 | Password | Profile
|-----------------------|----------|--------------
| master@email.com      | 12345678 | 1
| tec1@email.com        | 12345678 | 2
| tec2@email.com        | 12345678 | 2


# API Endpoints
* POST /login: Auth login
* POST /api/users: Create a new user
* GET /api/tasks: Get all tasks
* POST /api/tasks: Create a new task
* GET /api/tasks/:id: Get a task by ID
* PUT /api/tasks/:id: Update a task
* DELETE /api/tasks/:id: Delete a task