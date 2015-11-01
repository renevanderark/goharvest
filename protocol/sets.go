package protocol

// ListSets represents a list of Sets
//
//  http://www.openarchives.org/OAI/openarchivesprotocol.html#Set
//
type ListSets struct {
	Set []Set `xml:"set"`
}

// Set is an optional construct for grouping items for the purpose of
// selective harvesting.
//
// Repositories may organize items into sets. Set organization may be
// flat, i.e. a simple list, or hierarchical.
// Multiple hierarchies with distinct, independent top-level nodes are allowed.
// Hierarchical organization of sets is expressed in the syntax of the setSpec
// parameter as described below.
// When a repository defines a set organization it must include set membership
// information in the headers of items returned in response to the
// ListIdentifiers, ListRecords and GetRecord requests.
//
// A Set has:
// - setSpec -- a colon [:] separated list indicating the path from the root
// of the set hierarchy to the respective node. Each element in the list is a
// string consisting of any valid URI unreserved characters, which must not
// contain any colons [:]. Since a setSpec forms a unique identifier for the
// set within the repository, it must be unique for each set. Flat set
// organizations have only sets with setSpec that do not contain any colons [:].
// - setName -- a short human-readable string naming the set.
// - setDescription -- an optional and repeatable container that may hold
// community-specific XML-encoded data about the set; the accompanying
// Implementation Guidelines document provides suggestions regarding the
// usage of this container.
//
// http://www.openarchives.org/OAI/openarchivesprotocol.html#Set
//
type Set struct {
	SetSpec        string      `xml:"setSpec"`
	SetName        string      `xml:"setName"`
	SetDescription Description `xml:"setDescription"`
}
