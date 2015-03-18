// Data structure for the OAI-PMH protocol request:
package oai

import (
	"encoding/xml"
	"io/ioutil"
)

// Reads OAI PMH response XML from a file
func FromFile(filename string) (oaiResponse *Response) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// Unmarshall all the data
	err = xml.Unmarshal(bytes, &oaiResponse)
	if err != nil {
		panic(err)
	}

	return
}
