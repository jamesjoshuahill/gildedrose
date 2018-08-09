package gildedrose_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGildedrose(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gilded Rose Suite")
}
