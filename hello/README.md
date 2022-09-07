# Hello, world

Source: [go.dev > Tutorial: Get started with Go](https://go.dev/doc/tutorial/getting-started)

## About

This is the hello world must have tutorial, with a difference! The tutorial also explains how to import a module and call a public function from that module.

## Installation

Install the modules:

```sh
go mod tidy
```

Output:

```sh
go: finding module for package rsc.io/quote
go: found rsc.io/quote in rsc.io/quote v1.5.2
go: found example.com/greetings in example.com/greetings v0.0.0-00010101000000-000000000000
```

## Run

Run the hello.go program

```sh
go run .
```

Output:

```text
Hello, World!
Don't communicate by sharing memory, share memory by communicating.
Great to see you, Gladys!
Messages:
➡️  Hail, Gladys! Well met!
➡️  Hail, Samantha! Well met!
➡️  Hi, Darrin. Welcome!
```

To build an executable to your `GOBIN` executable directory (should be set in PATH)

```sh
# Check the project will build:
go build
hello.exe
# check if GOBIN is already set:
go env | grep GOBIN
# if GOBIN isn't set, it can be set:
go env -w GOBIN=C:\path\to\your\bin
# e.g.:
# go env -w GOBIN=C:\laragon\usr\bin
# Install the executable hello.exe
go install
```

The `hello.exe` will be executable from the directory set in GOBIN
