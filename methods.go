// Methods for harvesting an OAI-PMH service
package oai

import (
	"net/http"
	"encoding/xml"
	"io/ioutil"
	"strings"
)

func (resp *OAIResponse) ResumptionToken() (hasResumptionToken bool, resumptionToken string) {
	hasResumptionToken = false
	resumptionToken = ""
	if resp == nil { return }

	resumptionToken =  resp.ListIdentifiers.ResumptionToken

	if resumptionToken == "" {
		resumptionToken =  resp.ListRecords.ResumptionToken
	}

	if resumptionToken != "" { hasResumptionToken = true }

	return
}

func (req *OAIRequest) String() (url string) {
	qs := []string{req.BaseUrl, "?set=", req.Set, "&metadataPrefix=", req.MetadataPrefix, "&verb=", req.Verb}

	if req.ResumptionToken != "" {
		qs = append(qs, "&resumptionToken=")
		qs = append(qs, req.ResumptionToken)
	}

	if req.Identifier != "" {
		qs = append(qs, "&identifier=")
		qs = append(qs, req.Identifier)
	}
	return strings.Join(qs, "")
}

func (req *OAIRequest) Perform() (oaiResponse *OAIResponse) {
	resp, err := http.Get(req.String())
	if err != nil { panic(err) }

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { panic(err) }

	err = xml.Unmarshal(body, &oaiResponse)
	if err != nil { panic(err) }

	return
}

func (req *OAIRequest) Harvest(batchCallback func(*OAIResponse)) {
	oaiResponse := req.Perform()
	batchCallback(oaiResponse)
	hasResumptionToken, resumptionToken := oaiResponse.ResumptionToken()
	if hasResumptionToken == true {
		req.ResumptionToken = resumptionToken
		req.Harvest(batchCallback)
	}
}
