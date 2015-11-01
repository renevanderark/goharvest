package protocol

// Metadata is a single manifestation of the metadata from an item.
//
// The OAI-PMH supports items with multiple manifestations (formats)
// of metadata.
// At a minimum, repositories must be able to return records with metadata
// expressed in the Dublin Core format, without any qualification. Optionally,
// a repository may also disseminate other formats of metadata.
// The specific metadata format of the record to be disseminated is specified
// by means of an argument
// -- the metadataPrefix
// -- in the GetRecord or ListRecords request that produces the record.
// The ListMetadataFormats request returns the list of all metadata formats
// available from a repository, or for a specific item (which can be specified
// as an argument to the ListMetadataFormats request).
//
// http://www.openarchives.org/OAI/openarchivesprotocol.html#Record
type Metadata struct {
	Body []byte `xml:",innerxml"`
}

// GoString returns the body as a string
func (md Metadata) String() string {
	return string(md.Body)
}
