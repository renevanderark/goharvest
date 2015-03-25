package oai

// Harvest the identifiers of a complete OAI set
// call the identifier callback function for each Header
func (req *Request) HarvestRecords(callback func(*Record)) {
	req.Verb = "ListRecords"
	req.Harvest(func(resp *Response) {
		records := resp.ListRecords.Records
		for _, record := range records {
			callback(&record)
		}
	})
}
