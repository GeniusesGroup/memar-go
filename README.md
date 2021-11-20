# libgo
`libgo` abbreviation of `Go language library` is a repository wrapper to store all implementation of GeniusesGroup and others protocols and algorithms to make a digital platform more easily in better performance!
You can use all packages exported APIs and also use [library CLI](./lib-cli) to access some useful APIs from command line easily!

## Goals
- Compile an application as **Unikernel** instead of regular OSs binaries.
- Develope high available and distributed software without any admin in any infrastructure layers (DevOps culture goal).
- Let service developers act Lean and Agile in their organization.

## Not Goals
- 

## Protocols
Read more about each protocol or library in its [RFC](https://github.com/GeniusesGroup/RFCs)
As [suggest here](https://github.com/golang/go/issues/48087) to comply with the standards we add [protocol](./protocol) package and all other libgo packages implement this package. You can implement these protocols in your own way if our packages can't satisfied you or your requirements!   
A standard is a formalized protocol accepted by most of the parties that implement it. A protocol is not a set of rules. A protocol is the thing those rules describe the rules of. This is why programs implement a protocol and comply with a standard.

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

# Enterprise
Contact us by [this](ict@sabz.city) or  [this](omidhekayati@gmail.com) if you need enterprise support for developing high available and distributed software. See features available in enterprise package:
- Develope exclusive features in very short time
- Bug fixing quickly
- 

# Abbreviations
- UI >> any User Interface  --  GUI >> Graphic User Interface  --  VUI >> Voice User Interface  --  CLI >> Command Line Interface
