package action_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMetalJanitorAction(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Metal Janitor Suite")
}

func getTestData(name string) []byte {
	data, err := ioutil.ReadFile(filepath.Join("testdata", name))
	Expect(err).NotTo(HaveOccurred())

	return data
}
