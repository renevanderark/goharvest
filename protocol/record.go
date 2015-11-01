package protocol

// GetRecord is used to retrieve an individual metadata record from a repository.
//
// Required arguments specify the identifier of the item from which the record
// is requested and the format of the metadata that should be included in the
// record. Depending on the level at which a repository tracks deletions, a
// header with a "deleted" value for the status attribute may be returned,
// in case the metadata format specified by the metadataPrefix is no longer
// available from the repository or from the specified item.
//
// http://www.openarchives.org/OAI/openarchivesprotocol.html#GetRecord
//
type GetRecord struct {
	Record Record `xml:"record"`
}

// ListRecords is a verb used to harvest records from a repository.
//
// Optional arguments permit selective harvesting of records based on set
// membership and/or datestamp. Depending on the repository's support for
// deletions, a returned header may have a status attribute of "deleted"
// if a record matching the arguments specified in the request has been deleted.
// No metadata will be present for records with deleted status.
//
// http://www.openarchives.org/OAI/openarchivesprotocol.html#FlowControl
type ListRecords struct {
	Records         []Record `xml:"record"`
	ResumptionToken string   `xml:"resumptionToken"`
}

// Record is metadata expressed in a single format.
//
// A record is returned in an XML-encoded byte stream in response to an OAI-PMH
// request for metadata from an item.
// A record is identified unambiguously by the combination of the unique
// identifier of the item from which the record is available, the metadataPrefix
// identifying the metadata format of the record, and the datestamp of
// the record.
// The XML-encoding of records is organized into the following parts:
// header, metadata, about
//
// http://www.openarchives.org/OAI/openarchivesprotocol.html#Record
type Record struct {
	Header   Header   `xml:"header"`
	Metadata Metadata `xml:"metadata"`
	About    About    `xml:"about"`
}
