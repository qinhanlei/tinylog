Tiny simple naive

## Usage
`go get github.com/qinhanlei/tinylog`

#### example

```golang
package main

import (
	log "github.com/qinhanlei/tinylog"
)

func init() { log.Init(".") }

func main() {
	log.Debug("Hello Debug")
	log.Info("Hello Info")
	log.Warn("Hello Warn")
	log.Error("Hello Error")
	log.Fatal("Hello Fatal")
}
```

#### output
<!-- ![](https://raw.githubusercontent.com/qinhanlei/tinylog/master/test.png) -->
<div align="left"><img width="500" height="90" src="https://raw.githubusercontent.com/qinhanlei/tinylog/master/test.png"/></div>
