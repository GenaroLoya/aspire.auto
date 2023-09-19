package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Animals River testing", func() {

	Describe("Testing resolve vailid move", func() {
		var EntityTest []Entity

		BeforeEach(func() {
			EntityTest = make([]Entity, len(EntityList))
			copy(EntityTest, EntityList)
		})

		It("should return false when remove Col", func() {
			err, result := checkRemove(0, &EntityTest)
			Expect(err).To(BeNil())
			Expect(result).To(Equal(false))
		})

		It("should return false when remove Wolf", func() {
			err, result := checkRemove(1, &EntityTest)
			Expect(err).To(BeNil())
			Expect(result).To(Equal(false))
		})

		It("should return false when remove Sheep", func() {
			err, result := checkRemove(2, &EntityTest)
			Expect(err).To(BeNil())
			Expect(result).To(Equal(true))
		})
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Animal River Suite")
}
