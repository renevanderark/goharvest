go-oai
======

First steps into parallel OAI scrape lib. TODO:
- support Identify verb
- human readable formatter for OAIResponse struct 
- support channels / examples for channels


Example
----
```go
package main

import (
	"github.com/renevanderark/oai"
	"fmt"
	"bufio"
	"os"
)

func waitForKey() {
	fmt.Printf("\n--- Press ENTER to continue demo ---\n")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}

// Print the OAI Record's metadata body to stdout
func dumpMeta(resp *oai.OAIResponse) {
	fmt.Printf("%s\n", resp.GetRecord.Record.Metadata.Body)
}

// Print the OAI Record's about body to stdout
func dumpAbout(resp *oai.OAIResponse) {
	fmt.Printf("%s\n", resp.GetRecord.Record.About.Body)
}

// Print the OAI Response object to stdout
func dump(resp *oai.OAIResponse) {
	_, resTok := resp.ResumptionToken()
	fmt.Printf("%+v\n", resp)
	fmt.Printf("---- %s ----\n\n", resTok)
}

func main() {
	// Perform ListSets, pass dump func as callback
	req := &oai.OAIRequest{
		BaseUrl: "http://services.kb.nl/mdo/oai",
		Verb: "ListSets",
	}
	fmt.Printf("ListSets:\n%s", req)
	waitForKey()
	req.Harvest(dump)
	waitForKey()

	// Perform ListMetadataFormats, pass dump func as callback
	req = &oai.OAIRequest{
		BaseUrl: "http://memory.loc.gov/cgi-bin/oai2_0",
		Verb: "ListMetadataFormats",
	}
	fmt.Printf("ListMetadataFormats:\n%s", req)
	waitForKey()
	req.Harvest(dump)
	waitForKey()

	// Perform GetRecord, pass dumpMeta func as callback
	req = &oai.OAIRequest{
		BaseUrl: "http://services.kb.nl/mdo/oai", 
		Set: "DTS",
		MetadataPrefix: "dcx",
		Verb: "GetRecord",
		Identifier: "DTS:dts:7929:mpeg21",
	}
	fmt.Printf("GetRecord: \n%s", req)
	waitForKey()
	req.Harvest(dumpMeta)
	waitForKey()

	// Perform ListIdentifiers, pass dump func as callback:
	// req.Harvest will iterate until out of resumption tokens
	// at each iteration dump will be called with an *oai.OAIResponse
	req = &oai.OAIRequest{
		BaseUrl: "http://services.kb.nl/mdo/oai", 
		Set: "DTS",
		MetadataPrefix: "dcx",
		Verb: "ListIdentifiers",
	}
	fmt.Printf("ListIdentifiers:\n%s", req)
	waitForKey()
	req.Harvest(dump)

	// Perform ListRecords, pass dump func as callback:
	// req.Harvest will iterate until out of resumption tokens
	// at each iteration dump will be called with an *oai.OAIResponse
	req = &oai.OAIRequest{
		BaseUrl: "http://memory.loc.gov/cgi-bin/oai2_0", 
		Set: "bbc",
		MetadataPrefix: "oai_dc",
		Verb: "ListRecords",
		From: "2008-01-01T00:00:00Z",
	}
	fmt.Printf("ListRecords:\n%s", req)
	waitForKey()
	req.Harvest(dump)
}
```
