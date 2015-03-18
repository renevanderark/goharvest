package oai

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// Perform an HTTP GET request using the OAI Requests fields
// and return an OAI Response reference
func (req *OAIRequest) Perform() (oaiResponse *OAIResponse) {
	// Perform the GET request
	resp, err := http.Get(req.String())
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
