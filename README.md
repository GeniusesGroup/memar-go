# libgo   
[![GoDoc](https://pkg.go.dev/badge/github.com/GeniusesGroup/libgo)](https://pkg.go.dev/github.com/GeniusesGroup/libgo)
[![Go Report](https://goreportcard.com/badge/github.com/GeniusesGroup/libgo)](https://goreportcard.com/report/github.com/GeniusesGroup/libgo)

An **application** developing **framework** provide ZeroOps(zero operations), edge computing, ... that let you develope both server and client applications in Go without need to think more about any fundamental requirements, Just develope business services and user interfaces (now just graphical interface - gui), build apps as OS images or OS applications and easily just run first server node and let it distributes by many factors with inside logics not need external decision makers (automating software deployment) like Kubernetes(K8s).

In other word, `libgo` abbreviation of `Go language library` is a repository wrapper to store all implementation of GeniusesGroup and others protocols and algorithms to make a digital software more easily in better performance.
You can use all packages exported APIs, go generator mechanism or by [library commands](#commands-cla) to access some useful APIs from command line easily.

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
- init go project by `go mod init`
- add `libgo` to your project dependency
- install `libgo` with `lang_eng` or your desire language
- initialize the project with desire domain e.g. `google.com`
- Complete manifest in main package.
- Add other data to main package if needed.
- Implement protocols logic by autogenerate some codes not write them manually.
- build your organization app with desire tags and target OS and hardwares.
- run your desire version from /bin/ directory. Strongly suggest run app by systemd on linux or other app manager on other OS.

or easily run the following command under your project, just replace your domain name.
```
mkdir {domain}
git init
go mod init
go get -u github.com/GeniusesGroup/libgo
go install github.com/GeniusesGroup/libgo -tags "lang_eng"
libgo app init -d={domain}
libgo app build
```

## Commands (CLA)
libgo has a the command-line client for the some generator APIs implement in [modules](./modules/). It provides simple access to all APIs functions to make an application, a GUI app, ....

You can get list of all commands and their helps with `libgo help`. We just list some of important commands here that you can run them from within a Go module or any where in your project directory:
- **Initialize a project:** `libgo app init -idn=[internet-domain-name]`
- **Add new domain module:** `libgo mod new -dn=[domain-name]`
- **Build the apps(os images):** `libgo app build`
- **Run the app(os image):** `libgo app run`

## Build tags
- **dev_mode**: first check and change `AppMode_Dev` const in protocol package to desire behavior
- **tcp_listener**:

## Enterprise
Contact us by [this](mailto:ict@geniuses.group) or [this](mailto:omidhekayati@gmail.com) if you need enterprise support for developing high available and distributed software. See features available in enterprise package:
- Develope exclusive features in very short time
- Bug fixing quickly
- 

## Related Projects
- [Clive is an operating system designed to work in distributed and cloud computing environments.](https://github.com/fjballest/clive)
- [SQLc](sqlc.dev)
- [EntGo](https://entgo.io/)
- [go-zero](https://github.com/zeromicro/go-zero) e.g. (microservice system), (fully compatible with net/http), (middlewares are supported), ...
or [really relativetime?? Why not monotonic time??](https://github.com/zeromicro/go-zero/blob/master/core/timex/relativetime.go)


## Code style
[Read more here](./BEST-PRACTICE.md)
