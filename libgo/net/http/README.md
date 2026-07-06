# HTTP
Implement the [HTTP working group of IETF](httpwg.org) protocol.

[See why we develope this package if you prefer image vs code||text](https://viewer.diagrams.net/?tags=%7B%7D&highlight=0000ff&edit=_blank&layers=1&nav=1#R7Vtbc5s4FP41fmwGCTD2Y%2BKk2Yfdmcxmpt3mpaOADGpl5JXl2u6vXwHiJmGHeG2ILy8xOkgCnfN9OheRgT2ZrR85mkd%2FsQDTAbSC9cC%2BH0AILAjkTyLZZJIhgJkg5CRQnUrBM%2FmN85FKuiQBXtQ6CsaoIPO60GdxjH1RkyHO2arebcpo%2FalzFGJD8Owjakq%2FkkBEmXQEvVL%2BByZhlD8ZDMfZnRnKO6uVLCIUsFVFZD8M7AlnTGRXs%2FUE00R5uV6ycZ%2B33C1ejONYtBnwyCdf4ruv0feX2%2Fj59uXHC%2F%2Fy%2B5ObzfIL0aVasM%2FmG%2FXCYpNrAQdSKarJuIhYyGJEH0rpHWfLOMDJoyzZKvv8ydhcCoEU%2FsBCbJSF0VIwKYrEjKq7chV8808y%2FsbNm9%2Bq9%2B7XavKstclbayKyYbYDVPtb2rbAWLXLoUmjOvIJczLDAnMly1aeLHerlpVowZbcxztUm6MV8RCLHf3sAguSRJjJt%2BHSABbHFAnyq%2F4eSKE5LPqVBpcXyubvsD8w7B9hFEh1FIwiLE54Q9bSuNB6XU6nyV0NH3XrryIi8PMcpepZyX2hbmn1SMwFXu%2FWsqkVNWCkOKU2FSfn2KqkqG0pWVShp20dSY9eL5RR0K%2FD3t0N%2BoJmVZJVOLeFZgckBjwFYkCDGGIzx1LyN%2F53iRciNTZf%2BsnFwJPqGVK5lrtXyY1hmFxJfC8y7sh%2BJA7NHgPvfgeRQB9EshuIBGADkYbHIpJt6N1UpIUoZX7vm1ChhFx5TkvlwWMpb2Qoz9yW4uA2CYtkK2YxriulrsHSs0K3ssOAlvuLGlVsMd1tMG097xbzVsznNlgvl7Xeh9QTnhiJRYkee1xHT0HFfIpsmWpUNa7TJnIcbSLdyWV6MCZKEVYse3%2FQjY8Eur0wd4XcJUAOnGjY6rlvx63Q6zJuBQ1k%2FeiB6409GlVp%2Fsm6eTNJTFrHzPXstvzfApD%2FyeyxrYUibseENMO2swyXdQL3Hy8Ds4JzGPd7ajFfaw%2FcqwsG1vBAPhgA0LMTHh4JeacV%2BF1x1zXuPAN3sdQ8%2FBwJMTcQKPd6UYed9C%2FsJ54wyngJyymhVBMhSsJYNn2JiSRmSFyUID6it%2BrGjARBGg81%2BZ06uA%2FgeqBVVztwTdcz7DR0NKsNlLyG7HxNMHLdD2YCM%2Fe%2BHt8caGPPo7g3N3bYb53aMiBwEqmwY72dCnd7hFMg44RS4f7OcOBpsAMa7Ng7Lx24dwP3%2FnUj8MfMTXVG9Z%2Bbwl2HOZkirQUlfvL7MY50XP08rOlIx27Q4dGOdKBj6NDcpi4gv2%2Fvjq1mC3eTZ3l6crRvmuXp%2BVrHWRY8Vl3ppLL7K%2Bo6Rp1ZU5qihTjvzN6FdaX3nlZCs8Kyi%2Fs%2BRYsF8VMDIC5M8dZdoS1Te6uaaeecrrUnr%2FSJnI7PZ6BZq3lMa2ZPS2FGsySJz%2BaM0fPlnF7ELCzbG%2Bfe9xnF%2BXLOcZzDcE6fqGvO5VC5ci43CNhi2b44Z5ufkVwm57zRgfycPlHnnGuo96R6FCwpOeBM6ckLs2kijXD6l2MUpGZH%2Fs%2Fz5Z%2F%2B8YILTP6NOuWfWR%2B6TP457vimfrRk6wWl1l7PmKr4oKwrDpoVqwv3e9rHWbChqNit33tfbed8eTfSsoC9WadP1DnnzLrJ1e9tC0oKYlVAdqCPbmWz%2FIfOzLjlv8XaD%2F8B)

## Why - net/http disadvantages
- [ServeHTTP](https://github.com/golang/go/blob/master/src/net/http/server.go#L86) method belong to Handler interface depended to package itself by using request structure and force to remain in the net/http forever or refactor your services codes.
- Mix connection/stream data and request to one structure as [Request](https://github.com/golang/go/blob/master/src/net/http/request.go#L103)
- Use unnecessary pointer data like [URL](https://github.com/golang/go/blob/master/src/net/http/request.go#L124) in the Request structure.
- Unnecessary version [unmarshal](https://github.com/golang/go/blob/master/src/net/http/request.go#L132) in each request when [helper method](https://github.com/golang/go/blob/master/src/net/http/request.go#L399) exist, It isn't very performance penalty to just unmarshal and compare after in each request Because it uses very rare requirement.
- Declare some specific data in a request like [form data](https://github.com/golang/go/blob/master/src/net/http/request.go#L245) in this way why not have bodyAsJSON, ...
- When HTTP is a string-based protocol why use many integers like [response status code](https://github.com/golang/go/blob/master/src/net/http/server.go#L1134) to have many runtime base logic.
- net/http made [body](https://github.com/golang/go/blob/master/src/net/http/transfer.go#L809) to use by the concurrent situation, but it is bad logic to use the body in two different goroutines.
- So many unnecessary memory allocations, like allocate independently allocate for each header key and value in both [Request](https://github.com/golang/go/blob/master/src/net/http/request.go#L1076) and [Response](https://github.com/golang/go/blob/master/src/net/http/response.go#L191)
- Like many other libraries in the Go ecosystem, errors handle and declare in very bad shape and location. Error isn't about log message for developers, It is a message to show to users, So at least it must include locale messages, not just English.
- ...

## Why - github.com/valyala/fasthttp disadvantages
- high allocs/op even more than net/http in just parse phase. We know this package highly uses the pool to fool GC but it can't fool itself because anyway huge copy(on parse and reset logic) need to fill those allocations. even so many unneeded copy occur e.g. [URI.SetPath()](https://github.com/valyala/fasthttp/blob/3ff6aaa5917f40eeb5cdcb4272c58210f161f0ea/uri.go#L177) or [all header values](https://github.com/valyala/fasthttp/blob/7eeb00e1ccc54b29a6a165c6a27d5dfa96b416ca/header.go#L339)
- Anything is byte slice, but RFC says HTTP is string base protocol. It is so easy to not copy buffers but change type from byte slice to string. We know it is needed unsafe package, but it is worth to use unsafe package and easily can be contracted by package and consumers to know they must copy any data if they want to hold for a long time, otherwise, a memory leak occurs easily by any mistake logics.
- Implement many logics that don't belong to the HTTP layer:
    - It has worker mechanism just to limit the number of concurrent and in [some way](https://github.com/valyala/fasthttp/blob/9f11af296864153ee45341d3f2fe0f5178fd6210/workerpool.go#L147) it will block itself. Is it a good layer to handle such low layer requirements?
    - How data came to package as application layer such as TCP, TLS, ...
    - Limiting such as IP limiting does not belong to the HTTP package.
- Like many other libraries in the Go ecosystem, errors handle and declare in very bad shape and location. Error isn't about log message for developers, It is a message to show to users, So at least it must include locale messages, not just English.
- Use too much pool! e.g. [`uriPool`](https://github.com/valyala/fasthttp/blob/bcf7e8e94422403a93145c6ae7a8eb2224e0436b/uri.go#L30) used in locations where it just be to reuse underling `fullURI []byte` buffer, Due to `URI` can be on the stack. No benchmark provide to show is it worth to use pool?
- ...

## Goals
- Respect [HTTP protocols](../../protocol/http.go) and implement all requirements to be comply with those interfaces as http rules. So any other packages can easily be changeable with this package.

## Not Goals

## Protocols
- HTTP/1 : https://www.rfc-editor.org/rfc/rfc9110
- HTTP/2 : https://www.rfc-editor.org/rfc/rfc9113
- HTTP/3 : https://www.rfc-editor.org/rfc/rfc9114

# Abbreviations
