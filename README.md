<br />
<div align="center">
  <h3 align="center">Sabaody-KV</h3>
  <p align="center">
    A minimalist, high-performance in-memory key-value cache built from scratch in Go.
    <br />
    <a href="#architecture"><strong>Explore the docs »</strong></a>
  </p>
</div>

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#architecture--design">Architecture & Design</a></li>
    <li><a href="#benchmarks--load-testing">Benchmarks & Load Testing</a></li>
    <li><a href="#getting-started">Getting Started</a></li>
    <li><a href="#roadmap--future-features">Roadmap & Future Features</a></li>
    <li><a href="#current-problems">Current Problems</a></li>
  </ol>
</details>

## About The Project
> "What I cannot create, I do not understand." – Richard Feynman

**Sabaody-KV** is a minimalist in-memory key-value cache built from scratch in Go. This project is a deep dive into OS  internals, networking, and high-concurrency systems.

## Architecture & Design

<div align="center">
<img src="https://raw.githubusercontent.com/cocvu99/sabaody-kv/refs/heads/main/docs/image/%5B20260316%5D%20Sabaody-kv%20Architecture%20Diagram.jpg" alt="Architecture Diagram" width="800">

</div>

## Benchmarks & Load Testing
(TODO: Đưa số liệu TPS/RPS, P99 Latency và ảnh chụp màn hình Grafana).

## Getting Started

### Prerequisites
* Go 1.24+
* Linux/macOS (for advanced syscall features)

### Running the Server
```bash
# Clone the repository
git clone https://github.com/cocvu99/sabaody-kv

# Start the TCP server (Using Thread-pool method)
go run thread-pool/main.go
```

## Roadmap & Future Features
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

## Current Problems
(Vấn đề hiện tại: Bài toán xử lý TCP connection đứt gãy đột ngột, tối ưu hóa memory khi lượng key lớn).