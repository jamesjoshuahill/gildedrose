package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var binaryPath string

func TestGildedRose(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var _ = SynchronizedBeforeSuite(func() []byte {
	path, err := gexec.Build("github.com/jamesjoshuahill/gildedrose/cmd/gildedrose")
	Expect(err).NotTo(HaveOccurred())
	return []byte(path)
}, func(data []byte) {
	binaryPath = string(data)
})
