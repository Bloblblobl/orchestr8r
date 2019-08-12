package media_manager

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMediaManager(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MediaManager Suite")
}
