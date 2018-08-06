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
	RunSpecs(t, "Gilded Rose Suite")
}

var _ = SynchronizedBeforeSuite(func() []byte {
	binaryPath, err := gexec.Build("github.com/jamesjoshuahill/gildedrose")
	Expect(err).NotTo(HaveOccurred())
	return []byte(binaryPath)
}, func(data []byte) {
	binaryPath = string(data)
})
