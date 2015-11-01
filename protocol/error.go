package protocol

// Error represents an OAI error
//
// In event of an error or exception condition, repositories must
// indicate OAI-PMH errors, distinguished from HTTP Status-Codes,
// by including one or more error elements in the response.
// While one error element is sufficient to indicate the presence
// of the error or exception condition, repositories should report
// all errors or exceptions that arise from processing the request.
// Each error element must have a code attribute that must be from the
// following table; each error element may also have a free text string
// value to provide information about the error that is useful to a human
// reader. These strings are not defined by the OAI-PMH.
//
// Error Codes
// - badArgument	The request includes illegal arguments, is missing
// required arguments, includes a repeated argument, or values for arguments
// have an illegal syntax.
// - badResumptionToken	The value of the resumptionToken argument is
// invalid or expired.
// - badVerb	Value of the verb argument is not a legal OAI-PMH verb, the
// verb argument is missing, or the verb argument is repeated.
// - cannotDisseminateFormat	The metadata format identified by the value
// given for the metadataPrefix argument is not supported by the item or by
// the repository.
// - idDoesNotExist	The value of the identifier argument is unknown or illegal
//  in this repository.
// - noRecordsMatch	The combination of the values of the from, until, set and
//  metadataPrefix arguments results in an empty list.
// - noMetadataFormats	There are no metadata formats available for the
//  specified item.
// - noSetHierarchy	The repository does not support sets.
//
// http://www.openarchives.org/OAI/openarchivesprotocol.html#ErrorConditions
//
type Error struct {
	Code    string `xml:"code,attr"`
	Message string `xml:",chardata"`
}
