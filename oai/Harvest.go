package oai

// Perform a harvest of a complete OAI set, or simply one request
// call the batchCallback function argument with the OAI responses
func (req *OAIRequest) Harvest(batchCallback func(*OAIResponse)) {
	// Use Perform to get the OAI response
	oaiResponse := req.Perform()

	// Execute the callback function with the response
	batchCallback(oaiResponse)

	// Check for a resumptionToken
	hasResumptionToken, resumptionToken := oaiResponse.ResumptionToken()

	// Harvest further if there is a resumption token
	if hasResumptionToken == true {
		req.Set = ""
		req.MetadataPrefix = ""
		req.From = ""
		req.ResumptionToken = resumptionToken
		req.Harvest(batchCallback)
	}
}
