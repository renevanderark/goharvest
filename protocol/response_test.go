package protocol

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestResponseToken(t *testing.T) {

	Convey("get false for a response that does not have a resumption token", t, func() {
		response := Response{}
		So(response.HasResumptionToken(), ShouldBeFalse)
	})

	Convey("get true for a response that has resumption token in response.ListIdentifiers.ResumptionToken ", t, func() {
		response := Response{
			ListIdentifiers: ListIdentifiers{
				ResumptionToken: "metadataPrefix%3Dovai_dc%26offset%3D275",
			},
		}
		So(response.HasResumptionToken(), ShouldBeTrue)
	})

	Convey("get true for a response that has resumption token in response.ListRecords.ResumptionToken ", t, func() {
		response := Response{
			ListRecords: ListRecords{
				ResumptionToken: "metadataPrefix%3Doaaai_dc%26offset%3D275",
			},
		}

		So(response.HasResumptionToken(), ShouldBeTrue)
	})

	Convey("Get resumption token from response.ListIdentifiers.ResumptionToken ", t, func() {
		response := Response{
			ListIdentifiers: ListIdentifiers{
				ResumptionToken: "metadataPrefix%3Doai_dc%26offset%3D375",
			},
		}
		So(response.GetResumptionToken(), ShouldEqual, "metadataPrefix%3Doai_dc%26offset%3D375")
	})

	Convey("Get resumption token from response.ListRecords.ResumptionToken ", t, func() {
		response := Response{
			ListRecords: ListRecords{
				ResumptionToken: "metadataPrefix%3Doai_dc%26offset%3D275",
			},
		}

		So(response.GetResumptionToken(), ShouldEqual, "metadataPrefix%3Doai_dc%26offset%3D275")
	})
}
