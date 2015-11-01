package protocol

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// Request encapsulates the information to get information using OAI-PMH protocol
type Request struct {
	BaseURL         string
	Set             string
	MetadataPrefix  string
	verb            string
	Identifier      string
	ResumptionToken string
	From            string
	Until           string
}

// Clear sets to default all the fields
func (request *Request) Clear() {
	request.Set = ""
	request.MetadataPrefix = ""
	request.verb = ""
	request.Identifier = ""
	request.ResumptionToken = ""
	request.From = ""
	request.Until = ""
}

// Identify returns the identification details of the repository
func (request *Request) Identify() ([]byte, error) {
	return request.getVerb("Identify")
}

// GetFormats returns the content of the request for formats
func (request *Request) GetFormats() ([]byte, error) {
	return request.getVerb("ListMetadataFormats")
}

// GetSets returns the content of the request to get all the sets
func (request *Request) GetSets() ([]byte, error) {
	return request.getVerb("ListSets")
}

// GetRecords returns the content of the request to get all
// the records for a particular metadata format
func (request *Request) GetRecords(prefix string) ([]byte, error) {
	request.MetadataPrefix = prefix
	return request.getVerb("ListRecords")
}

// GetIdentifiers returns the content of the request to get all
// the identifiers for a particular metadata format
func (request *Request) GetIdentifiers(prefix string) ([]byte, error) {
	request.MetadataPrefix = prefix
	return request.getVerb("ListIdentifiers")
}

func (request *Request) getVerb(verb string) ([]byte, error) {
	request.verb = verb
	return request.get()
}

// IsValidResponse parses a content and returns the error
func (request *Request) IsValidResponse(content []byte) error {
	_, err := request.Parse(content)
	if err != nil {
		return err
	}
	return nil
}

// Parse returns the content of the remote repository
func (request *Request) Parse(content []byte) (*Response, error) {
	response, err := request.getResponse(content)
	if err != nil {
		return response, err
	}
	return response, err
}

// It performs an HTTP request, reads the body of the request and returns it
func (request *Request) get() ([]byte, error) {
	var (
		content []byte
	)
	url := request.getFullURL()

	resp, err := http.Get(url)
	if err != nil {
		errorMsg := "<br />" + err.Error()
		message := "There was a problem while trying to connect to that address: " + errorMsg
		return content, errors.New(message)
	}
	defer resp.Body.Close()

	// Read all the data
	content, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		errorMsg := "<br />" + err.Error()
		message := "There was a problem while trying to read the content of the page: " + errorMsg
		return content, errors.New(message)
	}
	return content, nil
}

// GetFullURL returns the full URL address for the request
// scheme:[BaseURL][?][parameter=value[&]]
// see http://www.w3.org/Addressing/URL/url-spec.txt
func (request *Request) getFullURL() string {
	parameters := []string{}

	addParameter := func(name, value string) {
		if value != "" {
			parameters = append(parameters, name+"="+value)
		}
	}
	addParameter("verb", request.verb)
	addParameter("set", request.Set)
	addParameter("metadataPrefix", request.MetadataPrefix)
	addParameter("identifier", request.Identifier)
	addParameter("from", request.From)
	addParameter("until", request.Until)
	addParameter("resumptionToken", request.ResumptionToken)

	query := "?" + strings.Join(parameters, "&")
	URL := request.BaseURL + query

	return URL
}

// Parse an array of bytes into an OAI Response and returns the Response
func (request *Request) getResponse(body []byte) (*Response, error) {
	response := &Response{}
	err := xml.Unmarshal(body, &response)
	if err != nil {
		errorMsg := "<br />" + err.Error()
		message := "There was a problem while trying parsing the XML:" + errorMsg
		return response, errors.New(message)
	}
	return response, nil
}
