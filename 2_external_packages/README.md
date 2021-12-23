Initiate module first with `go mod init test_module`
To get an external package: `go get github.com/badoux/checkmail`

If module is not initiated you need to have GOPATH variable to indicate where to install the package, if not defined then just define a place for it: `$export GOPATH=$HOME/go`

To remove unused dependencies from go.mod: `go mod tidy`