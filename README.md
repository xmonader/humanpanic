# humanpanic

`humanpanic` is a Go package designed to handle panics caused by human errors gracefully. It captures panic information, logs it, and provides a user-friendly message to aid in diagnosing the problem. It's inspired by [human-panic crate](https://github.com/rust-cli/human-panic)

## Features

- Captures panics with detailed stack trace and file information.
- Logs error details to a temporary file for further analysis.
- Provides functions to trigger panics with custom messages.

## Installation

To install the `humanpanic` package, use the following command:

```sh
go get github.com/xmonader/humanpanic
```

## Usage

### Importing the Package

```go
import "github.com/xmonader/humanpanic"
```

### Triggering a Panic

To trigger a panic with a custom message, use the `Panic` or `Panicf` functions.

```go
package main

import (
    "github.com/xmonader/humanpanic"
)

func main() {
    defer func() {
        if err := humanpanic.Recover(); err != nil {
            // Handle recovered error
        }
    }()

    // Trigger a panic with a custom message
    humanpanic.Panic("Something went wrong!")
}
```

### Recovering from a Panic

To recover from a panic and log the error details, use the Recover function.


```go
package main

import (
    "fmt"
    "github.com/xmonader/humanpanic"
)

func main() {
    defer func() {
        if err := humanpanic.Recover(); err != nil {
            fmt.Printf("Recovered from panic: %v\n", err)
        }
    }()

    // Trigger a panic
    humanpanic.Panic("This is a test panic!")
}
```
