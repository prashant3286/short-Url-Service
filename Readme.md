# URL Shortener Service

This is a simple URL shortener service implemented in Go using MongoDB as the database. It generates unique and shortened aliases for long URLs and provides real-time redirection to the original links.

## Features

- Generates unique and short aliases for long URLs
- Redirects users to the original link when a short URL is accessed
- Links expire after a configurable default time span
- Highly available and designed for minimal latency

## Requirements

- Go programming language
- MongoDB database

## Installation

1. Make sure Go is installed on your system. If not, refer to the [official Go installation guide](https://golang.org/doc/install).

2. Install the necessary Go packages:


3. Install and set up MongoDB on your local machine or a remote server. Refer to the [official MongoDB documentation](https://docs.mongodb.com/manual/installation/) for instructions.

4. Clone this repository:


5. Update the MongoDB connection URL in the `repository/url_repository.go` file:

```go
options.Client().ApplyURI("mongodb://localhost:27017")

Replace "mongodb://localhost:27017" with the appropriate MongoDB connection URL.
```

6. Build and run the application:
    ```go run main.go```

## Running the service

1. Make sure to install docker on your machine
2. Run `docker-compose up` 


### The server will start on http://localhost:8080.


