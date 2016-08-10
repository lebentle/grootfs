package unpacker_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"code.cloudfoundry.org/grootfs/cloner"
	"code.cloudfoundry.org/grootfs/cloner/unpacker"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Tar", func() {
	var (
		logger lager.Logger

		imgPath    string
		bundlePath string
		rootFSPath string

		tarUnpacker *unpacker.TarUnpacker

		stream *gbytes.Buffer
	)

	BeforeEach(func() {
		var err error

		bundlePath, err = ioutil.TempDir("", "")
		Expect(err).NotTo(HaveOccurred())
		rootFSPath = path.Join(bundlePath, "rootfs")

		tarUnpacker = unpacker.NewTarUnpacker()

		logger = lagertest.NewTestLogger("test-graph")

		imgPath, err = ioutil.TempDir("", "")
		Expect(err).NotTo(HaveOccurred())
		Expect(ioutil.WriteFile(path.Join(imgPath, "a_file"), []byte("hello-world"), 0600)).To(Succeed())
	})

	JustBeforeEach(func() {
		stream = gbytes.NewBuffer()
		sess, err := gexec.Start(exec.Command("tar", "-c", "-C", imgPath, "."), stream, nil)
		Expect(err).NotTo(HaveOccurred())
		Eventually(sess).Should(gexec.Exit(0))
	})

	AfterEach(func() {
		Expect(os.RemoveAll(imgPath)).To(Succeed())
		Expect(os.RemoveAll(bundlePath)).To(Succeed())
	})

	It("does write the image contents in the rootfs directory", func() {
		Expect(tarUnpacker.Unpack(logger, cloner.UnpackSpec{
			Stream:     stream,
			RootFSPath: rootFSPath,
		})).To(Succeed())

		filePath := path.Join(rootFSPath, "a_file")
		Expect(filePath).To(BeARegularFile())
		contents, err := ioutil.ReadFile(filePath)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(contents)).To(Equal("hello-world"))
	})

	Context("when it has whiteout files", func() {
		BeforeEach(func() {
			Expect(os.Mkdir(rootFSPath, 0755)).To(Succeed())

			// Add some pre-existing files in the rootfs
			Expect(ioutil.WriteFile(path.Join(rootFSPath, "b_file"), []byte(""), 0600)).To(Succeed())
			Expect(os.Mkdir(path.Join(rootFSPath, "a_dir"), 0755)).To(Succeed())
			Expect(ioutil.WriteFile(path.Join(rootFSPath, "a_dir", "a_file"), []byte(""), 0600)).To(Succeed())
			Expect(os.Mkdir(path.Join(rootFSPath, "b_dir"), 0755)).To(Succeed())
			Expect(ioutil.WriteFile(path.Join(rootFSPath, "b_dir", "a_file"), []byte(""), 0600)).To(Succeed())

			// Add some whiteouts
			Expect(ioutil.WriteFile(path.Join(imgPath, ".wh.b_file"), []byte(""), 0600)).To(Succeed())
			Expect(os.Mkdir(path.Join(imgPath, "a_dir"), 0755)).To(Succeed())
			Expect(ioutil.WriteFile(path.Join(imgPath, "a_dir", ".wh.a_file"), []byte(""), 0600)).To(Succeed())
			Expect(ioutil.WriteFile(path.Join(imgPath, ".wh.b_dir"), []byte(""), 0600)).To(Succeed())
		})

		It("deletes the pre-existing files", func() {
			Expect(tarUnpacker.Unpack(logger, cloner.UnpackSpec{
				Stream:     stream,
				RootFSPath: rootFSPath,
			})).To(Succeed())

			Expect(path.Join(rootFSPath, "b_file")).NotTo(BeAnExistingFile())
			Expect(path.Join(rootFSPath, "a_dir", "a_file")).NotTo(BeAnExistingFile())
		})

		It("deletes the pre-existing directories", func() {
			Expect(tarUnpacker.Unpack(logger, cloner.UnpackSpec{
				Stream:     stream,
				RootFSPath: rootFSPath,
			})).To(Succeed())

			Expect(path.Join(rootFSPath, "b_dir")).NotTo(BeAnExistingFile())
		})

		It("does not leak the whiteout files", func() {
			Expect(tarUnpacker.Unpack(logger, cloner.UnpackSpec{
				Stream:     stream,
				RootFSPath: rootFSPath,
			})).To(Succeed())

			Expect(path.Join(rootFSPath, ".wh.b_file")).NotTo(BeAnExistingFile())
			Expect(path.Join(rootFSPath, "a_dir", ".wh.a_file")).NotTo(BeAnExistingFile())
			Expect(path.Join(rootFSPath, ".wh.b_dir")).NotTo(BeAnExistingFile())
		})
	})

	Context("when it fails to untar", func() {
		JustBeforeEach(func() {
			stream = gbytes.NewBuffer()
			stream.Write([]byte("not-a-tar"))
		})

		It("returns an error", func() {
			Expect(tarUnpacker.Unpack(logger, cloner.UnpackSpec{
				Stream:     stream,
				RootFSPath: rootFSPath,
			})).NotTo(Succeed())
		})

		It("returns the command output", func() {
			Expect(tarUnpacker.Unpack(logger, cloner.UnpackSpec{
				Stream:     stream,
				RootFSPath: rootFSPath,
			})).To(
				MatchError(ContainSubstring("tar:")),
			)
		})
	})

	Context("when creating the target directory fails", func() {
		It("returns an error", func() {
			err := tarUnpacker.Unpack(logger, cloner.UnpackSpec{
				Stream:     stream,
				RootFSPath: "/some-destination/bundles/1000",
			})

			Expect(err).To(MatchError(ContainSubstring("making destination directory")))
		})
	})

	Context("when the target directory exists", func() {
		It("still works", func() {
			Expect(os.Mkdir(rootFSPath, 0770)).To(Succeed())

			Expect(tarUnpacker.Unpack(logger, cloner.UnpackSpec{
				Stream:     stream,
				RootFSPath: rootFSPath,
			})).To(Succeed())
		})
	})
})