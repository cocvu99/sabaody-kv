# Sabaody-KV 🫧

> "What I cannot create, I do not understand." – Richard Feynman

**Sabaody-KV** is a minimalist in-memory key-value cache built from scratch in Go. Much like the bubbles of Sabaody, the data is temporary and fleeting, but the underlying system is fundamental and robust. 

This project is a deep dive into database internals, networking, and high-concurrency systems.

---

## Getting Started

### Prerequisites
* Go 1.24+
* Linux/macOS (for advanced syscall features)

### Running the Server
```bash
# Clone the repository
git clone https://github.com/cocvu99/sabaody-kv

# Start the TCP server
go run thread-pool/main.go
```

## Roadmap & Core Features

### 1. Networking & Concurrency
- [x] **Persistent TCP Server**: Handles multiple requests over a single connection to minimize handshake overhead.
- [x] **Worker Pool**: Efficiently manages Goroutines to reuse system resources and limit parallel tasks.
- [ ] **I/O Multiplexing**: High-performance, non-blocking TCP server capable of handling thousands of concurrent connections (Epoll/Kqueue).
- [ ] **Shared-nothing Architecture**: Designed to minimize lock contention and boost horizontal scalability.

### 2. Protocol & Commands
- [ ] **RESP Implementation**: Full support for the Redis Serialization Protocol for client compatibility.
- [ ] **Key-Value Operations**: Standard commands: `SET`, `GET`, `DEL`, `EXISTS`.
- [ ] **Time-to-Live (TTL)**: Automatic data expiration using `EXPIRE` and `TTL` commands.
- [ ] **Atomic Operations**: Ensuring data integrity during concurrent access.

### 3. Advanced Data Structures
- [ ] **Sets & Sorted Sets**: Support for `SADD`, `SMEMBERS`, `ZADD`, and `ZRANK`.
- [ ] **Probabilistic Structures**: Efficient existence checks using Bloom Filters and frequency counting with Count-Min Sketch.
- [ ] **Geospatial & Bitmaps**: (Planned) Advanced data types for complex indexing.

### 4. System & Reliability
- [ ] **Cache Eviction**: Automated memory management via Approximated LRU/LFU policies.
- [ ] **Persistence**: Data durability through RDB snapshots and AOF (Append Only File).
- [ ] **Graceful Shutdown**: Ensures zero data loss and safe connection termination during exit.
- [ ] **Monitoring**: Real-time statistics via the `INFO` command.

---
### Todo 2026-Mar-10: 
- [ ] Re-use and maintain the connection between client and server
- [ ] Pass data from buffer (read) client and send to the server
- [ ] Handle all exception error cases: client closing connection, ... TODO
- [ ] TODO: Check all exception error case and handle them
