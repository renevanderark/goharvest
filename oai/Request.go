package oai

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"
)

// Request represents a request URL and query string to an OAI-PMH service
type Request struct {
	BaseURL         string
	Set             string
	MetadataPrefix  string
	Verb            string
	Identifier      string
	ResumptionToken string
	From            string
	Until           string
}

// GetFullURL represents the OAI Request in a string format
func (request *Request) GetFullURL() string {
	array := []string{}

	add := func(name, value string) {
		if value != "" {
			array = append(array, name+"="+value)
		}
	}

	add("verb", request.Verb)
	add("set", request.Set)
	add("metadataPrefix", request.MetadataPrefix)
	add("resumptionToken", request.ResumptionToken)
	add("identifier", request.Identifier)
	add("from", request.From)
	add("until", request.Until)

	URL := strings.Join([]string{request.BaseURL, "?", strings.Join(array, "&")}, "")

	return URL
}

// HarvestIdentifiers arvest the identifiers of a complete OAI set
// call the identifier callback function for each Header
func (request *Request) HarvestIdentifiers(callback func(*Header)) {
	request.Verb = "ListIdentifiers"
	request.Harvest(func(resp *Response) {
		headers := resp.ListIdentifiers.Headers
		for _, header := range headers {
			callback(&header)
		}
	})
}

// HarvestRecords harvest the identifiers of a complete OAI set
// call the identifier callback function for each Header
func (request *Request) HarvestRecords(callback func(*Record)) {
	request.Verb = "ListRecords"
	request.Harvest(func(resp *Response) {
		records := resp.ListRecords.Records
		for _, record := range records {
			callback(&record)
		}
	})
}

// ChannelHarvestIdentifiers harvest the identifiers of a complete OAI set
// send a reference of each Header to a channel
func (request *Request) ChannelHarvestIdentifiers(channels []chan *Header) {
	request.Verb = "ListIdentifiers"
	request.Harvest(func(resp *Response) {
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

// Harvest perform a harvest of a complete OAI set, or simply one request
// call the batchCallback function argument with the OAI responses
func (request *Request) Harvest(batchCallback func(*Response)) {
	// Use Perform to get the OAI response
	oaiResponse := request.Perform()

	// Execute the callback function with the response
	batchCallback(oaiResponse)

	// Check for a resumptionToken
	hasResumptionToken, resumptionToken := oaiResponse.ResumptionToken()

	// Harvest further if there is a resumption token
	if hasResumptionToken == true {
		request.Set = ""
		request.MetadataPrefix = ""
		request.From = ""
		request.ResumptionToken = resumptionToken
		request.Harvest(batchCallback)
	}
}

// Perform an HTTP GET request using the OAI Requests fields
// and return an OAI Response reference
func (request *Request) Perform() (oaiResponse *Response) {

	resp, err := http.Get(request.GetFullURL())
	if err != nil {
		panic(err)
	}

	// Make sure the response body object will be closed after
	// reading all the content body's data
	defer resp.Body.Close()

	// Read all the data
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Unmarshall all the data
	err = xml.Unmarshal(body, &oaiResponse)
	if err != nil {
		panic(err)
	}

	return
}

// ResumptionToken determine the resumption token in this Response
func (resp *Response) ResumptionToken() (hasResumptionToken bool, resumptionToken string) {
	hasResumptionToken = false
	resumptionToken = ""
	if resp == nil {
		return
	}

	// First attempt to obtain a resumption token from a ListIdentifiers response
	resumptionToken = resp.ListIdentifiers.ResumptionToken

	// Then attempt to obtain a resumption token from a ListRecords response
	if resumptionToken == "" {
		resumptionToken = resp.ListRecords.ResumptionToken
	}

	// If a non-empty resumption token turned up it can safely inferred that...
	if resumptionToken != "" {
		hasResumptionToken = true
	}

	return
}
