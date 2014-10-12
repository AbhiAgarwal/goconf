go-conf
=======

nconf inspired configuration for go

Simple usage:

```go
package main

import (
	"fmt"
	"github.com/abhiagarwal/goconf"
)

func main() {
	var parsedData map[string]interface{}
	parsedData = make(map[string]interface{})
	parsedData = goconf.Parse("go.conf")

	fmt.Println(parsedData)
}
```