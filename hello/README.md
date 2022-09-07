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
Hi, Gladys. Welcome!
or Great to see you, Gladys!
or Hail, Gladys! Well met!
greetings: empty name
exit status 1
```
