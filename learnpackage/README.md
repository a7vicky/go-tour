

The line of code package main specifies that this file belongs to the main package. The import "packagename" statement is used to import an existing package. packagename.FunctionName() is the syntax to call a function in a package.

$ go mod init learnpackage
The above command will create a file named go.mod. The following will be the contents of the file.

```
├── learnpackage
│   ├── go.mod
│   ├── main.go
│   └── simpleinterest
│       └── simpleinterest.go
```
