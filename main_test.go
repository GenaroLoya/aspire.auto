package main

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Aspires testing", func() {

	Describe("aspireLive", func() {
		It("should return the table with no error for a valid initial position", func() {
			table := []EnumState{DIRTY, CLEAN, CLEAN}
			initPos := 2
			initDir := TOR
			result, err := aspireLive(table, initPos, initDir)

			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(table))
		})

		It("should return an error for an invalid initial position", func() {
			table := []EnumState{CLEAN, DIRTY, CLEAN, DIRTY}
			initPos := 10
			initDir := TOR
			result, err := aspireLive(table, initPos, initDir)

			Expect(err).To(HaveOccurred())
			Expect(result).To(BeNil())
			Expect(err.Error()).To(Equal("Invalid initial position"))
		})

		It("should return an error for an empty table", func() {
			table := []EnumState{}
			initPos := 0
			initDir := TOR
			result, err := aspireLive(table, initPos, initDir)

			Expect(err).To(HaveOccurred())
			Expect(result).To(BeNil())
			Expect(err.Error()).To(Equal("Empty table"))
		})

		It("should return table[i] in zero", func() {
			table := []EnumState{CLEAN, DIRTY}
			initPos := 1
			initDir := TOR
			result, err := aspireLive(table, initPos, initDir)
			Expect(err).NotTo(HaveOccurred())
			Expect(result[initPos]).To(Equal(CLEAN))
		})

		It("should return all table in zero when arr is [1, 1]", func() {
			table := []EnumState{DIRTY, DIRTY}
			initPos := 0
			initDir := TOR
			result, err := aspireLive(table, initPos, initDir)

			expectZero := true

			fmt.Println(result)

			for _, in := range result {
				if in == DIRTY {
					expectZero = false
				}
			}

			Expect(err).NotTo(HaveOccurred())
			Expect(expectZero).To(Equal(true))
		})

		It("should return all zero in many Dirty table when init pos is 0", func() {
			table := []EnumState{DIRTY, DIRTY, DIRTY, DIRTY, DIRTY, DIRTY}
			initPos := 0
			initDir := TOR
			result, err := aspireLive(table, initPos, initDir)

			expectZero := true

			fmt.Println(result)

			for _, in := range result {
				if in == DIRTY {
					expectZero = false
				}
			}

			Expect(err).NotTo(HaveOccurred())
			Expect(expectZero).To(Equal(true))
		})

		It("should return all zero in partially Dirty table when init pos is any", func() {
			table := []EnumState{DIRTY, DIRTY, CLEAN, DIRTY, CLEAN, DIRTY}
			initPos := 0
			initDir := TOR
			result, err := aspireLive(table, initPos, initDir)

			expectZero := true

			fmt.Println(result)

			for _, in := range result {
				if in == DIRTY {
					expectZero = false
				}
			}

			Expect(err).NotTo(HaveOccurred())
			Expect(expectZero).To(Equal(true))
		})

		It("should return all zero in partially Dirty table when init pos is any and init dir is to left", func() {
			table := []EnumState{DIRTY, CLEAN, CLEAN, DIRTY, DIRTY, DIRTY}
			initPos := 0
			initDir := TOL
			result, err := aspireLive(table, initPos, initDir)

			expectZero := true

			fmt.Println(result)

			for _, in := range result {
				if in == DIRTY {
					expectZero = false
				}
			}

			Expect(err).NotTo(HaveOccurred())
			Expect(expectZero).To(Equal(true))
		})

		It("should return all zero in partially Dirty table when init pos is any and init dir is to left", func() {
			table := []EnumState{DIRTY, DIRTY, DIRTY, CLEAN, DIRTY}
			initPos := 3
			initDir := TOR
			result, err := aspireLive(table, initPos, initDir)

			expectZero := true

			fmt.Println(result)

			for _, in := range result {
				if in == DIRTY {
					expectZero = false
				}
			}

			Expect(err).NotTo(HaveOccurred())
			Expect(expectZero).To(Equal(true))
		})

		It("should return all zero in partially Dirty table when init pos is any and init dir is to left", func() {
			table := []EnumState{DIRTY, DIRTY, DIRTY}
			initPos := 3
			initDir := TOR
			result, err := aspireLive(table, initPos, initDir)

			expectZero := true

			fmt.Println(result)

			for _, in := range result {
				if in == DIRTY {
					expectZero = false
				}
			}

			Expect(err).NotTo(HaveOccurred())
			Expect(expectZero).To(Equal(true))
		})
	})

	Describe("Testing resolveAction", func() {
		var instructionsTest []Instruct

		BeforeEach(func() {
			instructionsTest = instructions
		})

		It("should return the expected action for a clean position", func() {
			result := resolveAction(CLEAN, instructionsTest)
			Expect(result.describAction).To(Equal(MOVE))
		})

		It("should return the expected action for a dirty position", func() {
			result := resolveAction(DIRTY, instructionsTest)
			Expect(result.describAction).To(Equal(CLEAR))
		})
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Aspire Suite")
}
