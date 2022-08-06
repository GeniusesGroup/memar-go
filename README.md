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
- init go project
- add `libgo` to your project dependency
- install `libgo` with `lang_eng` or your desire language
- initialize the project with desire domain e.g. `google.com`
- build your organization app
- run your desire version from /bin/ directory

or easily run the following command under your project:
```
go mod init
go get -u github.com/GeniusesGroup/libgo
go install github.com/GeniusesGroup/libgo -tags "lang_eng"
libgo app init -d=google.com
libgo build
```

## Most common libgo commands
From within a Go module or any where in your project directory:
- run `libgo app init -d=[internet-domain]`
- To add new domain run `libgo domain new -n=[domain-name]`


## Protocols
You can find protocol descriptions in its directory as now [protocol](./protocol/), [society](./society/), [ISO](./iso/)    
Read more about each protocol or library in its [RFC](https://github.com/GeniusesGroup/RFCs)
As [suggest here](https://github.com/golang/go/issues/48087) to comply with the standards we add [protocol](./protocol) package and all other libgo packages implement this package. You can implement these protocols in your own way if our packages can't satisfied you or your requirements.   
A standard is a formalized protocol accepted by most of the parties that implement it. A protocol is not a set of rules. A protocol is the thing those rules describe the rules of. This is why programs implement a protocol and comply with a standard.

## Industry Protocols
- Insurance     >> ACCORD
- Health Care   >> HL7
- Retail        >> GS1
- HR            >> HRXML

## Go build tags
Some functionality in files that have build tags `//go:build tag_name` or `// +build tag_name` in the first line just build when you provided in build time like `go build -tags "dev_mode tcp_listener"`. Build tag declarations must be at the very top of a .go files. Nothing, not even comments, can be above build tags. We prefer `go:build` over `+build` because as [describe in the proposal](https://go.googlesource.com/proposal/+/master/design/draft-gobuild.md#transition) but below chart is how to declare in `+build` style.

| Build Tag Syntax	            | Build Tag Sample	                        | Boolean Statement     |
| :---:                         | :---:                                     | :---:                 |
| Space-separated elements	    | // +build pro enterprise	                | pro OR enterprise     |
| Comma-separated elements      | // +build pro,enterprise	                | pro AND enterprise    |
| New line separated elements   | // +build pro<br />// +build enterprise   | pro AND enterprise    |
| Exclamation point elements    | // +build !pro	                        | NOT pro               |

- **dev_mode** :
- **tcp_listener** :

## Some useful git commands
- Make project version control by ```git init```
- Clone exiting repo by ```git clone ${repository path}```.
- Add libgo to project as submodule by ```git submodule add -b master https://github.com/GeniusesGroup/libgo```
- Clone existing project with just latest commits not all one ```git clone ${repository path} --recursive --shallow-submodules```
- Change libgo version by ```git checkout tag/${tag}``` or update by ```git submodule update -f --init --remote --checkout --recursive``` if needed.

## Some useful GO commands
- go build -race
- go tool compile -S {{file-name}}.go > {{file-name}}_C.S
- go tool objdump {{file-name}}.o > {{file-name}}_O.S
- go build -gcflags=-m {{file-name}}.go
- go run -gcflags='-m -m' {{file-name}}.go
- go build -ldflags "-X version=0.1"

## CLI - Command-Line Interface
lib-cli is the command-line client for the some generator APIs implement in [services](./services/). It provides simple access to all APIs functions to make a server, a GUI app & ....

### Make new project - Use git as version control
- Make project folder and suggest use your domain name for it.
- Make project version control by `git init`
- Instead aboves you can clone exiting repo by `git clone ${repository path}`.
- Install lib-cli by `go install github.com/GeniusesGroup/libgo/libgo/`
- Run lib-cli in a terminal by `libgo --desire-command`
- Choose desire services to make needed files or other actions.

### APIs
- Complete manifest in main package of service.
- Add other data to main package if needed.
- Add as many service you need by CLI services and add business logic to them.
- From CLI update service file to autogenerate some code for you.
- As you can see in file services logic layers are independent layer and you must just think locally. But if you need network stream data use `st protocol.Stream` in your each function parameters. Don't remove it even don't need it.

## RUN
- first check and change `AppMode_Dev` const in protocol package to desire behavior
- Build app by `go build`
- Strongly suggest run app by systemd on linux or other app manager on other OS.
- Otherwise easily run app by `./{{root-folder-name}}`


## Contribute Rules
- Write benchmarks and tests codes in different files as `{{file-name}}_test.go` and `{{file-name}}_bench_test.go`

# Enterprise
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

# Abbreviations & Definitions
- UI >> any User Interface  --  GUI >> Graphic User Interface  --  VUI >> Voice User Interface  --  CLI >> Command Line Interface
- **Modules**: a kind of collection of packages
- **Packages**: a kind of collection of files
- **dp**: domain protocol