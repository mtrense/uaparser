package uaparser_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUaparser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Uaparser Suite")
}
