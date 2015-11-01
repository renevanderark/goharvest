package protocol

// ListIdentifiers is an abbreviated verb form of ListRecords, retrieving only
// headers rather than records.
//
// Optional arguments permit selective harvesting of headers based on set
// membership and/or datestamp.
//
// Depending on the repository's support for deletions, a returned header
// may have a status attribute of "deleted" if a record matching the arguments
//
// specified in the request has been deleted.
//
// http://www.openarchives.org/OAI/openarchivesprotocol.html#ListIdentifiers
//
type ListIdentifiers struct {
	Headers         []Header `xml:"header"`
	ResumptionToken string   `xml:"resumptionToken"`
}

// Header contains the unique identifier of the item and properties necessary
// for selective harvesting.
//
// The header consists of the following parts:
// - the unique identifier -- the unique identifier of an item in a repository;
// - the datestamp -- the date of creation, modification or deletion of the
// record for the purpose of selective harvesting.
// - zero or more setSpec elements -- the set membership of the item for the
// purpose of selective harvesting.
// - an optional status attribute with a value of deleted indicates the withdrawal
// of availability of the specified metadata format for the item, dependent on
// the repository support for deletions.
//
// http://www.openarchives.org/OAI/openarchivesprotocol.html#Record
type Header struct {
	Identifier string   `xml:"identifier"`
	DateStamp  string   `xml:"datestamp"`
	SetSpec    []string `xml:"setSpec"`
	Status     string   `xml:"status,attr"`
}
