# Log to file [![Go Reference](https://pkg.go.dev/badge/github.com/Mydao57/logtofile.svg)](https://pkg.go.dev/github.com/Mydao57/logtofile)

Save your logs to a file to keep track of your project.

## Installation

You can use the `go` tool to install `logtofile`:

	go get github.com/Mydao57/logtofile

## Usage

In your Go app you can call the logger like this:

```go
package main

import (
	"log"
	"github.com/Mydao57/logtofile"
)

func init() {
	err := logger.NewLogger("./logs")
	if err != nil {
		log.Fatalf("could not create logger: %v", err)
	}
}

func main() {
	defer logger.Close()

	logger.INFO("This is an info message from main")
}
```

## Contributing

Contributions are welcome. Don’t hesitate to make a PR of your features