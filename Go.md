# Go programming language

## Some useful commands
- go build -race
- go tool compile -S {{file-name}}.go > {{file-name}}_C.S
- go tool objdump {{file-name}}.o > {{file-name}}_O.S
- go build -gcflags=-m {{file-name}}.go
- go run -gcflags='-m -m' {{file-name}}.go
- go build -ldflags "-X version=0.1"

## build tags
Some functionality in files that have build tags `//go:build tag_name` or `// +build tag_name` in the first line just build when you provided in build time like `go build -tags "dev_mode tcp_listener"`. Build tag declarations must be at the very top of a .go files. Nothing, not even comments, can be above build tags. We prefer `go:build` over `+build` because as [describe in the proposal](https://go.googlesource.com/proposal/+/master/design/draft-gobuild.md#transition) but below chart is how to declare in `+build` style.

| Build Tag Syntax	            | Build Tag Sample	                        | Boolean Statement     |
| :---:                         | :---:                                     | :---:                 |
| Space-separated elements	    | // +build pro enterprise	                | pro OR enterprise     |
| Comma-separated elements      | // +build pro,enterprise	                | pro AND enterprise    |
| New line separated elements   | // +build pro<br />// +build enterprise   | pro AND enterprise    |
| Exclamation point elements    | // +build !pro	                        | NOT pro               |

## Vulnerability Management for Go
The [govulncheck command](https://go.dev/blog/vuln) is a low-noise, reliable way for Go users to learn about known vulnerabilities that may affect their projects. Govulncheck analyzes your codebase and only surfaces vulnerabilities that actually affect you, based on which functions in your code are transitively calling vulnerable functions. To start using govulncheck, you can run the following from your project:
```
$ go install golang.org/x/vuln/cmd/govulncheck@latest
$ govulncheck ./...
```

## Directive comments
- `//go:generate`: This directive comment is used to specify a command that should be run by the go generate tool. It is typically placed before a code generation command, allowing you to automatically generate Go code.
- `//go:binary-only-package`: This directive comment is used to indicate that a package should be treated as a binary-only package, meaning that the source code for the package is not available. This is often used for packages that contain proprietary or platform-specific code.
- `//go:build`: This directive comment is used for build constraints. It allows you to control whether a file should be included in the build based on certain conditions, such as the operating system, architecture, or build tags.
- `//go:cgo_...`: There are several directive comments that start with cgo_, such as //go:cgo_import_dynamic and //go:cgo_export_dynamic. These comments are used in conjunction with cgo, the tool that allows Go code to call C code and vice versa. They provide instructions to the cgo tool on how to handle the C code.
- `//go:noinline`: This directive comment is used to indicate that a function should not be inlined by the compiler. Inlining is an optimization technique where the code of a function is inserted directly into the calling code, eliminating the function call overhead. Using this directive comment prevents the compiler from performing inlining for the specified function.
- `//go:nosplit`: This directive comment is used to indicate that a function should not be preempted (split) by the Go runtime scheduler. It is typically used for low-level functions that need precise control over their execution and should not be interrupted.
- `//go:linkname`: This directive comment is used to establish a link between Go code and non-Go code or external symbols. It allows you to refer to a symbol by a different name or from a different package.
- `//go:noescape`: This directive comment is used to indicate that a function's pointer arguments do not escape, meaning they are not stored or used beyond the lifetime of the function. This information allows the compiler to optimize the function's memory allocation.
- `//go:embed`: This directive comment is used to include static files or directories directly into the Go binary at compile time. It simplifies the process of bundling and distributing resources with your Go programs.
- `//go:generate go run`: This directive comment is a variant of the //go:generate directive. It specifies that the command following the directive should be executed by running a Go program using the go run command.
- `//go:build ...`: This directive comment is an extended form of the //go:build directive. It allows you to specify build constraints with more complex conditions using boolean operators, parentheses, and negation. This provides greater flexibility in controlling which files are included in the build.
- `//go:protofile`: This directive comment is used to specify the protobuf file associated with a Go source file. It is typically used in Go code that includes generated protobuf code, allowing the compiler to link the Go and protobuf files together correctly.
- `//go:nowritebarrier`: This directive comment is used to indicate that a function should be executed without write barriers. Write barriers are used by the garbage collector to track and update pointers during memory allocation and deallocation. Using this directive can be risky and should only be used in certain cases where manual memory management is required.
- `//go:norace`: This directive comment is used to disable the race detector for a specific function. The race detector is a tool in Go that helps identify concurrent access to shared variables that could lead to data races. This directive can be used when you are confident that a particular function is free from race conditions.
- `//go:buildignore`: This directive comment is used to exclude a file from the build process. It tells the Go build system to ignore the file, and it won't be included when compiling the package.
- `//go:generate goimports`: This directive comment is used in conjunction with the //go:generate directive to automatically run the goimports tool. goimports automatically updates and formats import statements in Go code, ensuring proper package imports and removing unused imports.
- `//go:embed pattern`: This directive comment is an extended form of //go:embed and allows you to specify a pattern to match files or directories for embedding. It provides more flexibility in selecting specific files or directories to include.
- `//go:nolint`: This directive comment is used to suppress specific linter warnings or errors for a particular line of code. It is often used when a linter rule flags a false positive or when there is a valid reason to ignore a linting issue temporarily.
- `//go:generate go test`: This directive comment is used in conjunction with the //go:generate directive to automatically run the go test command. It is commonly used to generate and execute test code for a package.
- `//go:uintptrescapes`: This directive comment is used to indicate that a uintptr value can escape to the heap. By default, the compiler assumes that uintptr values do not escape, but using this directive allows for more accurate escape analysis.
- `//go:build !constraint`: This directive comment is used to exclude a file from the build process based on a specific build constraint. It allows you to specify a constraint that should not be satisfied for the file to be included in the build.
- `//go:checkptr`: This directive comment is used to enable additional pointer safety checks in the code. It instructs the compiler to generate extra runtime checks to catch unsafe pointer operations and potential memory safety issues.
- `//go:nosplitcheck`: This directive comment is used to disable a nosplit check for a function. The nosplit check verifies that a function can execute without being preempted by the Go runtime scheduler. Using this directive can be risky and should only be used when necessary.
- `//go:noruntime`: This directive comment is used to indicate that a package does not depend on the Go runtime. It informs the compiler that the package can be used in a context where the Go runtime is not available or required.

## Other Organizations Styles
- [Uber](https://github.com/uber-go/guide/blob/master/style.md)
