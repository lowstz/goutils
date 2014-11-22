package requests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {

	Convey("Subject: All Https Requests\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			result, _ := GetHttpsRequest("http://google.com")
			So(string(result), ShouldEqual, 200)
		})
		Convey("Status Code Should Be 1200", func() {
			result, _ := GetRequest("http://google.com")
			So(string(result), ShouldEqual, 200)
		})

	})
}
