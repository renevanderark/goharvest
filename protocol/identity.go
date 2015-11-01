package protocol

// Identify is a verb used to retrieve information about a repository.
//
// Some of the information returned is required as part of the OAI-PMH.
// Repositories may also employ the Identify verb to return additional
// descriptive information.
// The response must include one instance of the following elements:
// - repositoryName : a human readable name for the repository;
// - baseURL : the base URL of the repository;
// - protocolVersion : the version of the OAI-PMH supported by the repository;
// - earliestDatestamp : a UTCdatetime that is the guaranteed lower limit
// of all datestamps recording changes, modifications, or deletions in the
// repository. A repository must not use datestamps lower than the one
// specified by the content of the earliestDatestamp element. earliestDatestamp
// must be expressed at the finest granularity supported by the repository.
// - deletedRecord : the manner in which the repository supports the notion
// of deleted records. Legitimate values are no ; transient ; persistent
// with meanings defined in the section on deletion.
// - granularity: the finest harvesting granularity supported by the repository.
// The legitimate values are YYYY-MM-DD and YYYY-MM-DDThh:mm:ssZ with meanings
// as defined in ISO8601.
//
// The response must include one or more instances of the following element:
// - adminEmail : the e-mail address of an administrator of the repository.
//
// The response may include multiple instances of the following optional
// elements:
//
// - compression : a compression encoding supported by the repository.
// The recommended values are those defined for the Content-Encoding header
// in Section 14.11 of RFC 2616 describing HTTP 1.1. A compression element
// should not be included for the identity encoding, which is implied.
// - description : an extensible mechanism for communities to describe their
// repositories. For example, the description container could be used to
// include collection-level metadata in the response to the Identify request.
// Implementation Guidelines are available to give directions with this respect.
// Each description container must be accompanied by the URL of an XML schema
// describing the structure of the description container.
//
// http://www.openarchives.org/OAI/openarchivesprotocol.html#Identify
type Identify struct {
	// must
	RepositoryName    string   `xml:"repositoryName"`
	BaseURL           string   `xml:"baseURL"`
	ProtocolVersion   string   `xml:"protocolVersion"`
	EarliestDatestamp string   `xml:"earliestDatestamp"`
	DeletedRecord     string   `xml:"deletedRecord"`
	Granularity       string   `xml:"granularity"`
	AdminEmail        []string `xml:"adminEmail"`
	// may
	Description []Description `xml:"description"`
}
