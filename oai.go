// Data structures for the OAI-PMH protocol defines:
// OAIRequest to represent a request URL and querystring to the OAI service
// OAIResponse and child fields for marshalling the XML responses

package oai

// Represents a request URL to an OAI-PMH service
type OAIRequest struct {
	BaseUrl, Set, MetadataPrefix, Verb, Identifier, ResumptionToken, From, Until string
}

type OAIHeader struct {
	Status string `xml:"status,attr"`
	Identifier string `xml:"identifier"`
	DateStamp string `xml:"datestamp"`
	SetSpec []string `xml:"setSpec"`
}

type Metadata struct {
	Body []byte `xml:",innerxml"`
}

type About struct {
	Body []byte `xml:",innerxml"`
}

type OAIRecord struct {
	Header OAIHeader `xml:"header"`
	Metadata Metadata `xml:"metadata"`
	About About `xml:"about"`
}


type ListIdentifiers struct {
	Header []OAIHeader `xml:"header"`
	ResumptionToken string `xml:"resumptionToken"`
}

type ListRecords struct {
	Record []OAIRecord `xml:"record"`
	ResumptionToken string `xml:"resumptionToken"`
}

type GetRecord struct {
	Record OAIRecord `xml:"record"`
}

type RequestNode struct {
	Verb string `xml:"verb,attr"`
	Set string `xml:"set,attr"`
	MetadataPrefix string `xml:"metadataPrefix,attr"`
}

type OAIError struct {
	Code string `xml:"code,attr"`
	Message string `xml:",chardata"`
}

type MetadataFormat struct {
	MetadataPrefix string `xml:"metadataPrefix"`
	Schema string `xml:"schema"`
	MetadataNamespace string `xml:"metadataNamespace"`
}

type ListMetadataFormats struct {
	MetadataFormat []MetadataFormat `xml:"metadataFormat"`
}

type SetDescription struct {
	Body []byte `xml:",innerxml"`
}

type Set struct {
	SetSpec string `xml:"setSpec"`
	SetName string `xml:"setName"`
	SetDescription SetDescription `xml:"setDescription"`

}

type ListSets struct {
	Set []Set `xml:"set"`
}

// The struct representation of an OAI-PMH XML response
type OAIResponse struct {
	ResponseDate string `xml:"responseDate"`
	Request RequestNode `xml:"request"`
	Error OAIError `xml:"error"`
	ListIdentifiers ListIdentifiers `xml:"ListIdentifiers"`
	ListRecords ListRecords `xml:"ListRecords"`
	GetRecord GetRecord `xml:"GetRecord"`
	ListMetadataFormats ListMetadataFormats `xml:"ListMetadataFormats"`
	ListSets ListSets `xml:"ListSets"`
}

