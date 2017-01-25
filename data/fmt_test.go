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

func runIntToStr(n int, fn intToStrFunc) float64 {
	start := time.Now()
	for i := 0; i < n; i++ {
		fn(i)
	}
	return time.Since(start).Seconds()
}

func TestIntToStr(t *testing.T) {
	n := 10000
	t1 := runIntToStr(n, fmtInt)
	t2 := runIntToStr(n, fmtfInt)
	t3 := runIntToStr(n, convInt)
	fmt.Printf("fmt.Sprint   * %d : %.6fs\n", n, t1)
	fmt.Printf("fmt.Sprintf  * %d : %.6fs\n", n, t2)
	fmt.Printf("strconv.Itoa * %d : %.6fs\n", n, t3)

	Convey("Test string conversion", t, func() {

		Convey("there is not much different between fmt.Sprint and fmt.Sprintf", func() {
			So(t2, ShouldAlmostEqual, t1, 0.01)
		})

		Convey("strconv.Itoa is the fastest one", func() {
			So(t3, ShouldBeLessThan, t2)
		})
	})
}
