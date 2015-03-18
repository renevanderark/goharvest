package oai

// Harvest the identifiers of a complete OAI set
// send a reference of each OAIHeader to a channel
func (req *OAIRequest) ChannelHarvestIdentifiers(channels []chan *OAIHeader) {
	req.Verb = "ListIdentifiers"
	req.Harvest(func(resp *OAIResponse) {
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
