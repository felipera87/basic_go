To run a file: `go run main.go`
The application need to have main package and main function, not necessary to packages

To create a module: `go mod init test_module`
Module is equivalent to package.json on node

To generate the binaries of the module: `go build`
Then you can run it directly on terminal: `./test_module`

To "install" the module to golang root: `go install`
It'll do the same as build, but on golang directories.
