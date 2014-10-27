// Data structure for the OAI-PMH protocol request:
package oai

import "strings"

// Represents a request URL and query string to an OAI-PMH service
type OAIRequest struct {
	BaseUrl, Set, MetadataPrefix, Verb, Identifier, ResumptionToken, From, Until string
}

// String representation of the OAI Request
func (req *OAIRequest) String() string {
	var part []string
	qs := []string{}

	if req.Verb != "" {
		part = []string{"verb", req.Verb,}
		qs = append(qs, strings.Join(part, "="))
	}

	if req.Set != "" {
		part = []string{"set", req.Set,}
		qs = append(qs, strings.Join(part, "="))
	}

	if req.MetadataPrefix != "" {
		part = []string{"metadataPrefix", req.MetadataPrefix,}
		qs = append(qs, strings.Join(part, "="))
	}

	if req.ResumptionToken != "" {
		part = []string{"resumptionToken", req.ResumptionToken,}
		qs = append(qs, strings.Join(part, "="))
	}

	if req.Identifier != "" {
		part = []string{"identifier", req.Identifier,}
		qs = append(qs, strings.Join(part, "="))
	}

	if req.From != "" {
		part = []string{"from", req.From,}
		qs = append(qs, strings.Join(part, "="))

	}

	if req.Until != "" {
		part = []string{"until",req.Until,}
		qs = append(qs, strings.Join(part, "="))
	}

	return strings.Join([]string{req.BaseUrl, "?",strings.Join(qs, "&"),}, "")
}

