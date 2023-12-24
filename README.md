# tide-go

## Introduction

tide-go is a flexible config parser for Go, designed to handle the TIDE configuration 
format. It provides an easy-to-use API for accessing configuration values from
TIDE files.

See [TIDE](https://github.com/tideconf/tide) for more information on the TIDE
configuration format.

> [!IMPORTANT]  
> This is no more than a hobby project at the moment. I have always been curious about the design and implementation of configuration frameworks, and this is my attempt at creating one. I am not sure if this will ever be used, but I am hoping to learn a bit more about the whole deal of configuration handling. If you are interested in this project, please feel free to contribute or provide feedback.

# Installation

To install tide-go, use `go get`:

```go
go get -u github.com/tideconf/tide-go
```

# Example

An example of using tide-go to parse a TIDE file, is available in the examples
directory.

```go
go run examples/example.go 
```

# Development

This is only capable of parsing TIDE files at the moment. It might be expanded
to support function calls / embedded logic in the future.
