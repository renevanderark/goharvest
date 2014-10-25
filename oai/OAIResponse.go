// Data structure for the OAI-PMH protocol responses:
package oai

import (
	"fmt"
)

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
	Headers []OAIHeader `xml:"header"`
	ResumptionToken string `xml:"resumptionToken"`
}

type ListRecords struct {
	Records []OAIRecord `xml:"record"`
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

type Description struct {
	Body []byte `xml:",innerxml"`
}

type Set struct {
	SetSpec string `xml:"setSpec"`
	SetName string `xml:"setName"`
	SetDescription Description `xml:"setDescription"`

}

type ListSets struct {
	Set []Set `xml:"set"`
}

type Identify struct {
	RepositoryName string `xml:"repositoryName"`
	BaseURL string `xml:"baseURL"`
	ProtocolVersion string `xml:"protocolVersion"`
	AdminEmail []string `xml:"adminEmail"`
	EarliestDatestamp string `xml:"earliestDatestamp"`
	DeletedRecord string `xml:"deletedRecord"`
	Granularity string `xml:"granularity"`
	Description []Description `xml:"description"`
}

// The struct representation of an OAI-PMH XML response
type OAIResponse struct {
	ResponseDate string `xml:"responseDate"`
	Request RequestNode `xml:"request"`
	Error OAIError `xml:"error"`

	Identify Identify `xml:"Identify"`
	ListMetadataFormats ListMetadataFormats `xml:"ListMetadataFormats"`
	ListSets ListSets `xml:"ListSets"`
	GetRecord GetRecord `xml:"GetRecord"`
	ListIdentifiers ListIdentifiers `xml:"ListIdentifiers"`
	ListRecords ListRecords `xml:"ListRecords"`
}


// Formatter for Metadata content
func (md Metadata) GoString() string { return fmt.Sprintf("%s", md.Body); }

// Formatter for Description content
func (desc Description) GoString() string { return fmt.Sprintf("%s", desc.Body); }

// Formatter for About content
func (ab About) GoString() string { return fmt.Sprintf("%s", ab.Body); }
