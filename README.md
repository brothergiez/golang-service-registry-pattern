# **Service Registry**

The Service Registry is a **Golang** and **MongoDB** based application designed to manage services in a microservices architecture. It provides mechanisms to dynamically register services, retrieve a list of registered services, and deregister services.

## **Key Features**
- **Register Service**: Register new services into the registry.
- **List Services**: Retrieve a list of registered services.
- **Deregister Service**: Remove a specific service from the registry.

---

## **System Requirements**
- **Go**: Version 1.20 or newer.
- **MongoDB**: Version 4.x or newer.
- **Dependencies**:
  - `github.com/joho/godotenv`: For reading `.env` files.
  - `github.com/gorilla/mux`: For HTTP routing.
  - `go.mongodb.org/mongo-driver`: For interacting with MongoDB.


## **Project Structure**
```
service-registry/
├── cmd/
│   └── server/           # Application entry point
│       └── main.go
├── internal/
│   ├── config/           # Application configuration module
│   ├── database/         # MongoDB connection module 
│   ├── registry/         # Core Service Registry module
│   │   ├── handler.go    # HTTP Handlers
│   │   ├── model.go      # Data structures for services
│   │   ├── repository.go # Repository (database access)
│   │   ├── router.go     # API route definitions
│   │   └── service.go    # Business logic
│   └── util/             # Additional utilities
└── go.mod

```

## **Configuration**
Create a `.env` file in the root directory to store application configuration:

```env
MONGO_URI=mongodb://localhost:27017
SERVER_PORT=3000
DATABASE_NAME=service_registry
COLLECTION_NAME=services
```


## **Installation**
Follow these steps to set up and run the Service Registry project:

1. **Clone the Repository**
   Clone the project from GitHub:
   ```bash
   git clone https://github.com/username/service-registry.git
   cd service-registry
   ```


2. **Install Dependencies Install the required Go modules:**
   ```bash
   go mod tidy
   ```

2. **Run the Server**
   ```bash
   go run cmd/server/main.go
   ```

## **API Endpoints**
### **1. Register Service**

- **URL**: `/services`
- **Method**: `POST`
- **Request Body**
The body should be a JSON object with the following fields:
```json
{
  "name": "example-service",
  "address": "127.0.0.1",
  "port": 3000
}
```
- **Response (201 Created)**
```json
{
  "message": "Service registered successfully"
}
```

### **2. List Services**

- **URL**: `/services`
- **Method**: `GET`
- **Response (200 OK)**
```json
[
  {
    "id": "64a1e6d2f6c2b93e4d64f0a5",
    "name": "example-service",
    "address": "127.0.0.1",
    "port": 3000,
    "registered_at": "2024-12-12T10:00:00Z"
  }
]
```

### **2. Deregister Service**

- **URL**: `/services/{id}`
- **Method**: `DELETE`
- **Response (200 OK)**
```json
{
  "message": "Service deregistered successfully"
}

```

## **Installation**
### **1. Manual Testing: Use curl or Postman to test the endpoints. Example:**
```bash
curl -X GET http://localhost:3000/services
```

### **2. Automated Testing: Add unit tests for repository, service, and handler modules. Run all tests using:**
```bash
go test ./... -v
```

## **License**
This project is licensed under the MIT License.

