
# Project Title

This my sandbox project for learning Golang with Beego framework. I am using bee tool to manage my project.


## Installation

After cloning the repository to run the project you must have golang installed on your computer. Recommended version is 18+. Link for Go install: https://go.dev/doc/install

(Optional) Now create workspace, so the VSCode can reference the needed dependencies. I found this very useful, when working with VSCode. This will generate a go.work file, where all module names will be stored. When creating a new module just remember to add `go work use $DIRECTORY` to add directory to workspace. Now go into parent directory and run commands given below.

```bash
    go work init
    go work use go-car-rental
```

Lastly, go into project directory and run

```bash
    go mod tidy
```

This will download all the dependencies needed and generate go.sum file. Now we can build and run the server using standard go commands.

```bash
    go build main.go
    ./go-car-rental
```

As an alternative you can install bee tool and use it to run project. More information on installing bee tool: https://beego.gocn.vip/beego/en-US/developing/bee/. Use `go install` instead of `go get` for newer versions of Golang.