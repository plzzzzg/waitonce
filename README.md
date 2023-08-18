# WaitOnce

[![Go Reference](https://pkg.go.dev/badge/github.com/plzzzzg/waitonce#section-readme.svg)](https://pkg.go.dev/github.com/plzzzzg/waitonce#section-readme)
[![Go](https://github.com/plzzzzg/waitonce/actions/workflows/go.yml/badge.svg)](https://github.com/plzzzzg/waitonce/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/plzzzzg/waitonce)](https://goreportcard.com/report/github.com/plzzzzg/waitonce)
[![codecov](https://codecov.io/gh/plzzzzg/waitonce/graph/badge.svg?token=VG7XY6OXZG)](https://codecov.io/gh/plzzzzg/waitonce)

WaitOnce is a simple tool that ensures the prerequisites are ready.

## Install

```shell
go get github.com/plzzzzg/waitonce
```


## Examples

```go
package main

import (
	"github.com/plzzzzg/waitonce"
	"time"
)

func main() {
	waitOnceID := "preload"

	go func() {
		if timeout := waitonce.GetOrCreate(waitOnceID).Wait(time.Second); timeout {
			// fallback when timeout
		} else {
			// do something after preloading done
		}
	}()

	// preload async
	go func() {
		// preloading
		time.Sleep(time.Second)
		waitonce.GetOrCreate(waitOnceID).Done()
	}()
}

```

## Licence

Licensed under the MIT License.