# VantaGrid

VantaGrid is an advanced HTTP honeypot and threat monitoring framework built in Go. It is designed for security researchers, organizations, and penetration testers to capture and analyze unauthorized web traffic in real time. VantaGrid records request metadata, fingerprints clients, scores suspicious behavior, and provides a REST API for analysis.

---

## Overview

VantaGrid simulates web endpoints and monitors traffic patterns without exposing real services. Each request is captured with the following data:

- IP address of the client  
- HTTP method and requested path  
- User-Agent string  
- Unique fingerprint for client identification  
- Numerical threat score for suspicious patterns  
- Timestamp of the interaction  

This information allows researchers and security teams to detect attack attempts, repeated probing, or anomalous behavior.

---

## Architecture

```text
VantaGrid/
├── main.go           # Entry point
├── core/
│   ├── scoring.go    # Threat scoring engine
│   ├── fingerprint.go# Unique client fingerprint generation
├── handlers/
│   ├── http.go       # HTTP request handler
│   ├── api.go        # REST API for events
├── storage/
│   ├── memory.go     # Thread-safe in-memory storage
├── types/
│   ├── event.go      # Event data structure
````

Module Breakdown

core: Handles scoring logic and fingerprint generation for client requests.

handlers: Processes HTTP requests and exposes API endpoints for monitoring.

storage: Maintains thread-safe storage of all captured events.

types: Provides structured data models for events.

Features

Modular architecture for easy extension

Multi-threaded request handling

Client fingerprinting and threat scoring

REST API endpoint for integration with dashboards or analysis tools

Real-time monitoring of suspicious HTTP traffic

Extensible to add new endpoints or threat scoring rules

Execution Flow

Incoming HTTP Request
        │
http.HandleFunc("/") → HTTPHandler.Handle()
        │
Compute client fingerprint (SHA1)
        │
Compute threat score based on method, path, and User-Agent
        │
Store event in memory (thread-safe)
        │
Return static HTTP response

Installation and Execution

Install Go (version 1.20+)

Clone the repository:
````
git clone <repo-url>
cd VantaGrid
````
Run the application:
````
go run main.go
````

By default, VantaGrid listens on port 8080:
````
http://0.0.0.0:8080/
````
REST API endpoint for events:
````
http://0.0.0.0:8080/api/events
````
Example Event

{
  "ip": "192.168.1.105",
  "path": "/admin",
  "method": "GET",
  "user_agent": "Mozilla/5.0",
  "score": 55,
  "fingerprint": "a12b3c4d5e6f7g8h9i0j1k2l3m4n5o6p7q8r9s0",
  "time": "2026-04-03T13:20:00Z"
}

````

