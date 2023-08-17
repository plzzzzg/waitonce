# WaitOnce

WaitOnce is a simple tool that ensures the prerequisites are ready.

## Install

```shell
go insall github.com/plzzzzg/waitonce
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
		if timeout := waitonce.GetOrCreateWaitOnce(waitOnceID).Wait(time.Second); timeout {
			// fallback when fallback
		} else {
			// do something after preloading done
		}
	}()

	// preload async
	go func() {
		// preloading
		time.Sleep(time.Second)
		waitonce.GetOrCreateWaitOnce(waitOnceID).Done()
	}()
}

```

## Licence

Licensed under the MIT License.