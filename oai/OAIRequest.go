// Data structure for the OAI-PMH protocol request:
package oai

import "strings"

// Represents a request URL and query string to an OAI-PMH service
type Request struct {
	BaseUrl, Set, MetadataPrefix, Verb, Identifier, ResumptionToken, From, Until string
}

// String representation of the OAI Request
func (req *Request) String() string {
	qs := []string{}

	add := func(name, value string) {
		if value != "" {
			qs = append(qs, name + "=" + value)
		}
	}

	add("verb", req.Verb)
	add("set", req.Set)
	add("metadataPrefix", req.MetadataPrefix)
	add("resumptionToken", req.ResumptionToken)
	add("identifier", req.Identifier)
	add("from", req.From)
	add("until", req.Until)

	return strings.Join([]string{req.BaseUrl, "?", strings.Join(qs, "&")}, "")
}
