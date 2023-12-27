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

```bash
go get -u github.com/tideconf/tide-go
```

# Usage

## Parsing a TIDE file

To parse a TIDE file, use the `ParseFile` function:

```bash
$ cat config.tide
database {
    type: string = "mysql"
    host: string = "localhost"
    port: integer = 3306
    credentials {
        username: string = "user"
        password: string = "pass"
    }
}

myApp {
    features: array[string] = ["feature1", "feature2", "feature3"]
    numbers: array[integer] = [1, 2, 3]
}
```

```go
package main

import (
    "fmt"
    "github.com/tideconf/tide-go"
)

func main() {
    cfg, err := tide.NewTIDE("./path/to/config.tide")
	if err != nil {
		log.Fatal(err)
	}

    port, err := cfg.GetInt("database.port")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Port:", port)
}
```

## Importing a TIDE file into another TIDE file

It is possible to merge configurations, with the imported configuration taking
precedence over the base configuration in case of conflicts. This is useful for
splitting up configuration into multiple files.

Circular imports are not allowed.

To import a TIDE file into another TIDE file, use the `ImportFile` function:

```bash
$ cat config.tide
import "additional_settings"

database {
    host: string = "localhost"
    port: integer = 3306
}

$ cat additional_settings.tide
logging {
    level: string = "info"
}
```

## Environment variables

TIDE configuration values can be overridden by environment variables. The
environment variable name is the uppercased path to the configuration value,
with the path separator replaced by an underscore.

For example `database.credentials.username` would be overridden by the
`DATABASE_CREDENTIALS_USERNAME` environment variable.

# Example

An example of using tide-go to parse a TIDE file, is available in the examples
directory.

```go
go run examples/example.go 
```

# Contributing

Contributions are welcome! Please feel free to submit a pull request or open an
issue if you find a bug or have a feature request.

# License

tide-go is licensed under the Apache 2.0 license. See the LICENSE file for more
information.