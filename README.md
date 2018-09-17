# tinylog
simple and stupid logger


## Usage
`got get github.com/qinhanlei/tinylog`

Colored for linux/macOS only now


## Example
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


## Todo
rolling log file
