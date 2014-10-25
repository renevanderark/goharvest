package main

import (
	"github.com/renevanderark/goharvest/oai"
	"fmt"
	"bufio"
	"os"
)

func waitForKey() {
	fmt.Printf("\n--- Press ENTER to continue demo ---\n")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}


// Print the OAI Response object to stdout
func dump(resp *oai.OAIResponse) {
	_, resTok := resp.ResumptionToken()
	fmt.Printf("%#v\n", resp)

	if resTok != "" {
		fmt.Printf("---- Resumption token = \"%s\" ----\n\n", resTok)
	}
	waitForKey()
}

func main() {
	// Perform Identify, pass dump func as callback
	req := &oai.OAIRequest{
		BaseUrl: "http://services.kb.nl/mdo/oai",
		Verb: "Identify",
	}
	fmt.Printf("Identify:\n%s", req)
	waitForKey()
	req.Harvest(dump)

	// Perform ListSets, pass dump func as callback
	req = &oai.OAIRequest{
		BaseUrl: "http://services.kb.nl/mdo/oai",
		Verb: "ListSets",
	}
	fmt.Printf("ListSets:\n%s", req)
	waitForKey()
	req.Harvest(dump)

	// Perform ListMetadataFormats, pass dump func as callback
	req = &oai.OAIRequest{
		BaseUrl: "http://memory.loc.gov/cgi-bin/oai2_0",
		Verb: "ListMetadataFormats",
	}
	fmt.Printf("ListMetadataFormats:\n%s", req)
	waitForKey()
	req.Harvest(dump)

	// Perform GetRecord, pass dump func as callback
	req = &oai.OAIRequest{
		BaseUrl: "http://services.kb.nl/mdo/oai", 
		Set: "DTS",
		MetadataPrefix: "dcx",
		Verb: "GetRecord",
		Identifier: "DTS:dts:7929:mpeg21",
	}
	fmt.Printf("GetRecord: \n%s", req)
	waitForKey()
	req.Harvest(dump)

	// Perform ListIdentifiers, pass dump func as callback:
	// req.Harvest will iterate until out of resumption tokens
	// at each iteration dump will be called with an *oai.OAIResponse
	req = &oai.OAIRequest{
		BaseUrl: "http://services.kb.nl/mdo/oai", 
		Set: "DTS",
		MetadataPrefix: "dcx",
		Verb: "ListIdentifiers",
		From: "2012-09-06T014:00:00.000Z",
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
		From: "2010-07-19T20:01:36Z",
	}
	fmt.Printf("ListRecords:\n%s", req)
	waitForKey()
	req.Harvest(dump)
}
