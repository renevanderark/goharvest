package protocol

// ListMetadataFormats holds the formats received from server
//
// OAI-PMH supports the dissemination of records in multiple metadata
// formats from a repository.
// The ListMetadataFormats request returns the list of all metadata
// formats available from a repository
//
// http://www.openarchives.org/OAI/openarchivesprotocol.html#MetadataNamespaces
//
type ListMetadataFormats struct {
	MetadataFormat []MetadataFormat `xml:"metadataFormat"`
}

// MetadataFormat is a metadata format available from a repository
//
// It contains:
// 1. The metadataPrefix - a string to specify the metadata format in OAI-PMH
// requests issued to the repository. metadataPrefix consists of any valid URI
//  unreserved characters. metadataPrefix arguments are used in ListRecords,
//  ListIdentifiers, and GetRecord requests to retrieve records, or the headers
//   of records that include metadata in the format specified by the
//    metadataPrefix;
// 2. The metadata schema URL - the URL of an XML schema to test validity of
// metadata expressed according to the format;
// 3. The XML namespace URI that is a global identifier of the metadata format.
//  (http://www.w3.org/TR/1999/REC-xml-names-19990114/Overview.html)
//
// http://www.openarchives.org/OAI/openarchivesprotocol.html#MetadataNamespaces
//
type MetadataFormat struct {
	MetadataPrefix    string `xml:"metadataPrefix"`
	Schema            string `xml:"schema"`
	MetadataNamespace string `xml:"metadataNamespace"`
}
