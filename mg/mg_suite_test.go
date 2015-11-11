package mg_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMg(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mg Suite")
}
