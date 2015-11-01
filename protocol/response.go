package protocol

// Response encapsulates the information from a harvest request
//
// All responses to OAI-PMH requests must be well-formed XML instance documents.
// Encoding of the XML must use the UTF-8 representation of Unicode. Character
// references, rather than entity references, must be used. Character references
// allow XML responses to be treated as stand-alone documents that can be
// manipulated without dependency on entity declarations external to the document.
//
// The XML data for all responses to OAI-PMH requests must validate against the
// XML Schema shown at the end of this section . As can be seen from that schema,
// responses to OAI-PMH requests have the following common markup:
//
// (ignored) The first tag output is an XML declaration where the version is
// always 1.0 and the encoding is always UTF-8, eg: <?xml version="1.0" encoding="UTF-8" ?>
// (ignored) The remaining content is enclosed in a root element with the name OAI-PMH.
//
// For all responses, the first two children of the root element are:
// - responseDate -- a UTCdatetime indicating the time and date that the
// response was sent. This must be expressed in UTC
// - request -- indicating the protocol request that generated this response.
//
// The third child of the root element is either:
// - error -- an element that must be used in case of an error or exception condition;
// - (Identify, ListMetadataFormats, ListSets, GetRecord, ListIdentifiers,
// ListRecords) an element with the same name as the verb of the respective
// OAI-PMH request.
//
// http://www.openarchives.org/OAI/openarchivesprotocol.html#XMLResponse
//
type Response struct {
	ResponseDate string      `xml:"responseDate"`
	Request      RequestNode `xml:"request"`
	Error        Error       `xml:"error"`

	Identify            Identify            `xml:"Identify"`
	ListMetadataFormats ListMetadataFormats `xml:"ListMetadataFormats"`
	ListSets            ListSets            `xml:"ListSets"`
	GetRecord           GetRecord           `xml:"GetRecord"`
	ListIdentifiers     ListIdentifiers     `xml:"ListIdentifiers"`
	ListRecords         ListRecords         `xml:"ListRecords"`
}

// RequestNode is indicating the protocol request that generated this response.
//
// The rules for generating the request element are as follows:
// 1. The content of the request element must always be the base URL of the protocol request;
// 2. The only valid attributes for the request element are the keys of the key=value pairs of protocol request. The attribute values must be the corresponding values of those key=value pairs;
// 3. In cases where the request that generated this response did not result in an error or exception condition, the attributes and attribute values of the request element must match the key=value pairs of the protocol request;
// 4. In cases where the request that generated this response resulted in a badVerb or badArgument error condition, the repository must return the base URL of the protocol request only. Attributes must not be provided in these cases.
//
// http://www.openarchives.org/OAI/openarchivesprotocol.html#XMLResponse
//
type RequestNode struct {
	Verb           string `xml:"verb,attr"`
	Set            string `xml:"set,attr"`
	MetadataPrefix string `xml:"metadataPrefix,attr"`
}

// HasResumptionToken determines if the request has or not a ResumptionToken
func (response *Response) HasResumptionToken() bool {
	return response.ListIdentifiers.ResumptionToken != "" || response.ListRecords.ResumptionToken != ""
}

// GetResumptionToken returns the resumption token or an empty string
// if it does not have a token
func (response *Response) GetResumptionToken() string {
	var resumptionToken string

	// First attempt to obtain a resumption token from a ListIdentifiers response
	resumptionToken = response.ListIdentifiers.ResumptionToken

	// Then attempt to obtain a resumption token from a ListRecords response
	if resumptionToken == "" {
		resumptionToken = response.ListRecords.ResumptionToken
	}
	return resumptionToken
}
