goharvest
======

OAI-PMH harvester implementations for golang.

Hello OAI
---
```go
package main

import (
	"github.com/renevanderark/goharvest/oai"
	"fmt"
)

func main() {
	(&oai.Request{
		BaseURL:"http://services.kb.nl/mdo/oai", Set:"DTS", MetadataPrefix:"dcx",
		From: "2012-09-06T014:00:00.000Z",
	}).HarvestRecords(func (record *oai.Record) {
		fmt.Printf("%s\n\n", record.Metadata.Body[0:500])
	})
}
```


Demo sources
---
Sources for the demo's can be found in the bin dir

Prerequisites
---
- The go tool
- git

Get started
---

Clone the project into your Go workspace
```sh
$ cd $GOPATH/src
$ mkdir -p github.com/renevanderark
$ cd github.com/renevanderark
$ git clone https://github.com/renevanderark/goharvest.git
```

Starting the demo's:

```sh
$ go run $GOPATH/src/github.com/renevanderark/goharvest/bin/oai_demo/main.go
$ go run $GOPATH/src/github.com/renevanderark/goharvest/bin/oai_harvest_demo1/main.go
$ go run $GOPATH/src/github.com/renevanderark/goharvest/bin/oai_harvest_demo2/main.go
$ go run $GOPATH/src/github.com/renevanderark/goharvest/bin/hello_oai/main.go
$ go run $GOPATH/src/github.com/renevanderark/goharvest/bin/channel_harvest_demo1/main.go
```

