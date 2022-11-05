package dult

import (
	"github.com/ac-i/dult/speed"
)

// QuickDULT tests upload and download speed of a network connection.
func QuickDULT(opt speed.Options) (*speed.TestResult, error) {
	return speed.QuickDULT(opt) //nolint:wrapcheck
}
