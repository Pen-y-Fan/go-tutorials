# Create a Go module

Source: [Tutorial: Create a Go module](https://go.dev/doc/tutorial/create-module)

## About

This tutorial's sequence includes seven brief topics that each illustrate a different part of the Go language.

1. [x] Create a module -- Write a small module with functions you can call from another module.
2. [x] Call your code from another module -- Import and use your new module.
3. [x] Return and handle an error -- Add simple error handling.
4. [x] Return a random greeting -- Handle data in slices (Go's dynamically-sized arrays).
5. [x] Return greetings for multiple people -- Store key/value pairs in a map.
6. [x] Add a test -- Use Go's built-in unit testing features to test your code.
7. [x] Compile and install the application -- Compile and install your code locally.

## Test

To run the test

```sh
cd greetings
go test -v
```

You will see the verbose output (-v is optional)

```text
=== RUN   TestHelloName
--- PASS: TestHelloName (0.00s)
=== RUN   TestHelloEmpty
--- PASS: TestHelloEmpty (0.00s)
=== RUN   TestHellosNames
--- PASS: TestHellosNames (0.00s)
=== RUN   TestHellosEmpty
--- PASS: TestHellosEmpty (0.00s)
PASS
ok      example.com/greetings   0.205s
```
