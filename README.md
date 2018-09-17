Tiny simple naive

## Usage
`go get github.com/qinhanlei/tinylog`

```golang
package main

import (
	log "github.com/qinhanlei/tinylog"
)

func main() {
	log.Init(".")
	log.Debug("Hello Debug")
	log.Error("Hello Error")
}
```
