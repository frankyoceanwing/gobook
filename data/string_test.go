package data

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

type intsToStringFunc func([]int) string

func intsToString(ints []int) string {
	s := "["
	for i, v := range ints {
		if i != 0 {
			s += " ,"
		}
		s += strconv.Itoa(v)
	}
	s += "]"
	return s
}

func intsToStringWithJoin(ints []int) string {
	strs := make([]string, len(ints))
	for i, v := range ints {
		strs[i] = strconv.Itoa(v)
	}
	return "[" + strings.Join(strs, " ,") + "]"
}

func intsToStringWithBuffer(ints []int) string {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range ints {
		if i != 0 {
			buf.WriteString(" ,")
		}
		buf.WriteString(strconv.Itoa(v))
	}
	buf.WriteString("]")
	return buf.String()
}

func runIntsToString(fn intsToStringFunc) (string, float64) {
	ints := make([]int, 100000)
	for i := range ints {
		ints[i] = i + 1
	}
	start := time.Now()
	s := fn(ints)
	return s, time.Since(start).Seconds()
}

func TestIntsToString(t *testing.T) {
	s1, t1 := runIntsToString(intsToString)
	s2, t2 := runIntsToString(intsToStringWithBuffer)
	s3, t3 := runIntsToString(intsToStringWithJoin)

	Convey("Test string concatenation", t, func() {

		Convey("bytes.Buffer is faster than '+'", func() {
			So(s1 == s2, ShouldBeTrue)
			So(t2, ShouldBeLessThan, t1)
		})

		Convey("strings.Join is the fastest one", func() {
			So(s2 == s3, ShouldBeTrue)
			So(t3, ShouldBeLessThan, t2)
		})
	})
}
