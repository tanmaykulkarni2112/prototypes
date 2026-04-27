# DDoS Agent

A concurrent request simulator that demonstrates hitting a server endpoint multiple times using Go goroutines.

## Overview

This project consists of two main components:

### 1. Go Client (`main.go`)
A concurrent HTTP client that simulates high load by making multiple requests to a server endpoint using goroutines.

**How it works:**
- Spawns 50 worker goroutines
- Each worker makes 1000 HTTP requests to `http://localhost:3000/home`
- Uses `sync.WaitGroup` to synchronize goroutines and wait for all workers to complete
- Uses `atomic.AddInt64()` to safely track request count without race conditions
- Each client has a 5-second timeout per request
- No artificial delays between requests for maximum concurrency

**Key features:**
- Concurrent execution using goroutines (lightweight Go threads)
- Thread-safe request counting with atomic operations
- Proper cleanup and synchronization

### 2. Node.js Server (`nodeServer/`)

**Boilerplate setup:**
- Generated using `npx quick-express-cli`
- Express.js server running on port 3000

**Modified endpoints:**
- **`GET /home`** - Main endpoint being tested
  - Increments and logs a counter to track incoming requests
  - Logs the number of requests received to console
  - Returns a simple string response

- **`GET /count`** - Status endpoint
  - Returns the total number of hits to `/home` as JSON

- **`GET /api/home`** - Alternative endpoint
  - Returns JSON response

## How to run

### Start the Node.js server:
```bash
cd nodeServer
npm install
npm start
```

### Run the Go client (in a separate terminal):
```bash
go run main.go
```

## Project Goals

- Demonstrate concurrent HTTP requests using Go goroutines
- Measure server load handling
- Explore high-concurrency patterns in Go
- Test scalability of the Node.js endpoint
