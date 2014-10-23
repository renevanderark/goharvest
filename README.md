go-oai
======

First steps into parallel OAI scrape lib. TODO:
- support date parameters
- support channels / examples for channels
- example for parsing metadata content body

Example
----
Some example code for this package

```go
package main

import (
	"github.com/renevanderark/oai"
	"fmt"
)

// Print the OAI Record's metadata body to stdout
func dumpMeta(resp *oai.OAIResponse) {
	fmt.Printf("%s\n", resp.GetRecord.Record.Metadata.Body)
}

// Print the OAI Response object to stdout
func dump(resp *oai.OAIResponse) {
	_, resTok := resp.ResumptionToken()
	fmt.Printf("%+v\n", resp)
	fmt.Printf("---- %s ----\n\n", resTok)
}

func main() {

	// Perform a GetRecord, pass dumpMeta lambda as callback
	req := &oai.OAIRequest{
		BaseUrl: "http://services.kb.nl/mdo/oai", 
		Set: "DTS",
		MetadataPrefix: "dcx",
		Verb: "GetRecord",
		Identifier: "DTS:dts:7929:mpeg21",
	}
	fmt.Printf("%s\n", req)
	req.Harvest(dumpMeta)

	// Perform a ListIdentifiers, pass dump lambda as callback:
	// req.Harvest will iterate until out of resumption tokens
	// at each iteration dump will be called with an *oai.OAIResponse
	req = &oai.OAIRequest{
		BaseUrl: "http://services.kb.nl/mdo/oai", 
		Set: "DTS",
		MetadataPrefix: "didl",
		Verb: "ListIdentifiers", 
	}
	fmt.Printf("%s\n", req)
	req.Harvest(dump)
}
```
