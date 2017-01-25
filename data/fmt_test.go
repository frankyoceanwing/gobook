package data

import (
	"strconv"
	"testing"
	"time"

	"fmt"

	. "github.com/smartystreets/goconvey/convey"
)

type intToStrFunc func(int) string

func fmtInt(i int) string {
	return fmt.Sprint(i)
}

func fmtfInt(i int) string {
	return fmt.Sprintf("%d", i)
}

func convInt(i int) string {
	return strconv.Itoa(i)
}

func runIntToStr(fn intToStrFunc) float64 {
	start := time.Now()
	for i := 0; i < 100000; i++ {
		fn(i)
	}
	return time.Since(start).Seconds()
}

func TestIntToStr(t *testing.T) {
	t1 := runIntToStr(fmtInt)
	t2 := runIntToStr(fmtfInt)
	t3 := runIntToStr(convInt)

	Convey("Test string conversion", t, func() {

		Convey("fmt.Sprintf is faster than fmt.Sprint", func() {
			So(t2, ShouldBeLessThan, t1)
		})

		Convey("strconv.Itoa is the fastest one", func() {
			So(t3, ShouldBeLessThan, t2)
		})
	})
}
