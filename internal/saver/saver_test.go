package saver_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-solution-api/internal/flusher"
	"github.com/ozoncp/ocp-solution-api/internal/mocks"
	"github.com/ozoncp/ocp-solution-api/internal/models"
	"github.com/ozoncp/ocp-solution-api/internal/saver"
)

var _ = Describe("Saver", func() {
	var (
		ctrl                *gomock.Controller
		mockRepo            *mocks.MockRepo
		batchSize           int
		f                   flusher.Flusher
		err                 error
		solutions           []models.Solution
		capacity            uint
		forgetAllOnOverflow bool
		s                   saver.Saver
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		solutions = []models.Solution{
			*models.NewSolution(uint64(1), uint64(1)),
			*models.NewSolution(uint64(2), uint64(1)),
			*models.NewSolution(uint64(3), uint64(1)),
			*models.NewSolution(uint64(4), uint64(1)),
			*models.NewSolution(uint64(5), uint64(1)),
		}
	})

	JustBeforeEach(func() {
		f, err = flusher.New(mockRepo, batchSize)
		if s, err = saver.New(capacity, f, forgetAllOnOverflow); s != nil {
			for _, solution := range solutions {
				s.Save(solution)
			}
			err = s.Close()
		}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("nil Saver with 0 capacity", func() {
		BeforeEach(func() {
			batchSize = 2
			capacity = 0

			mockRepo.EXPECT().AddSolutions(gomock.Any()).Return(nil).Times(0)
		})

		It("", func() {
			Expect(f).ShouldNot(BeNil())
			Expect(s).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(ContainSubstring("zero Saver capacity"))
		})
	})

	Context("nil Saver with nil Flusher", func() {
		BeforeEach(func() {
			batchSize = 0
			capacity = 1

			mockRepo.EXPECT().AddSolutions(gomock.Any()).Return(nil).Times(0)
		})

		It("", func() {
			Expect(f).Should(BeNil())
			Expect(s).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(ContainSubstring("got nil Flusher"))
		})
	})

	Context("saves all", func() {
		BeforeEach(func() {
			batchSize = 1
			capacity = 1

			mockRepo.EXPECT().AddSolutions(gomock.Any()).Return(nil).Times(5)
		})

		It("", func() {
			Expect(f).ShouldNot(BeNil())
			Expect(s).ShouldNot(BeNil())
			Expect(err).Should(BeNil())
		})
	})

	Context("saves part", func() {
		var (
			repoError error
		)

		BeforeEach(func() {
			batchSize = 1
			capacity = 1
			repoError = errors.New("out of memory")

			gomock.InOrder(
				mockRepo.EXPECT().AddSolutions(gomock.Len(batchSize)).Return(nil),
				mockRepo.EXPECT().AddSolutions(gomock.Any()).Return(repoError),
			)
		})

		It("", func() {
			Expect(f).ShouldNot(BeNil())
			Expect(s).ShouldNot(BeNil())
			Expect(err).ShouldNot(BeNil())
			Expect(err.Error()).Should(ContainSubstring("out of memory"))
			Expect(err.Error()).Should(ContainSubstring("lost 4 solutions"))
		})
	})
})
