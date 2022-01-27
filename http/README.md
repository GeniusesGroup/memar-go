# HTTP

## Why - net/http disadvantages
- [ServeHTTP](https://github.com/golang/go/blob/master/src/net/http/server.go#L86) method belong to Handler interface depended to package itself by using request structure and force to remain in the net/http forever or refactor your services codes.
- Mix connection/stream data and request to one structure as [Request](https://github.com/golang/go/blob/master/src/net/http/request.go#L103)
- Use unnecessary pointer data like [URL](https://github.com/golang/go/blob/master/src/net/http/request.go#L124) in the Request structure.
- Unnecessary version [unmarshal](https://github.com/golang/go/blob/master/src/net/http/request.go#L132) in each request when [helper method](https://github.com/golang/go/blob/master/src/net/http/request.go#L399) exist, It isn't very performance penalty to just unmarshal and compare after in each request Because it uses very rare requirement.
- Declare some specific data in a request like [form data](https://github.com/golang/go/blob/master/src/net/http/request.go#L245) in this way why not have bodyAsJSON, ...
- When HTTP is a string-based protocol why use many integers like [response status code](https://github.com/golang/go/blob/master/src/net/http/server.go#L1134) to have many runtime base logic.
- net/http made [body](https://github.com/golang/go/blob/master/src/net/http/transfer.go#L809) to use by the concurrent situation, but it is impossible to use the body in two different goroutines.
- So many unnecessary memory allocations, like allocate independently allocate for each header key and value in both [Request](https://github.com/golang/go/blob/master/src/net/http/request.go#L1076) and [Response](https://github.com/golang/go/blob/master/src/net/http/response.go#L191)
- Like many other added libraries to go, errors handle and declare in very bad shape and location.
- ...

## Why - github.com/valyala/fasthttp disadvantages
- high allocs/op even more than net/http in just parse phase. We know this package highly uses the pool to fool GC but it can't fool itself because anyway huge copy(on parse and reset logic) need to fill those allocations. even so many unneeded copy occur e.g. [URI.SetPath()](https://github.com/valyala/fasthttp/blob/3ff6aaa5917f40eeb5cdcb4272c58210f161f0ea/uri.go#L177) or [all header values](https://github.com/valyala/fasthttp/blob/7eeb00e1ccc54b29a6a165c6a27d5dfa96b416ca/header.go#L339)
- Anything is byte slice, but RFC says HTTP is string base protocol. It is so easy to not copy buffers but change type from byte slice to string. We know it is needed unsafe package, but it worse it and easily can be contracted by package and consumers to know they must copy any data if they want to hold for a long time, otherwise, a memory leak occurs easily by any mistake logics.
- Implement many logics that don't belong to the HTTP layer:
    - has worker mechanism just to limit the number of concurrent and in [some way](https://github.com/valyala/fasthttp/blob/9f11af296864153ee45341d3f2fe0f5178fd6210/workerpool.go#L147) it will block itself. And is it a good layer to handle such low layer requirements?
    - How data came to package as application layer such as TCP, TLS, ...
    - Limiting such as IP limiting does not belong to the HTTP package.
- Like many other libraries in the Go ecosystem, errors handle and declare in very bad shape and location. Error isn't about log message for developers, It is a message to show to users, So at least it must include locale messages, not just English.
- ...

## Goals
- Respect [HTTP protocol](../protocol/http.go) and implement all requirements to be comply with those interfaces as http rules. So any other packages can easily be changeable with this package.

## Not Goals

## Protocols
HTTP/1 : https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol
HTTP/2 : https://httpwg.org/specs/rfc7540.html
HTTP/3 : https://quicwg.org/base-drafts/draft-ietf-quic-http.html

# Abbreviations
