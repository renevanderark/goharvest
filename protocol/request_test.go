package protocol

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRequestURL(t *testing.T) {
	request := Request{
		BaseURL: "http://eprints.uwe.ac.uk/",
	}
	Convey("get just address", t, func() {
		So(request.getFullURL(), ShouldEqual, "http://eprints.uwe.ac.uk/?")
	})

	Convey("get request URL with verb", t, func() {
		request2 := Request{
			BaseURL: "http://eprints.uwe.ac.uk/",
			verb:    "identify",
		}

		So(request2.getFullURL(), ShouldEqual, "http://eprints.uwe.ac.uk/?verb=identify")
	})

	Convey("get request URL with set", t, func() {
		request3 := Request{
			BaseURL: "http://eprints.uwe.ac.uk/",
			Set:     "history",
		}
		So(request3.getFullURL(), ShouldEqual, "http://eprints.uwe.ac.uk/?set=history")
	})

	Convey("get request URL with prefix", t, func() {
		request4 := Request{
			BaseURL:        "http://eprints.uwe.ac.uk/",
			MetadataPrefix: "prefix",
		}
		So(request4.getFullURL(), ShouldEqual, "http://eprints.uwe.ac.uk/?metadataPrefix=prefix")
	})

	Convey("get request URL with prefix, verb and set", t, func() {
		request5 := Request{
			BaseURL:        "http://eprints.uwe.ac.uk/",
			Set:            "history",
			MetadataPrefix: "prefix",
			verb:           "identify",
		}
		So(request5.getFullURL(), ShouldEqual, "http://eprints.uwe.ac.uk/?verb=identify&set=history&metadataPrefix=prefix")
	})

	Convey("get request URL with resumptionToken", t, func() {
		request6 := Request{
			BaseURL:         "http://eprints.uwe.ac.uk/",
			ResumptionToken: "metadataPrefix%3Doai_dc%26offset%3D275",
		}
		So(request6.getFullURL(), ShouldEqual, "http://eprints.uwe.ac.uk/?resumptionToken=metadataPrefix%3Doai_dc%26offset%3D275")
	})

	Convey("get request URL with identifier", t, func() {
		request7 := Request{
			BaseURL:    "http://eprints.uwe.ac.uk/",
			Identifier: "79",
		}
		So(request7.getFullURL(), ShouldEqual, "http://eprints.uwe.ac.uk/?identifier=79")
	})

	Convey("get request URL with from", t, func() {
		request8 := Request{
			BaseURL: "http://eprints.uwe.ac.uk/",
			From:    "2009",
		}
		So(request8.getFullURL(), ShouldEqual, "http://eprints.uwe.ac.uk/?from=2009")
	})

	Convey("get request URL with until", t, func() {
		request9 := Request{
			BaseURL: "http://eprints.uwe.ac.uk/",
			Until:   "2015",
		}
		So(request9.getFullURL(), ShouldEqual, "http://eprints.uwe.ac.uk/?until=2015")
	})

	Convey("get request full URL: prefix, verb, set, identifier, from, until and resumptionTokenidentifier", t, func() {
		request10 := Request{
			BaseURL:         "http://eprints.uwe.ac.uk/",
			Set:             "history",
			MetadataPrefix:  "prefix",
			verb:            "identify",
			Identifier:      "79",
			From:            "2009",
			Until:           "2015",
			ResumptionToken: "metadataPrefix%3Doai_dc%26offset%3D275",
		}
		expectedURL := "http://eprints.uwe.ac.uk/?verb=identify&set=history&metadataPrefix=prefix&identifier=79&from=2009&until=2015&resumptionToken=metadataPrefix%3Doai_dc%26offset%3D275"
		So(request10.getFullURL(), ShouldEqual, expectedURL)
	})
}

func TestParse(t *testing.T) {
	request := Request{
		BaseURL: "http://eprints.uwe.ac.uk/",
	}
	Convey("get just address", t, func() {
		wrongFormat := []byte("Invalid format for OAI ")
		_, err := request.Parse(wrongFormat)
		So(err, ShouldNotBeNil)
	})

	Convey("does not give error for a good format", t, func() {

		good := `<?xml version='1.0' encoding='UTF-8'?>
              <?xml-stylesheet type='text/xsl' href='/oai2.xsl' ?>
              <OAI-PMH xmlns="http://www.openarchives.org/OAI/2.0/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/ http://www.openarchives.org/OAI/2.0/OAI-PMH.xsd">
                <responseDate>2015-10-08T21:04:23Z</responseDate>
                <request verb="Identify">http://eprints.uwe.ac.uk/cgi/oai2</request>
                <Identify>
                  <repositoryName>UWE Research Repository</repositoryName>
                  <baseURL>http://eprints.uwe.ac.uk/cgi/oai2</baseURL>
                  <protocolVersion>2.0</protocolVersion>
                  <adminEmail>eprints@uwe.ac.uk</adminEmail>
                  <earliestDatestamp>2012-08-23T15:48:51Z</earliestDatestamp>
                  <deletedRecord>persistent</deletedRecord>
                  <granularity>YYYY-MM-DDThh:mm:ssZ</granularity>
                  <description>
                    <oai-identifier xmlns="http://www.openarchives.org/OAI/2.0/oai-identifier" xsi:schemaLocation="http://www.openarchives.org/OAI/2.0/oai-identifier http://www.openarchives.org/OAI/2.0/oai-identifier.xsd">
                      <scheme>oai</scheme>
                      <repositoryIdentifier>eprints.uwe.ac.uk</repositoryIdentifier>
                      <delimiter>:</delimiter>
                      <sampleIdentifier>oai:eprints.uwe.ac.uk:6923</sampleIdentifier>
                    </oai-identifier>
                  </description>
                  <description>
                    <eprints xmlns="http://www.openarchives.org/OAI/1.1/eprints" xsi:schemaLocation="http://www.openarchives.org/OAI/1.1/eprints http://www.openarchives.org/OAI/1.1/eprints.xsd">
                      <content>
                        <text>OAI Site description has not been configured.
              </text>
                      </content>
                      <metadataPolicy>
                        <text>No metadata policy defined.
              This server has not yet been fully configured.
              Please contact the admin for more information, but if in doubt assume that
              NO rights at all are granted to this data.
              </text>
                      </metadataPolicy>
                      <dataPolicy>
                        <text>No data policy defined.
              This server has not yet been fully configured.
              Please contact the admin for more information, but if in doubt assume that
              NO rights at all are granted to this data.
              </text>
                      </dataPolicy>
                      <submissionPolicy>
                        <text>No submission-data policy defined.
              This server has not yet been fully configured.
              </text>
                      </submissionPolicy>
                      <comment>This system is running eprints server software (EPrints 3.3.12) developed at the University of Southampton. For more information see http://www.eprints.org/</comment>
                    </eprints>
                  </description>
                </Identify>
              </OAI-PMH>`
		goodFormat := []byte(good)
		_, err := request.Parse(goodFormat)
		So(err, ShouldBeNil)
	})
}
