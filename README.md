# ☕ Online Coffee RESTful API (Golang Microservices)

A lightweight **RESTful API for managing coffee products**, built using **pure Go (`net/http`)** with **microservice-oriented design principles**.  
This project intentionally avoids frameworks to gain a deep understanding of Go’s HTTP internals, routing, and request handling.

---

## 🚀 Features

- RESTful CRUD operations for coffee products
- Built using **standard Go libraries only**
- Custom HTTP routing using `http.ServeMux`
- Regex-based URL parameter extraction
- Graceful server shutdown using OS signals
- Clean separation of handlers, data, and server logic
- JSON request and response handling

---

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **HTTP Server:** `net/http`
- **Routing:** Custom routing + Regex
- **Logging:** `log`
- **Graceful Shutdown:** `context`, `os/signal`, `syscall`
- **Data Store:** In-memory (learning purpose)

---

## 📁 Project Structure

Microservices with go/
│
├── main.go
├── handlers/
│ └── products.go
├── data/
│ └── products.go
├── go.mod
└── README.md

---

## ▶️ Running the Application

```bash:-
go run main.go

Server starts on 
http://localhost:9090


## 📡 API Endpoints

| Method | Endpoint       | Description             |
| -----: | -------------- | ----------------------- |
|    GET | /products      | Get all coffee products |
|   POST | /products      | Add a new product       |
|    PUT | /products/{id} | Update product by ID    |
| DELETE | /products/{id} | Delete product by ID    |

## Postman Requests

### 1. Get all products
- Method: GET
- URL: http://localhost:9090/products
- Headers: Not required
- Body: None

---

### 2. Add a product
- Method: POST
- URL: http://localhost:9090/products
- Headers:
  - Content-Type: application/json
- Body (raw → JSON):
{
  "name": "Espresso",
  "price": 120

}


---

### 3. Update a product
- Method: PUT
- URL: http://localhost:9090/products/1
- Headers:
  - Content-Type: application/json
- Body (raw → JSON):
{
  "name": "Latte",
  "price": 150

}

---

### 4. Delete a product
- Method: DELETE
- URL: http://localhost:9090/products/1
- Headers: Not required
- Body: None

---

## 🧪 Sample cURL Requests

Get all products:-
curl http://localhost:9090/products

Add a product:-
curl -X POST http://localhost:9090/products \
-H "Content-Type: application/json" \
-d '{"name":"Espresso","price":120}'

Update a product:- 
curl -X PUT http://localhost:9090/products/1 \
-H "Content-Type: application/json" \
-d '{"name":"Latte","price":150}'

---

## Core Concepts Demonstrated

- Building REST APIs using only net/http
- Implementing custom ServeHTTP handlers
- Explicit HTTP method handling
- Regex-based URI validation and ID extraction
- Proper use of http.ResponseWriter and HTTP status codes
- Graceful shutdown using context with timeouts

---

## Graceful Shutdown

The server listens for OS interrupt signals and shuts down gracefully, allowing active requests to complete before exiting.

---

## Learning Objective

This project was built to:

- Strengthen Golang backend fundamentals
- Understand HTTP request and response lifecycle
- Learn microservice-ready project structuring
- Avoid framework abstraction at early stages

---

## Future Enhancements

- Persistent database integration
- Authentication and authorization
- API versioning
- Docker and containerization
- Unit and integration testing
- Swagger and OpenAPI documentation

---

## Author

Hrishab Pachange  
Computer Science Student and Golang Developer  
AI/ML Engineer Intern  
REST APIs and Microservices Enthusiast


