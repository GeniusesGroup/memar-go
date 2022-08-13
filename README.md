# libgo   
[![GoDoc](https://pkg.go.dev/badge/github.com/GeniusesGroup/libgo)](https://pkg.go.dev/github.com/GeniusesGroup/libgo)
[![Go Report](https://goreportcard.com/badge/github.com/GeniusesGroup/libgo)](https://goreportcard.com/report/github.com/GeniusesGroup/libgo)

An **application** developing **framework** provide ZeroOps(zero operations), edge computing, ... that let you develope both server and client applications in Go without need to think more about any fundamental requirements, Just develope business services and user interfaces (now just graphical interface - gui), build apps as OS images or OS applications and easily just run first server node and let it distributes by many factors with inside logics not need external decision makers (automating software deployment) like Kubernetes(K8s).

In other word, `libgo` abbreviation of `Go language library` is a repository wrapper to store all implementation of GeniusesGroup and others protocols and algorithms to make a digital software more easily in better performance.
You can use all packages exported APIs, go generator mechanism or by [library CLI](#CLI) to access some useful APIs from command line easily.

## Goals
- Provide complete framework to develope any purpose distributed application with **low||no code**.
- No(minimum) dependency on any other repositories.
- Compile an application as **Unikernel** instead of regular OSs binaries.
- Develope high available and distributed software without any admin in any infrastructure layers (DevOps culture goal).
- Let service developers act Lean and Agile in their organization.
- [Reinvent the wheel](https://en.wikipedia.org/wiki/Reinventing_the_wheel)

## Not Goals
- 

## Installation
- Make project directory and suggest use your internet domain name for it.
- initialize project version control. If you use git run `git init` or `git clone ${repository path}`.
- init go project by `git init` and `go mod init`
- add `libgo` to your project dependency
- install `libgo` with `lang_eng` or your desire language
- initialize the project with desire domain e.g. `google.com`
- build your organization app by `libgo build` or by `go build` with desire tags and target OS and hardwares.
- run your desire version from /bin/ directory. Strongly suggest run app by systemd on linux or other app manager on other OS.

or easily run the following command under your project, just replace your domain name.
```
mkdir {domain}
git init
go mod init
go get -u github.com/GeniusesGroup/libgo
go install github.com/GeniusesGroup/libgo -tags "lang_eng"
libgo app init -d={domain}
libgo build
```

## Commands (CLA)
libgo has a the command-line client for the some generator APIs implement in [modules](./modules/). It provides simple access to all APIs functions to make an application, a GUI app, ....

You can get list of all commands and their helps with `libgo help`. We just list some of important commands here that you can run them from within a Go module or any where in your project directory:
- **Initialize a project:** `libgo app init -d=[internet-domain]`
- **Add new domain module:** `libgo domain new -n=[domain-name]`

## Build tags
- **dev_mode**: first check and change `AppMode_Dev` const in protocol package to desire behavior
- **tcp_listener**:

### Developing hints
- Complete manifest in main package of service.
- Add other data to main package if needed.
- Implement protocols logic by autogenerate some codes not write them manually.
- Don't think about network when you develope a business service logic. Use `st protocol.Stream` as stream data not network stream and don't remove it even don't need it from `Process` method arguments.

## Protocols
You can find protocol descriptions in its directory as now [protocol](./protocol/), [society](./society/), [ISO](./iso/)    
Read more about each protocol or library in its [RFC](https://github.com/GeniusesGroup/RFCs)
As [suggest here](https://github.com/golang/go/issues/48087) to comply with the standards we add [protocol](./protocol) package and all other libgo packages implement this package. You can implement these protocols in your own way if our packages can't satisfied you or your requirements.   
A standard is a formalized protocol accepted by most of the parties that implement it. A protocol is not a set of rules. A protocol is the thing those rules describe the rules of. This is why programs implement a protocol and comply with a standard.

### Industry Protocols
- Insurance     >> ACCORD
- Health Care   >> HL7
- Retail        >> GS1
- HR            >> HRXML

## GIT
Git is not the best version control mechanism for a software project, but it is the most use one.

### Some useful commands
- Make project version control by ```git init```
- Clone exiting repo by ```git clone ${repository path}```.
- Add libgo to project as submodule by ```git submodule add -b master https://github.com/GeniusesGroup/libgo```
- Clone existing project with just latest commits not all one ```git clone ${repository path} --recursive --shallow-submodules```
- Change libgo version by ```git checkout tag/${tag}``` or update by ```git submodule update -f --init --remote --checkout --recursive``` if needed.

## Go
### Some useful commands
- go build -race
- go tool compile -S {{file-name}}.go > {{file-name}}_C.S
- go tool objdump {{file-name}}.o > {{file-name}}_O.S
- go build -gcflags=-m {{file-name}}.go
- go run -gcflags='-m -m' {{file-name}}.go
- go build -ldflags "-X version=0.1"

### build tags
Some functionality in files that have build tags `//go:build tag_name` or `// +build tag_name` in the first line just build when you provided in build time like `go build -tags "dev_mode tcp_listener"`. Build tag declarations must be at the very top of a .go files. Nothing, not even comments, can be above build tags. We prefer `go:build` over `+build` because as [describe in the proposal](https://go.googlesource.com/proposal/+/master/design/draft-gobuild.md#transition) but below chart is how to declare in `+build` style.

| Build Tag Syntax	            | Build Tag Sample	                        | Boolean Statement     |
| :---:                         | :---:                                     | :---:                 |
| Space-separated elements	    | // +build pro enterprise	                | pro OR enterprise     |
| Comma-separated elements      | // +build pro,enterprise	                | pro AND enterprise    |
| New line separated elements   | // +build pro<br />// +build enterprise   | pro AND enterprise    |
| Exclamation point elements    | // +build !pro	                        | NOT pro               |

## Contribute Rules
- Write benchmarks and tests codes in different files as `{{file-name}}_test.go` and `{{file-name}}_bench_test.go`

## Enterprise
Contact us by [this](mailto:ict@geniuses.group) or [this](mailto:omidhekayati@gmail.com) if you need enterprise support for developing high available and distributed software. See features available in enterprise package:
- Develope exclusive features in very short time
- Bug fixing quickly
- 

## Good Idea, Bad implementation!
- [SQLc](sqlc.dev)
- [EntGo](https://entgo.io/)
- [go-zero](https://github.com/zeromicro/go-zero) e.g. (microservice system), (fully compatible with net/http), (middlewares are supported), ...
or [really relativetime?? Why not monotonic time??](https://github.com/zeromicro/go-zero/blob/master/core/timex/relativetime.go)

## Related Projects
- [Clive is an operating system designed to work in distributed and cloud computing environments.](https://github.com/fjballest/clive)

## Abbreviations & Definitions
- **UI**: (any) User Interface
    - **GUI**: Graphic User Interface
    - **VUI**: Voice User Interface
    - **CLI**: Command Line Interface
    - **CLA**: Command Line Arguments
- **Modules**: a kind of collection of packages
- **Packages**: a kind of collection of files
- **dp**: domain protocol
- **init**: initialize call just after an object allocate.
- **reinit**: re-initialize call when allocated object want to reuse immediately or pass to a pool to reuse later. It will prevent memory leak by remove any references in the object.
- **deinit**: de-initialize call just before an object want to de-allocated (GC).
- **open**:
- **reset**:
- **close**:
