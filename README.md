# Microblog Application
Microblog is a simplified microblogging platform similar to Twitter. It allows users to publish tweets, follow other users, and view timelines of the users they follow.

---

## Features

 - Publish tweets (max 280 characters).
 - Follow/unfollow other users.
 - View the timeline of tweets from users you follow.
 - High-performance backend powered by MongoDB.
 - Caching with Redis for fast timeline retrieval.
 - JWT-based authentication for secure API access.

---

## Tech Stack

- **Go (Golang)**: Backend development.
- **MongoDB**: NoSQL database for storing users and tweets.
- **Redis**: In-memory caching for fast responses.
- **Docker**: Containerization.

---

## Project Structure

```plaintext
microblog/
├── application/
│   ├── ports/           # Interfaces for use cases
│   └── usecases/        # Business logic
├── cmd/                 # Application entry point
├── domain/
│   ├── repository/      # Repository interfaces
│   └── models/          # Core domain models (User, Tweet)
├── infrastructure/
│   ├── adapters/
│   │   ├── cache/       # Redis implementation
│   │   └── persistence/ # MongoDB implementation
│   ├── database/        # MongoDB and Redis connection logic
│   ├── logger/          # Centralized logging
│   └── server/          # HTTP server and routes
├── interface/
│   ├── controllers/     # HTTP controller logic
│   └── dto/             # Data Transfer Objects
├── security/
│   └── auth/            # Authentication logic
├── services/            # Application services for each domain
├── tests/               # integration tests
```

## API Endpoints
### Public Endpoints
- GET /generate_token: Generate a JWT token.
### Private Endpoints
- POST	/api/user/follow
    - request
    ```
    {
        "user_id": "1",
        "target_id": "5"
    }
    ```
  
    - response
    ```
    {
        "message": "User followed successfully"
    }
    ```
  
- GET	/api/user/timeline
  - response
  ```
  {
    "tweets": [
        {
            "id": "20241203111150",
            "user_id": "2",
            "content": "This is a tweet 2",
            "created_at": "2024-12-03T14:11:50.619Z"
        },
        {
            "id": "20241203095046",
            "user_id": "2",
            "content": "This is a tweet 2",
            "created_at": "2024-12-03T12:50:46.317Z"
        },
        {
            "id": "20241203095043",
            "user_id": "2",
            "content": "This is a tweet",
            "created_at": "2024-12-03T12:50:43.577Z"
        },
        {
            "id": "20241203095041",
            "user_id": "5",
            "content": "This is a tweet",
            "created_at": "2024-12-03T12:50:41.645Z"
        },
        {
            "id": "20241203095037",
            "user_id": "5",
            "content": "This is a tweet 2",
            "created_at": "2024-12-03T12:50:37.887Z"
        }
    ]
  }
  ```
  
- POST	/api/tweet/
    - request
    ```
    {
        "user_id":"1",
        "content":"This is a tweet"
    }
    ```
  
    - response
    ```
    {
        "message": "Tweet published successfully"
    }
    ```

## Setup

### Prerequisites

Make sure you have the following installed:
 - Docker
 - Go (v1.20 or later)
 - MongoDB
 - Redis

### Running Locally

#### With Docker Compose

0.	Run tests:
1. Units
```bash
go test ./... -v
```
Integration
```bash
go test ./test/integration
```

1.	Clone the repository:
```bash
git clone https://github.com/nvalfre/microblog.git
cd microblog
```

2.	Build and run the services in the desired env (only dev atm):
```bash
APP_ENV=dev docker-compose up --build 
```

#### Without Docker

1.	Set up MongoDB and Redis locally.

2.	Clone the repository:
```bash
git clone https://github.com/nvalfre/microblog.git
cd microblog
```

3. Setup redis and mongo in localhost and default ports.

4. Build and run the services:
```bash
go run main.go
```

