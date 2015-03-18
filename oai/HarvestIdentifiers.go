package oai

// Harvest the identifiers of a complete OAI set
// call the identifier callback function for each OAIHeader
func (req *OAIRequest) HarvestIdentifiers(callback func(*OAIHeader)) {
	req.Verb = "ListIdentifiers"
	req.Harvest(func(resp *OAIResponse) {
		headers := resp.ListIdentifiers.Headers
		for _, header := range headers {
			callback(&header)
		}
	})
}
