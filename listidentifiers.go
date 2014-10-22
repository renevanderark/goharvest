package oai

import (
	"net/http"
	"io/ioutil"
	"strings"
	"fmt"
)

func listIdentifiers(baseUrl string, set string, metadataPrefix string, resumptionToken string) string {
	qs := []string{baseUrl, "?set=", set, "&metadataPrefix=", metadataPrefix, "&verb=ListIdentifiers"}

	if resumptionToken != "" {
		qs = append(qs, "&resumptionToken=")
		qs = append(qs, resumptionToken)
	}

	resp, err := http.Get(strings.Join(qs, ""))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s", body)
}

func (oai *OAI) ListIdentifiers() string {
	return listIdentifiers(oai.BaseUrl, oai.Set, oai.MetadataPrefix, oai.ResumptionToken);
}
