# libgo
`libgo` abbreviation of `Go language library` is a repository wrapper to store all implementation of SabzCity and others protocols and algorithms to make a digital platform more easily in better performance!
You can use all packages exported APIs and also use [library CLI](./lib-cli) to access some useful APIs from command line easily!

## Some useful git command
- Make project version control by ```git init```
- Clone exiting repo by ```git clone ${repository path}```.
- Add libgo to project as submodule by ```git submodule add -b master https://github.com/SabzCity/libgo```
- Clone existing project with just latest commits not all one ```git clone ${repository path} --recursive --shallow-submodules```
- Change libgo version by ```git checkout tag/${tag}``` or update by ```git submodule update -f --init --remote --checkout --recursive``` if needed.
