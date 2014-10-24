// Methods for harvesting an OAI-PMH service
package oai

import (
	"net/http"
	"encoding/xml"
	"io/ioutil"
	"strings"
)

// Determine the resumption token in this OAIResponse
func (resp *OAIResponse) ResumptionToken() (hasResumptionToken bool, resumptionToken string) {
	hasResumptionToken = false
	resumptionToken = ""
	if resp == nil { return }

	// First attempt to obtain a resumption token from a ListIdentifiers response
	resumptionToken =  resp.ListIdentifiers.ResumptionToken

	// Then attempt to obtain a resumption token from a ListRecords response
	if resumptionToken == "" {
		resumptionToken =  resp.ListRecords.ResumptionToken
	}

	// If a non-empty resumption token turned up it can safely inferred that...
	if resumptionToken != "" { hasResumptionToken = true }

	return
}

// String representation of the OAI Request
func (req *OAIRequest) String() string {
	qs := []string{req.BaseUrl, "?set=", req.Set, "&metadataPrefix=", req.MetadataPrefix, "&verb=", req.Verb}

	if req.ResumptionToken != "" {
		qs = append(qs, "&resumptionToken=")
		qs = append(qs, req.ResumptionToken)
	}

	if req.Identifier != "" {
		qs = append(qs, "&identifier=")
		qs = append(qs, req.Identifier)
	}

	if req.From != "" {
		qs = append(qs, "&from=")
		qs = append(qs, req.From)
	}

	if req.Until != "" {
		qs = append(qs, "&until=")
		qs = append(qs, req.Until)
	}
	return strings.Join(qs, "")
}

// Perform an HTTP GET request using the OAI Requests fields
// and return an OAI Response reference
func (req *OAIRequest) Perform() (oaiResponse *OAIResponse) {
	// Perform the GET request
	resp, err := http.Get(req.String())
	if err != nil { panic(err) }

	// Make sure the response body object will be closed after
	// reading all the content body's data
	defer resp.Body.Close()

	// Read all the data
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { panic(err) }

	// Unmarshall all the data 
	err = xml.Unmarshal(body, &oaiResponse)
	if err != nil { panic(err) }

	return
}

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

// Reads OAI PMH response XML from a file
func FromFile(filename string) (oaiResponse *OAIResponse) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil { panic(err) }

	// Unmarshall all the data 
	err = xml.Unmarshal(bytes, &oaiResponse)
	if err != nil { panic(err) }

	return
}

