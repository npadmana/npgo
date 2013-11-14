package sf_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sf Suite")
}
