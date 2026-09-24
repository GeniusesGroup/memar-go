# Storage
Memar provide local storage engine in many abstraction as low level as `Page` to many indexing packages.

## Levels
Each storage type (e.g. file, ...) MUST provide 3 automatic storage layer as `Local`, `Cache`, `Distributed`

### Local

### Cache
StoragesCache is the interface that can implement by any Application to provide cache storage mechanism. Keep them on volatile memories and just save very common on non-volatile storages. All GUI & edge nodes of any software must use cache in many cases to improve performance.

Cache CAN occur in many layers such as GUI, Edge, Fog, Cloud, Source of truth, ...

### Distributed


## Embedded database
An [embedded database system](https://en.wikipedia.org/wiki/Embedded_database) is a database management system which is tightly integrated with an application software; it is embedded in the application. There is no need for administrator.

## Embedded Projects
Github related repository can found [here](https://github.com/topics/embedded-database)
- [Native DB](https://github.com/vincent-herlemont/native_db) -- Rust
- [sqlite](https://github.com/sqlite/sqlite) -- C
- [ChaiSQL](https://github.com/chaisql/chai) -- Go
- [Bow](https://github.com/zippoxer/bow) -- Go
- [CloverDB](https://github.com/ostafen/clover) -- Go
- [TinyDB](https://github.com/msiemens/tinydb) -- Python

### Standalone projects
- [Aerospike](https://aerospike.com/) -- C
