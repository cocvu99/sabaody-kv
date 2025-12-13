## sabaody-kv
Cache data is like bubbles - temporary, subject to bursting (deleting), but very important and effective. Written in golang


## Core Features ✨
- [x] **Implement a TCP Server that a client can send multiple requests to server**: Build a TCP server can handle requests from different clients 

- [ ] **Implement Thread Pool server**: Manage a pool of worker threads to reuse resources and control the number of parallel processing tasks.

- [ ] **Implement I/O multiplexing TCP Server**: Build a high-performance TCP server using I/O Multiplexing (non-blocking) to handle thousands of concurrent connections.

- [ ] **Implement RESP protocol**: Implement the standard Redis communication protocol to ensure compatibility with existing Redis Clients.

- [ ] **Implement PING, SET, GET, TTL command**: Fundamental commands to check connection, store, retrieve, and delete key-value data. Manage data lifecycle, allowing keys to be automatically deleted after a specific time duration.

- [ ] **Implement EXPIRE, DEL, EXISTS command**: Fundamental commands to check connection, store, retrieve, and delete key-value data. Manage data lifecycle, allowing keys to be automatically deleted after a specific time duration.

- [ ] **Implement Simple Set commands: SADD, SREM, SISMEMBER, SMEMBER**: Support collections of unique elements, useful for fast membership checks.

- [ ] **Implement Sorted Set commands: ZADD, ZSCORE, ZRANK**: Collections of unique elements ordered by score, commonly used for leaderboards.
    + [ ] *Use a 3rd party lib*
    + [ ] *Or implement skip list/b+tree yourself (optional)*

- [ ] **Implement Count-min sketch (CMS commands)**: Probabilistic data structures for existence checks or frequency counting with minimal memory footprint.
    + [ ] *CMS.INITBYDIM*
    + [ ] *CMS.INCRBY*
    + [ ] *CMS.QUERY*

- [ ] **Implement Bloom filter commands**: Probabilistic data structures for existence checks or frequency counting with minimal memory footprint.
    + [ ] *BF.RESERVE*
    + [ ] *BF.MADD*
    + [ ] *BF.EXISTS*

- [ ] **Implement INFO commands**: Provide real-time system statistics (memory usage, total keys, expiring keys...) essential for monitoring and troubleshooting.
    + [ ] *expires*
    + [ ] *avg_ttl*

- [ ] **Implement Cache eviction**: Strategy to automatically free up memory when full by removing least frequently or recently used keys.
    + [ ] *random*
    + [ ] *Approximated LRU*
    + [ ] *Approximated LFU (optional)*


- [ ] **Implement Graceful shutdown**: Ensure safe server termination: stop accepting new requests, finish ongoing tasks, and persist data before exiting.

- [ ] **Implement Shared-nothing architecture**: Architecture design that minimizes resource sharing between threads to reduce lock contention and boost performance.

- [ ] **Complete the Core Features**: 
    + [ ] *Measure benchmark*


## Advanced Features ✨

- [ ] **Advanced commands**: 
    + [ ] *Geospatial*: https://redis.io/docs/latest/develop/data-types/geospatial/
    + [ ] *List*
    + [ ] *Bitmap*
    + [ ] *Hyperloglog*
    + [ ] *Queue*
...
    
- [ ] **Cache pipeline**: Allow clients to send a batch of commands without waiting for individual responses, increasing network throughput.
    + Source: https://redis.io/docs/latest/develop/using-commands/pipelining/

- [ ] **Authentication**: Security mechanism requiring clients to authenticate via password before being authorized to execute any other commands.

- [ ] **Data persistence: RDB, AOF**: Mechanisms to persist data from RAM to disk for recovery after a server restart.

- [ ] **To be continued (TBC)**:



<!-- - [x] **Persistent, Immutable Storage**: Flushes full Memtables to immutable, sorted SSTable (.sst) files on disk
- [x] **Efficient Lookups**: SSTables are highly structured with block-based indexes, Bloom filters, Block cache, and Table cache
to minimize disk reads, especially for non-existent keys.
- [x] **Automatic Compaction**: process to merge SSTables, reclaim space, and optimize read performance.
- [x] **Single-Process Safety**: Implements an exclusive file lock to prevent concurrent access from multiple processes, ensuring data integrity.
### Future works
- [ ] **Leveled Compaction**: implement LevelDB's leveled compaction algorithm,
where files are organized into levels (L0, L1, L2...). This provides better scalability and more predictable performance by running smaller, more targeted compactions. -->


## Optional
- [ ] **Implement a thread-safe counter using Mutex**: 


# How to run code

TCP-server

```go build main.go```

Start Server:
# Terminal 1
```cd tcp-server```
```go run main.go```

# Terminal 2
Using telnet
```telnet localhost 3000```


using monitor script
```pip install psutil```