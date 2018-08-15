package main_test

import (
	"os/exec"
	"path/filepath"

	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Gilded Rose", func() {
	It("updates all items correctly", func() {
		expectedOutput := readTestData("golden_master.txt")

		session, err := gexec.Start(exec.Command(binaryPath), GinkgoWriter, GinkgoWriter)

		Expect(err).NotTo(HaveOccurred())
		Eventually(session).Should(gexec.Exit(0))
		Expect(string(session.Out.Contents())).To(Equal(expectedOutput))
	})

	PIt("updates all items correctly including conjured", func() {
		expectedOutput := readTestData("golden_master_with_conjured_item.txt")

		session, err := gexec.Start(exec.Command(binaryPath), GinkgoWriter, GinkgoWriter)

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
