package protocol

// Description is an extensible mechanism for communities to
// describe their repositories.
//
// For example, the description container could be used to include
// collection-level metadata in the response to the Identify request.
// Implementation Guidelines are available to give directions with
// this respect.
// Each description container must be accompanied by the URL of an
// XML schema describing the structure of the description container.
//
// http://www.openarchives.org/OAI/openarchivesprotocol.html#Identify
//
type Description struct {
	Body []byte `xml:",innerxml"`
}

// String returns the string representation
func (ab Description) String() string {
	return string(ab.Body)
}
