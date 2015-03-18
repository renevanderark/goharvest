package main

import (
	"fmt"
	"github.com/renevanderark/goharvest/oai"
	"time"
)

type Digest struct {
	receivingChannels int
	baseRequest       *oai.OAIRequest
}

func dump(resp *oai.OAIResponse) {
	fmt.Printf("%s\n\n", resp.GetRecord.Record.Metadata.Body[0:1000])
}

func (digest *Digest) getRecord(identifier string) {
	req := &oai.OAIRequest{
		BaseUrl:        digest.baseRequest.BaseUrl,
		MetadataPrefix: digest.baseRequest.MetadataPrefix,
		Verb:           "GetRecord",
		Identifier:     identifier,
	}
	fmt.Println(req)
	req.Harvest(dump)
}

func (digest *Digest) digestIdentifiers(c chan *oai.OAIHeader) {
	hdr := <-c

	if hdr != nil {
		digest.getRecord(hdr.Identifier)
		digest.digestIdentifiers(c)
	} else {
		digest.receivingChannels--
		fmt.Printf("OPEN: %d\n\n", digest.receivingChannels)
	}
}

func main() {
	req := &oai.OAIRequest{
		BaseUrl:        "http://services.kb.nl/mdo/oai",
		Set:            "DTS",
		MetadataPrefix: "didl",
		From:           "2012-09-06T014:00:00.000Z",
	}
	digestChannels := []chan *oai.OAIHeader{}

	digest := &Digest{
		receivingChannels: 16,
		baseRequest: &oai.OAIRequest{
			BaseUrl:        "http://services.kb.nl/mdo/oai",
			MetadataPrefix: "didl",
		},
	}

	for i := 0; i < digest.receivingChannels; i++ {
		digestChannels = append(digestChannels, make(chan *oai.OAIHeader))
		go digest.digestIdentifiers(digestChannels[i])
	}

	req.ChannelHarvestIdentifiers(digestChannels)

	for digest.receivingChannels > 0 {
		time.Sleep(1)
	}
}
