package flusher_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-solution-api/internal/flusher"
	"github.com/ozoncp/ocp-solution-api/internal/mocks"
	"github.com/ozoncp/ocp-solution-api/internal/models"
)

var _ = Describe("Flusher", func() {
	var (
		ctrl      *gomock.Controller
		mockRepo  *mocks.MockRepo
		batchSize int
		f         flusher.Flusher
		solutions []models.Solution
		remaining []models.Solution
		err       error
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
	})

	JustBeforeEach(func() {
		if f, err = flusher.New(mockRepo, batchSize); f != nil {
			remaining, err = f.FlushSolutions(context.Background(), solutions)
		}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("nil Repo", func() {
		BeforeEach(func() {
			batchSize = 1
			mockRepo = nil
		})

		It("", func() {
			Expect(f).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(ContainSubstring("got nil Repo"))
		})
	})

	Context("zero batch size", func() {
		BeforeEach(func() {
			batchSize = 0
		})

		It("", func() {
			Expect(f).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(ContainSubstring("batchSize < 1"))
		})
	})

	Context("negative batch size", func() {
		BeforeEach(func() {
			batchSize = -1
		})

		It("", func() {
			Expect(f).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(ContainSubstring("batchSize < 1"))
		})
	})

	Context("noop with empty solutions", func() {
		BeforeEach(func() {
			batchSize = 2
			solutions = []models.Solution{}

			mockRepo.EXPECT().AddSolutions(context.Background(), gomock.Any()).Return(nil).Times(0)
		})

		It("", func() {
			Expect(f).ShouldNot(BeNil())
			Expect(remaining).Should(BeNil())
			Expect(err).Should(BeNil())
		})
	})

	Context("saves all solutions", func() {
		BeforeEach(func() {
			batchSize = 2
			solutions = []models.Solution{{}, {}, {}}

			mockRepo.EXPECT().AddSolutions(context.Background(), gomock.Any()).Return(nil).Times(2)
		})

		It("", func() {
			Expect(f).ShouldNot(BeNil())
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
			solutions = []models.Solution{{}, {}, {}, {}, {}}
			halfSize = len(solutions) / 2
			batchSize = halfSize
			repoError = errors.New("out of memory")

			gomock.InOrder(
				mockRepo.EXPECT().AddSolutions(context.Background(), gomock.Len(batchSize)).Return(nil),
				mockRepo.EXPECT().AddSolutions(context.Background(), gomock.Any()).Return(repoError),
			)
		})

		It("", func() {
			Expect(f).ShouldNot(BeNil())
			Expect(err).Should(BeEquivalentTo(repoError))
			Expect(remaining).ShouldNot(BeNil())
			Expect(len(remaining)).To(Equal(len(solutions) - batchSize))
		})
	})
})
