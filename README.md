# go-gin-api

![Go](https://img.shields.io/badge/Go-1.21-blue) 
![Gin](https://img.shields.io/badge/Gin-1.9-lightgrey) 
![Gorm](https://img.shields.io/badge/Gorm-2.0-lightgrey) 
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-18-blue) 
![RabbitMQ](https://img.shields.io/badge/RabbitMQ-3.12-orange)
![Docker](https://img.shields.io/badge/Docker-24.0-blue)

**Student Management API using Go, Gin, GORM, PostgreSQL & RabbitMQ**

A RESTful API that supports full CRUD operations for students and uses event-driven messaging with RabbitMQ. Includes Docker Compose setup for easy local development.

---

## Tech Stack

- **Language:** Go  
- **Framework:** Gin  
- **ORM:** GORM  
- **Database:** PostgreSQL  
- **Messaging:** RabbitMQ  
- **Containerization:** Docker / Docker Compose  

---

## Features

- Full CRUD operations for students (`POST`, `GET`, `PUT`, `DELETE`)  
- GET endpoint to list all students  
- Messaging with RabbitMQ for asynchronous processing (e.g., student registration confirmation)  
- Structured project following Go best practices  
- Run locally via Docker Compose  

---

## Setup

### Prerequisites

- **Go 1.21+** – to run the API  
- **Docker 24+** – to containerize Postgres, RabbitMQ, and the API  
- **Docker Compose 2+** – to run all services with one command  

---

### Runing with Docker Compose

Run all services (PostgreSQL + RabbitMQ + API):

```bash
docker-compose up -d
```

This will start:

* PostgreSQL database

* RabbitMQ message broker

* Go Gin API

API will be available at: http://localhost:8080

---

### Endpoints

| Method | Endpoint               | Description                        |
|--------|-----------------------|-------------------------------------|
| GET    | /students             | Get all students                    |
| GET    | /students/:id         | Get a student by ID                 |
| GET    | /students/cpf/:cpf    | Search for a student by CPF         |
| POST   | /students             | Create a new student                |
| PATCH  | /students/:id         | Update an existing student by ID    |
| DELETE | /students/:id         | Delete a student by ID              | 

---

## Messaging with RabbitMQ

This project uses RabbitMQ for asynchronous messaging. 

- **Consumer:** `consumer.go` listens to the `students_queue` queue.
- It processes messages related to student data (for example, creating or updating students asynchronously).
- Ensure RabbitMQ is running (`docker-compose up`) before starting the consumer.
- Start the consumer:

```bash
go run consumer.go
```

---

### Project Structure
```text
go-gin-api/
├── main.go # Entry point of the application
├── consumer.go # RabbitMQ consumer
├── controllers/ # API endpoint handlers
├── routes/ # Route definitions
├── models/ # Database models (GORM)
├── database/ # PostgreSQL connection and setup
├── messaging/ # RabbitMQ connection and helpers
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

---

## License

MIT License

---

### Notes

This project is part of a [Dev Journey](https://github.com/FernandaIshida/dev-journey/blob/main/README.md): exploring backend development across multiple technologies:

- **Go, Gin, Gorm** – building REST APIs with event-driven architecture  
- **PostgreSQL** – relational database experience  
- **RabbitMQ** – asynchronous messaging  
- **Python & Django/DRF / FastAPI** – building APIs, microsservices, and automations  
- **Docker** – containerization and environment management  

It demonstrates full CRUD functionality, API design best practices, and event-driven development, while connecting this Go project to my wider learning path in backend engineering.

