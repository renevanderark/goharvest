package oai


// Harvest the identifiers of a complete OAI set
// call the identifier callback function for each OAIHeader
func (req *OAIRequest) HarvestRecords(callback func(*OAIRecord)) {
	req.Verb = "ListRecords"
	req.Harvest(func(resp *OAIResponse) {
		records := resp.ListRecords.Records
		for _, record := range records {
			callback(&record)
		}
	})
}
