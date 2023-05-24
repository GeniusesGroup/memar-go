# libgo Best Practice

### Developing hints
- Don't think about network when you develope a business service logic. Use `st protocol.Stream` as stream data not network stream and don't remove it even don't need it from `Process` method arguments.

## Protocols
You can find protocol descriptions in its directory as now [protocol](./protocol/), [ISO](./iso/)    
Read more about each protocol or library in its [RFC](https://github.com/GeniusesGroup/RFCs)
As [suggest here](https://github.com/golang/go/issues/48087) to comply with the standards we add [protocol](./protocol) package and all other libgo packages implement this package. You can implement these protocols in your own way if our packages can't satisfied you or your requirements.   
A standard is a formalized protocol accepted by most of the parties that implement it. A protocol is not a set of rules. A protocol is the thing those rules describe the rules of. This is why programs implement a protocol and comply with a standard.

## Optimize memory allocation and de-allocation
- use `go build -gcflags=-m {{file-name}}.go` to determine `heap escape`

### build tags
Some functionality in files that have build tags `//go:build tag_name`

## Contribute Rules
- Write benchmarks and tests codes in different files as `{{file-name}}_test.go` and `{{file-name}}_bench_test.go`
- Naming by [common naming convention](https://en.wikipedia.org/wiki/Naming_convention_(programming))

## Structure methods
All structure declare in libgo should have these methods:
- **Init**: initialize call just after an object allocate.
- **Reinit**: re-initialize call when allocated object want to reuse immediately or pass to a pool to reuse later. It will prevent memory leak by remove any references in the object.
- **Deinit**: de-initialize call just before an object want to de-allocated (GC).

All structure declare in libgo can have these methods too:
- **Alloc**: call to allocate the object. It must be called on pointer to the struct not direct use e.g. `t *embed`.
- **Dealloc**: call to de-allocate the object. It must be called on pointer to the struct not direct use e.g. `t *embed`.
- **clone**:
- **open**:
- **reset**:
- **close**:

## Rules
- Use `var` keyword when declare new variable instead of `:=`. Just use `:=` in `for`, `range`, `switch .(type)`
- Write one logic per line and avoid to use `;` to write multiple logic in one line e.g.
```go
    if err := Logic(); err != nil {
        return
    }

    var err = Logic()
    if err != nil {
        return
    }
```
- Don't access fields in any way, just use methods to access structure fields.

## GIT
Git is not the best version control mechanism for a software project, but it is the most use one.

### Some useful commands
- Make project version control by ```git init```
- Clone exiting repo by ```git clone ${repository path}```.
- Add libgo to project as submodule by ```git submodule add -b master https://github.com/GeniusesGroup/libgo```
- Clone existing project with just latest commits not all one ```git clone ${repository path} --recursive --shallow-submodules```
- Change libgo version by ```git checkout tag/${tag}``` or update by ```git submodule update -f --init --remote --checkout --recursive``` if needed.

## Methods logic
- **Get**: get primary key and return full record data
- **Find**: get secondary key and return primary keys
- **Filter**: get at least one key that one of them is not index and must do some logic in runtime and usually(not always) return primary keys.
- **List**: get primary key and return secondary keys
- **Search**: 

## Abbreviations & Definitions
- **UI**: (any) User Interface
    - **GUI**: Graphic User Interface
    - **VUI**: Voice User Interface
    - **CLI**: Command Line Interface
    - **CLA**: Command Line Arguments
- **Modules**: a kind of collection of packages
- **Packages**: a kind of collection of files
- **dp**: domain protocol
- **init**: initialize an object
- **reinit**: re-initialize to reuse an object
- **deinit**: de-initialize an object
- **open**:
- **reset**:
- **close**:
- **SCM**: Source control management (SCM) systems provide a running history of code development and help to resolve conflicts when merging contributions from multiple sources.
