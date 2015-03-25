package oai

// Harvest the identifiers of a complete OAI set
// send a reference of each Header to a channel
func (req *Request) ChannelHarvestIdentifiers(channels []chan *Header) {
	req.Verb = "ListIdentifiers"
	req.Harvest(func(resp *Response) {
		headers := resp.ListIdentifiers.Headers
		i := 0
		for _, header := range headers {
			channels[i] <- &header
			i++
			if i == len(channels) {
				i = 0
			}
		}

		// If there is no more resumption token, send nil to all
		// the channels to signal the harvest is done
		hasResumptionToken, _ := resp.ResumptionToken()
		if !hasResumptionToken {
			for _, channel := range channels {
				channel <- nil
			}
		}
	})
}
