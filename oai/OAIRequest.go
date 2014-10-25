// Data structure for the OAI-PMH protocol request:
package oai

import "strings"

// Represents a request URL and query string to an OAI-PMH service
type OAIRequest struct {
	BaseUrl, Set, MetadataPrefix, Verb, Identifier, ResumptionToken, From, Until string
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

