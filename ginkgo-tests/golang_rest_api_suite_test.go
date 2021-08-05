package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGolangRestApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GolangRestApi Suite")
}

