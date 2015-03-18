package oai

// Harvest the identifiers of a complete OAI set
// call the identifier callback function for each Header
func (req *Request) HarvestIdentifiers(callback func(*Header)) {
	req.Verb = "ListIdentifiers"
	req.Harvest(func(resp *Response) {
		headers := resp.ListIdentifiers.Headers
		for _, header := range headers {
			callback(&header)
		}
	})
}
