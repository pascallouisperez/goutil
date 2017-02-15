package errors

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type ErrorsSuite struct{}

var _ = Suite(&ErrorsSuite{})

func (_ *ErrorsSuite) TestNiceMessage(c *C) {
	err := New("%s,%d", "message", 123)
	c.Assert(err, ErrorMatches, "errors_test.go:[0-9]+: message,123")
}
