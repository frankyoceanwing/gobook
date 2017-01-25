package data

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFloatCompare(t *testing.T) {

	Convey("Test float", t, func() {
		a := 0.1 // float64
		Convey("When b's type is float64", func() {
			b := 0.1
			So(a == b, ShouldBeTrue)
		})

		Convey("When b's type is float32", func() {
			var b float32 = 0.1
			So(a == float64(b), ShouldBeFalse)
		})
	})
}
