package integration_test

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strconv"
	"strings"
	"testing"
	"time"

	"code.cloudfoundry.org/grootfs/integration"
	"code.cloudfoundry.org/grootfs/integration/runner"
	"code.cloudfoundry.org/grootfs/testhelpers"
	"code.cloudfoundry.org/lager"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var (
	GrootFSBin    string
	Runner        runner.Runner
	StorePath     string
	NamespacerBin string
	mountPath     string

	GrootUsername    string
	GrootUID         int
	GrootGID         int
	RegistryUsername string
	RegistryPassword string
	GrootfsTestUid   int
	GrootfsTestGid   int
)

const xfsMountPath = "/mnt/xfs-%d"

func TestGroot(t *testing.T) {
	var (
		TardisBin string
	)

	RegisterFailHandler(Fail)

	SynchronizedBeforeSuite(func() []byte {
		grootFSBin, err := gexec.Build("code.cloudfoundry.org/grootfs")
		Expect(err).NotTo(HaveOccurred())
		grootFSBin = integration.MakeBinaryAccessibleToEveryone(grootFSBin)

		tardisBin, err := gexec.Build("code.cloudfoundry.org/grootfs/store/filesystems/overlayxfs/tardis")
		Expect(err).NotTo(HaveOccurred())
		tardisBin = integration.MakeBinaryAccessibleToEveryone(tardisBin)
		testhelpers.SuidBinary(tardisBin)

		namespacerBin, err := gexec.Build("code.cloudfoundry.org/grootfs/integration/namespacer")
		Expect(err).NotTo(HaveOccurred())

		return []byte(grootFSBin + ":" + tardisBin + ":" + namespacerBin)
	}, func(data []byte) {
		var err error
		binaries := strings.Split(string(data), ":")
		GrootFSBin = string(binaries[0])
		TardisBin = string(binaries[1])
		tmpNamespacerBin := string(binaries[2])

		rand.Seed(time.Now().UnixNano())
		NamespacerBin = fmt.Sprintf("/tmp/namespacer-%d", rand.Int())
		_, _, err = runCommand(exec.Command("cp", tmpNamespacerBin, NamespacerBin))
		Expect(err).NotTo(HaveOccurred())

		grootUser, err := user.Lookup(os.Getenv("GROOTFS_USER"))
		Expect(err).NotTo(HaveOccurred())

		GrootUsername := grootUser.Username

		GrootUID, err = strconv.Atoi(grootUser.Uid)
		Expect(err).NotTo(HaveOccurred())

		GrootGID, err = strconv.Atoi(grootUser.Gid)
		Expect(err).NotTo(HaveOccurred())

		GrootfsTestUid, _ = strconv.Atoi(os.Getenv("GROOTFS_TEST_UID"))
		GrootfsTestGid, _ = strconv.Atoi(os.Getenv("GROOTFS_TEST_GID"))

		fmt.Fprintf(os.Stderr, "============> RUNNING %s TESTS (%d:%d) <=============", "OVERLAY-XFS", GrootfsTestUid, GrootfsTestGid)
	})

	SynchronizedAfterSuite(func() {
	}, func() {
		gexec.CleanupBuildArtifacts()
	})

	BeforeEach(func() {
		testhelpers.ReseedRandomNumberGenerator()

		mountPath = fmt.Sprintf(xfsMountPath, GinkgoParallelNode())
		StorePath = path.Join(mountPath, "store")

		RegistryUsername = os.Getenv("DOCKER_REGISTRY_USERNAME")
		RegistryPassword = os.Getenv("DOCKER_REGISTRY_PASSWORD")

		Runner = runner.Runner{
			GrootFSBin: GrootFSBin,
			StorePath:  StorePath,
			TardisBin:  TardisBin,
		}.WithLogLevel(lager.DEBUG).WithStderr(GinkgoWriter).RunningAsUser(GrootfsTestUid, GrootfsTestGid)
	})

	AfterEach(func() {
		testhelpers.CleanUpOverlayMounts(StorePath)

		Expect(os.RemoveAll(StorePath)).To(Succeed())
	})

	RunSpecs(t, "Integration Suite")
}

func writeMegabytes(outputPath string, mb int) error {
	_, stderr, err := runCommand(exec.Command("dd", "if=/dev/zero", fmt.Sprintf("of=%s", outputPath), "bs=1048576", fmt.Sprintf("count=%d", mb)))
	if err != nil {
		return errors.New(stderr)
	}
	return nil
}

func mountByDefault() bool {
	return GrootfsTestUid == 0
}

func runCommand(command *exec.Cmd) (string, string, error) {
	stdout, stderr := gbytes.NewBuffer(), gbytes.NewBuffer()
	command.Stdout = io.MultiWriter(GinkgoWriter, stdout)
	command.Stderr = io.MultiWriter(GinkgoWriter, stderr)
	err := command.Run()
	return string(stdout.Contents()), string(stderr.Contents()), err
}
