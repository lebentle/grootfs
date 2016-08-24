package groot_test

import (
	"io/ioutil"
	"os/exec"
	"path"

	"code.cloudfoundry.org/grootfs/integration"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Create", func() {
	var imagePath string

	BeforeEach(func() {
		var err error
		imagePath, err = ioutil.TempDir("", "")
		Expect(err).NotTo(HaveOccurred())

		Expect(ioutil.WriteFile(path.Join(imagePath, "foo"), []byte("hello-world"), 0644)).To(Succeed())
	})

	Context("when no --store option is given", func() {
		It("uses the default store path", func() {
			Expect("/var/lib/grootfs/bundles").ToNot(BeAnExistingFile())

			cmd := exec.Command(GrootFSBin, "create", imagePath, "random-id")
			sess, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Eventually(sess).Should(gbytes.Say("making directory `/var/lib/grootfs`"))
			Expect(err).NotTo(HaveOccurred())
			// It will fail at this point, because /var/lib/grootfs doesn't exist
			Eventually(sess).Should(gexec.Exit(1))
		})
	})

	Context("when two rootfses are using the same image", func() {
		It("isolates them", func() {
			bundle := integration.CreateBundle(GrootFSBin, StorePath, imagePath, "random-id")
			anotherBundle := integration.CreateBundle(GrootFSBin, StorePath, imagePath, "another-random-id")
			Expect(ioutil.WriteFile(path.Join(bundle.RootFSPath(), "bar"), []byte("hello-world"), 0644)).To(Succeed())
			Expect(path.Join(anotherBundle.RootFSPath(), "bar")).NotTo(BeARegularFile())
		})
	})

	Context("when the id is already being used", func() {
		BeforeEach(func() {
			Expect(integration.CreateBundle(GrootFSBin, StorePath, imagePath, "random-id")).NotTo(BeNil())
		})

		It("fails and produces a useful error", func() {
			cmd := exec.Command(GrootFSBin, "--store", StorePath, "create", imagePath, "random-id")
			sess, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Eventually(sess.Out).Should(gbytes.Say("bundle for id `random-id` already exists"))
			Expect(err).NotTo(HaveOccurred())
			Eventually(sess).Should(gexec.Exit(1))
		})
	})

	Context("when groot does not have permissions to apply the requested mapping", func() {
		It("returns the newuidmap output in the stdout", func() {
			cmd := exec.Command(
				GrootFSBin, "--store", StorePath,
				"create",
				"--uid-mapping", "1:1:65000",
				imagePath,
				"some-id",
			)

			buffer := gbytes.NewBuffer()
			sess, err := gexec.Start(cmd, buffer, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Expect(sess.Wait()).NotTo(gexec.Exit(0))

			Eventually(buffer).Should(gbytes.Say(`range [\[\)0-9\-]* -> [\[\)0-9\-]* not allowed`))
		})

		It("does not leak the bundle directory", func() {
			cmd := exec.Command(
				GrootFSBin, "--store", StorePath,
				"create",
				"--uid-mapping", "1:1:65000",
				imagePath,
				"some-id",
			)

			sess, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Expect(sess.Wait()).NotTo(gexec.Exit(0))

			Expect(path.Join(StorePath, "bundles", "some-id")).ToNot(BeAnExistingFile())
		})
	})

	Context("when the id is not provided", func() {
		It("fails", func() {
			cmd := exec.Command(GrootFSBin, "--store", StorePath, "create", imagePath)
			sess, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(sess).Should(gexec.Exit(1))
		})
	})

	Context("when a mappings flag is invalid", func() {
		It("fails when the uid mapping is invalid", func() {
			cmd := exec.Command(
				GrootFSBin, "--store", StorePath,
				"create", imagePath,
				"--uid-mapping", "1:hello:65000",
				"some-id",
			)

			sess, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Expect(sess.Wait()).NotTo(gexec.Exit(0))
		})

		It("fails when the gid mapping is invalid", func() {
			cmd := exec.Command(
				GrootFSBin, "--store", StorePath,
				"create", imagePath,
				"--gid-mapping", "1:groot:65000",
				"some-id",
			)

			sess, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Expect(sess.Wait()).NotTo(gexec.Exit(0))
		})
	})
})
