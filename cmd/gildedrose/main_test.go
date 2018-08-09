package main_test

import (
	"os/exec"

	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Gilded Rose", func() {
	Context("when run for 30 days", func() {
		It("updates the items correctly", func() {
			contents, err := ioutil.ReadFile("../../GOLDEN_MASTER.txt")
			Expect(err).NotTo(HaveOccurred())
			expectedOutput := string(contents)
			cmd := exec.Command(binaryPath)

			session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))
			Expect(string(session.Out.Contents())).To(Equal(expectedOutput))
		})
	})
})
