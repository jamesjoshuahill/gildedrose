package main_test

import (
	"os/exec"

	"io/ioutil"

	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Gilded Rose", func() {
	It("updates all items correctly", func() {
		expectedOutput := readTestData("golden_master_with_conjured_item.txt")
		cmd := exec.Command(binaryPath)

		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

		Expect(err).NotTo(HaveOccurred())
		Eventually(session).Should(gexec.Exit(0))
		Expect(string(session.Out.Contents())).To(Equal(expectedOutput))
	})
})

func readTestData(name string) string {
	path := filepath.Join("testdata", name)
	contents, err := ioutil.ReadFile(path)
	Expect(err).NotTo(HaveOccurred())
	return string(contents)
}
