package main

import (
	"fmt"

	"github.com/renevanderark/goharvest/oai"
)

// Print the OAI Response object to stdout
func dump(resp *oai.Response) {
	fmt.Printf("%#v\n", resp)

}

func main() {
	req := (&oai.Request{
		BaseUrl: "http://www.edshare.soton.ac.uk/cgi/oai2",
		Verb:    "Identify",
	})
	req.Harvest(dump)

}
