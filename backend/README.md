# Team-9 – Gym Workout Tracker - Backend Application 🏋️


## 💪 Technologies Used

- **Go (Golang)** – Backend logic and HTTP API
- **Gin** – Web framework for Go
- **GORM** – ORM for database interactions
- **SQLite** – Lightweight embedded database
- **Docker** – Containerization for consistent environments
- **Docker Compose** – Orchestrates multi-container environments
- **Elasticsearch** – Log storage and search engine
- **Logstash** – Log pipeline to forward logs to Elasticsearch

---

## 📁 Project Structure

```
backend/
├── cmd/                  # Application entry point
│   └── main.go
├── internal/
│   ├── handlers/         # HTTP route handlers
│   ├── models/           # Domain models (WorkoutDay, Exercise)
│   ├── repositories/     # Database persistence logic
│   └── services/         # Business logic
├── tests/                # Integration and unit tests
├── utils/                # Helper functions
├── .env-example          # Sample environment variables
├── Dockerfile            # Docker image definition
├── docker-compose.yml    # Docker Compose services (API, Elasticsearch, etc.)
├── go.mod                # Go module definition
├── go.sum                # Module checksums
├── logstash.conf         # Logstash pipeline config
├── Makefile              # Dev commands (build, test, run)
└── README.md             # Project documentation
```

---

## 🚀 Getting Started

### 🔧 Prerequisites

- [Go](https://golang.org/dl/)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### 📦 Installation

1. **Clone the Repository**

```bash
git clone https://github.com/ProgramadoresSemPatria/Team-9.git
cd Team-9/backend
```

2. **Create Your Environment File**

```bash
cp .env-example .env
```

> Update the `.env` file with your local configuration.

3. **Run with Docker Compose**

```bash
docker-compose up --build
```

This will spin up:
- The API server
- Elasticsearch
- Logstash

The backend API will be available at:  
`http://localhost:5000`

Elasticsearch will be accessible at:  
`http://localhost:9200`

---

## 🔬 Running Tests

Run unit and integration tests:

```bash
make tests
```

---

## 🔍 Logging with ELK

- **Logstash** is configured via `logstash.conf` to parse and send logs to **Elasticsearch**
- Make sure all services (API, Elasticsearch, Logstash) are running for full log functionality

---

## 🧰 API Functionality (WIP)

- [x] Create Flows
- [x] Create a workout day
- [x] Add exercises to a workout day
- [x] Retrieve workout days and their exercises
- [X] Update or delete workout days, flows and exercises
- [X] Authentication



*Made with 💪 by Team-9*

