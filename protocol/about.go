package protocol

// About is an optional and repeatable container to hold data about
// the metadata part of the record.
//
// The contents of an about container must conform to an XML Schema.
// Individual implementation communities may create XML Schema that define
// specific uses for the contents of about containers. Two common uses of
// about containers are:
// - rights statements: some repositories may find it desirable to attach
// terms of use to the metadata they make available through the OAI-PMH.
// No specific set of XML tags for rights expression is defined by OAI-PMH,
// but the about container is provided to allow for encapsulating
// community-defined rights tags.
// - provenance statements: One suggested use of the about container is
// to indicate the provenance of a metadata record, e.g. whether it has
// been harvested itself and if so from which repository, and when.
// An XML Schema for such a provenance container, as well as some
// supporting information is available from the accompanying
// Implementation Guidelines document.
//
// http://www.openarchives.org/OAI/openarchivesprotocol.html#Record
type About struct {
	Body []byte `xml:",innerxml"`
}

// String returns the string representation
func (ab About) String() string {
	return string(ab.Body)
}
