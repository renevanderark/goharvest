package main

import (
	"fmt"
	"github.com/renevanderark/goharvest/oai"
)

// Dump a snippet of the Record metadata
func dump(resp *oai.Response) {
	fmt.Printf("%s\n\n", resp.GetRecord.Record.Metadata.Body[0:500])
}

// Performs a GetRecord request for the record identified by the OAI Header
func getRecord(hdr *oai.Header) {
	req := &oai.Request{
		BaseUrl:        "http://services.kb.nl/mdo/oai",
		Set:            "DTS",
		MetadataPrefix: "dcx",
		Verb:           "GetRecord",
		Identifier:     hdr.Identifier,
	}

	req.Harvest(dump)
}

// Demonstrates harvesting using the ListIdentifiers verb with HarvestIdentifiers
func main() {
	req := &oai.Request{
		BaseUrl:        "http://services.kb.nl/mdo/oai",
		Set:            "DTS",
		MetadataPrefix: "dcx",
		Verb:           "ListIdentifiers",
		From:           "2012-09-06T014:00:00.000Z",
	}

	// HarvestIdentifiers passes each individual OAI header to the getRecord
	// function as an Header object
	req.HarvestIdentifiers(getRecord)
}
