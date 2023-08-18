# Memar - Golang version   
[![GoDoc](https://pkg.go.dev/badge/github.com/GeniusesGroup/memar-go)](https://pkg.go.dev/github.com/GeniusesGroup/memar-go)
[![Go Report](https://goreportcard.com/badge/github.com/GeniusesGroup/memar-go)](https://goreportcard.com/report/github.com/GeniusesGroup/memar-go)

`Memar` as a `Go language library` is a repository wrapper to store all implementation of [Memar](https://github.com/GeniusesGroup/memar) and others protocols and algorithms to make a digital software more easily in better performance.
You can use all packages exported APIs, go generator mechanism or by [library commands](#commands-cla) to access some useful APIs from command line easily.

If you want to get insight about memar, You MUST start reading interfaces in [protocol package](./protocol/).

## Installation
- Make project directory and suggest use your internet domain name for it.
- initialize project version control. If you use git run `git init` or `git clone ${repository path}`.
- init go project by `go mod init`
- add `memar` to your project dependency
- install `memar` with `lang_eng` or your desire language
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
go get -u github.com/GeniusesGroup/memar-go
go install github.com/GeniusesGroup/memar-go -tags "lang_eng"
memar app init -d={domain}
memar app build
```

## Commands (CLA)
memar has a the command-line client for the some generator APIs implement in [modules](./modules/). It provides simple access to all APIs functions to make an application, a GUI app, ....

You can get list of all commands and their helps with `memar help`. We just list some of important commands here that you can run them from within a Go module or any where in your project directory:
- **Initialize a project:** `memar app init -idn=[internet-domain-name]`
- **Add new domain module:** `memar mod new -dn=[domain-name]`
- **Build the apps(os images):** `memar app build`
- **Run the app(os image):** `memar app run`

## Build tags
- **dev_mode**: first check and change `AppMode_Dev` const in protocol package to desire behavior
- **tcp_listener**:

## Code style
[Read more here](./BEST-PRACTICE.md)

## Contribution Guide
- [Referencing issues](https://go.dev/doc/contribute#ref_issues)
