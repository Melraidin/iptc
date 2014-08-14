iptc
====

IPTC reader - Go (golang) wrapper for libiptcdata

Dependencies
============

Requires libiptcdata.

On OS X using Homebrew:
```brew install libiptcdata```

Usage
=====

```go
package main

import (
	"log"
	"os"

	"github.com/melraidin/iptc"
)

func main() {
	data, err := iptc.Open(os.Args[1])

	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	log.Printf("%v\n", data)
}
```

Output:

```
$ go run read.go resources/caption.jpg
2014/08/14 15:01:00 map[1:map[90:[1 b   2 5   4 7]] 2:map[0:2 116:Copyright 2014. All rights reserved. 120:Processed with VSCOcam with f2 preset]]
```
