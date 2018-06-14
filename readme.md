# go-healthchecks [![Build Status](https://travis-ci.org/frozzare/go-healthchecks.svg?branch=master)](https://travis-ci.org/frozzare/go-healthchecks) [![GoDoc](https://godoc.org/github.com/frozzare/go-healthchecks?status.svg)](http://godoc.org/github.com/frozzare/go-healthchecks) [![Go Report Card](https://goreportcard.com/badge/github.com/frozzare/go-healthchecks)](https://goreportcard.com/report/github.com/frozzare/go-healthchecks)

Max is a YAML-based task runner with Docker support.

Simple http client api for [healthchecks.io](https://healthchecks.io)

## Installation

```
go get -u github.com/frozzare/go-healthchecks
```

## Example

```
package main

import (
	"context"
	"log"

	"github.com/frozzare/go-healthchecks"
)

func main() {
	client := healthchecks.NewClient(nil)

	err := client.Success(context.Background(), "your-uuid-here")

	if err != nil {
		log.Fatal(err)
	}

	err = client.Fail(context.Background(), "your-uuid-here")

	if err != nil {
		log.Fatal(err)
	}
}
```

## License

MIT © [Fredrik Forsmo](https://github.com/frozzare)