# IPTV Browser

A lightweight IPTV aggregation service built in Go that collects and normalizes IPTV metadata such as countries, channels, streams, and logos into a unified data layer.

---

## Preview

![IPTV Browser Preview](static/Sample.png)

---

## Why This Project Exists

IPTV sources are often fragmented, inconsistent, and hard to consume directly.
This service acts as a unified aggregation layer that normalizes IPTV data into a structured format, making it easier to integrate into applications or downstream services.

---

## Features

* Aggregates IPTV metadata (countries, channels, streams, logos)
* Concurrent data fetching using Go routines
* Clean layered architecture (handler, service, repository)
* Lightweight and fast execution
* Easy to extend for caching, filtering, or additional sources

---

## Tech Stack

* Go (Golang)
* JSON processing
* REST architecture
* Goroutines & concurrency primitives

---

## Architecture

The project follows a simple layered architecture:

* **Handler Layer** → Handles HTTP requests and responses
* **Service Layer** → Contains business logic and data processing
* **Repository Layer** → Responsible for data fetching and external sources

This structure keeps the system modular, testable, and easy to scale.

---

## Project Structure

```text
iptv-browser/
├── internal/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   └── model/
├── static/
│   └── sample.png
├── cmd/
├── main.go
└── go.mod
```

---

## Concurrency Design

The system uses Go goroutines to fetch IPTV resources in parallel.
This reduces latency when aggregating data from multiple sources.
Synchronization is handled using wait groups to ensure safe completion of all concurrent tasks before returning results.

---

## Future Improvements

* Add caching layer (Redis)
* Add search and filtering capabilities
* Add pagination support
* Add OpenAPI/Swagger documentation
* Add Docker support
* Improve retry and error handling for external sources

---

## Disclaimer

This project only aggregates publicly available IPTV metadata and does not host or distribute any media content.

---

## License

MIT
