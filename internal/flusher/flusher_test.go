package flusher_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-solution-api/internal/flusher"
	"github.com/ozoncp/ocp-solution-api/internal/mocks"
	"github.com/ozoncp/ocp-solution-api/internal/solution"
)

var _ = Describe("Flusher", func() {
	var (
		ctrl      *gomock.Controller
		mockRepo  *mocks.MockRepo
		batchSize int
		f         flusher.Flusher
		solutions []solution.Solution
		remaining []solution.Solution
		err       error
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
	})

	JustBeforeEach(func() {
		f = flusher.New(mockRepo, batchSize)
		remaining, err = f.Flush(solutions)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("noop with empty solutions", func() {
		BeforeEach(func() {
			batchSize = 2
			solutions = []solution.Solution{}

			mockRepo.EXPECT().AddSolutions(gomock.Any()).Return(nil).Times(0)
		})

		It("", func() {
			Expect(remaining).Should(BeNil())
			Expect(err).Should(BeNil())
		})
	})

	Context("saves all solutions", func() {
		BeforeEach(func() {
			batchSize = 2
			solutions = []solution.Solution{{}, {}, {}}

			mockRepo.EXPECT().AddSolutions(gomock.Any()).Return(nil).Times(2)
		})

		It("", func() {
			Expect(remaining).Should(BeNil())
			Expect(err).Should(BeNil())
		})
	})

	Context("saves part of solutions", func() {
		var (
			halfSize  int
			repoError error
		)

		BeforeEach(func() {
			solutions = []solution.Solution{{}, {}, {}, {}, {}}
			halfSize = len(solutions) / 2
			batchSize = halfSize
			repoError = errors.New("out of memory")

			gomock.InOrder(
				mockRepo.EXPECT().AddSolutions(gomock.Len(batchSize)).Return(nil),
				mockRepo.EXPECT().AddSolutions(gomock.Any()).Return(repoError),
			)
		})

		It("", func() {
			Expect(err).Should(BeEquivalentTo(repoError))
			Expect(remaining).ShouldNot(BeNil())
			Expect(len(remaining)).To(Equal(len(solutions) - batchSize))
		})
	})
})
